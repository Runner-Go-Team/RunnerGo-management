package handler

import (
	"fmt"
	"github.com/go-omnibus/proof"
	"golang.org/x/net/context"
	"gorm.io/gen"
	"gorm.io/gorm"
	"kp-management/internal/pkg/biz/consts"
	"kp-management/internal/pkg/dal"
	"kp-management/internal/pkg/dal/model"
	"kp-management/internal/pkg/dal/query"
	"time"
)

func TimedTaskExec() {
	// 开启定时任务轮询
	for {
		ctx := context.Background()
		tx := query.Use(dal.DB()).TimedTaskConf
		// 组装查询条件
		conditions := make([]gen.Condition, 0)
		// 当前时间
		nowTime := time.Now().Unix()
		conditions = append(conditions, tx.Status.Eq(consts.TimedTaskInExec))
		// 从数据库当中，查出当前需要执行的定时任务
		timedTaskData, err := tx.WithContext(ctx).Where(conditions...).Find()

		// 当前时间的 时，分
		nowTimeInfo := time.Unix(nowTime, 0)
		nowYear := nowTimeInfo.Year()
		nowMonth := nowTimeInfo.Month()
		nowDay := nowTimeInfo.Day()
		nowHour := nowTimeInfo.Hour()
		nowMinute := nowTimeInfo.Minute()
		nowWeekday := nowTimeInfo.Weekday()

		if err == nil { // 查到了数据
			proof.Infof("定时任务--查到了数据：", timedTaskData)
			// 组装运行计划参数
			for _, timedTaskInfo := range timedTaskData {
				// 获取定时任务的执行时间相关数据
				tm := time.Unix(timedTaskInfo.TaskExecTime, 0)
				taskYear := tm.Year()
				taskMonth := tm.Month()
				taskDay := tm.Day()
				taskHour := tm.Hour()
				taskMinute := tm.Minute()
				taskWeekday := tm.Weekday()

				// 排除过期的定时任务
				if timedTaskInfo.TaskCloseTime < nowTime {
					// 把当前定时任务状态变成已过期
					_, err := tx.WithContext(ctx).Where(tx.TeamID.Eq(timedTaskInfo.TeamID)).
						Where(tx.PlanID.Eq(timedTaskInfo.PlanID)).
						Where(tx.SenceID.Eq(timedTaskInfo.SenceID)).
						UpdateColumn(tx.Status, consts.TimedTaskTimeout)
					if err != nil {
						proof.Infof("定时任务过期状态修改失败，err：", err)
					}
					proof.Infof("定时任务--设置为过期：", timedTaskInfo.TaskCloseTime, " 当前时间：", nowTime)
					continue
				}

				// 根据不同的任务频次，进行不同的运行逻辑
				switch timedTaskInfo.Frequency {
				case 0: // 一次
					if taskYear != nowYear || taskMonth != nowMonth || taskDay != nowDay || taskHour != nowHour || taskMinute != nowMinute {
						continue
					}
					proof.Infof("定时任务--频次一次：通过可运行")
				case 1: // 每天
					// 比较当前时间是否等于定时任务的时间
					if taskHour != nowHour || taskMinute != nowMinute {
						continue
					}

				case 2: // 每周
					// 比较当前周几是否等于定时任务的时间
					if taskWeekday != nowWeekday || taskHour != nowHour || taskMinute != nowMinute {
						continue
					}

				case 3: // 每月
					// 比较当前每月几号是否等于定时任务的时间
					if taskDay != nowDay || taskHour != nowHour || taskMinute != nowMinute {
						continue
					}
				}

				// 给当前任务加分布式锁，防止重复执行
				timedTaskKey := "TimeTaskRun:" + fmt.Sprintf("%d", timedTaskInfo.SenceID)
				setRedisErr := dal.GetRDB().SetNX(ctx, timedTaskKey, 1, time.Second*65).Err()
				if setRedisErr != nil {
					continue
				}

				// 执行定时任务计划
				err := runTimedTask(ctx, timedTaskInfo)
				if err != nil {
					proof.Infof("定时任务运行失败，任务信息：", timedTaskInfo, " err：", err)
				}
			}
		} else if err != gorm.ErrRecordNotFound {
			proof.Infof("定时任务查询数据库出错，err：", err)
			continue
		}

		// 睡眠一分钟，再循环执行
		time.Sleep(59 * time.Second)
	}
}

func runTimedTask(ctx context.Context, timedTaskInfo *model.TimedTaskConf) error {
	// 开始执行计划
	sceneIds := make([]int64, 0, 1)
	sceneIds = append(sceneIds, timedTaskInfo.SenceID)
	runStressParams := RunStressReq{
		PlanID:  timedTaskInfo.PlanID,
		TeamID:  timedTaskInfo.TeamID,
		SceneID: sceneIds,
		UserID:  timedTaskInfo.UserID,
	}
	proof.Infof("定时任务--开始执行计划，参数：", runStressParams)
	// 进入执行计划方法
	_, runErr := RunStress(ctx, runStressParams)
	proof.Infof("定时任务--执行结果，runErr：", runErr)
	if runErr != nil {
		return runErr
	}
	return nil
}
