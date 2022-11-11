package report

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/shopspring/decimal"

	"gorm.io/gen/field"

	"github.com/go-omnibus/proof"
	"go.mongodb.org/mongo-driver/bson"
	"gorm.io/gen"

	"kp-management/internal/pkg/biz/record"

	"kp-management/internal/pkg/biz/consts"
	"kp-management/internal/pkg/dal"
	"kp-management/internal/pkg/dal/mao"
	"kp-management/internal/pkg/dal/query"
	"kp-management/internal/pkg/dal/rao"
	"kp-management/internal/pkg/packer"
)

func CountByTeamID(ctx context.Context, teamID int64) (int64, error) {
	tx := query.Use(dal.DB()).Report

	return tx.WithContext(ctx).Where(tx.TeamID.Eq(teamID)).Count()
}

func ListByTeamID2(ctx context.Context, teamID int64, limit, offset int, keyword string, startTimeSec, endTimeSec int64, taskType, taskMode, status, sortTag int32) ([]*rao.Report, int64, error) {

	tx := query.Use(dal.DB()).Report

	conditions := make([]gen.Condition, 0)
	conditions = append(conditions, tx.TeamID.Eq(teamID))

	if keyword != "" {
		var reportIDs []int64

		planReportIDs, err := KeywordFindPlan(ctx, teamID, keyword)
		if err != nil {
			return nil, 0, err
		}
		reportIDs = append(reportIDs, planReportIDs...)

		sceneReportIDs, err := KeywordFindScene(ctx, teamID, keyword)
		if err != nil {
			return nil, 0, err
		}
		reportIDs = append(reportIDs, sceneReportIDs...)

		userReportIDs, err := KeywordFindUser(ctx, keyword)
		if err != nil {
			return nil, 0, err
		}
		reportIDs = append(reportIDs, userReportIDs...)

		if len(reportIDs) > 0 {
			conditions = append(conditions, tx.ID.In(reportIDs...))
		} else {
			conditions = append(conditions, tx.ID.In(0))
		}
	}

	if startTimeSec > 0 && endTimeSec > 0 {
		startTime := time.Unix(startTimeSec, 0)
		endTime := time.Unix(endTimeSec, 0)
		conditions = append(conditions, tx.CreatedAt.Between(startTime, endTime))
	}

	if taskType > 0 {
		conditions = append(conditions, tx.TaskType.Eq(taskType))
	}

	if taskMode > 0 {
		conditions = append(conditions, tx.TaskMode.Eq(taskMode))
	}

	if status > 0 {
		conditions = append(conditions, tx.Status.Eq(status))
	}

	sort := make([]field.Expr, 0)

	if sortTag == 0 { // 默认排序
		sort = append(sort, tx.Rank.Desc())
		sort = append(sort, tx.ID.Desc())
	}
	if sortTag == 1 { // 创建时间倒序
		sort = append(sort, tx.CreatedAt.Desc())
	}
	if sortTag == 2 { // 创建时间正序
		sort = append(sort, tx.CreatedAt)
	}
	if sortTag == 3 { // 修改时间倒序
		sort = append(sort, tx.UpdatedAt.Desc())
	}
	if sortTag == 4 { // 修改时间正序
		sort = append(sort, tx.UpdatedAt)
	}

	reports, cnt, err := tx.WithContext(ctx).Where(conditions...).
		Order(sort...).
		FindByPage(offset, limit)

	if err != nil {
		return nil, 0, err
	}

	var userIDs []int64
	for _, r := range reports {
		userIDs = append(userIDs, r.RunUserID)
	}

	u := query.Use(dal.DB()).User
	users, err := u.WithContext(ctx).Where(u.ID.In(userIDs...)).Find()
	if err != nil {
		return nil, 0, err
	}

	return packer.TransReportModelToRaoReportList(reports, users), cnt, nil
}

func KeywordFindPlan(ctx context.Context, teamID int64, keyword string) ([]int64, error) {
	var planIDs []int64

	p := dal.GetQuery().Plan
	err := p.WithContext(ctx).Where(p.TeamID.Eq(teamID), p.Name.Like(fmt.Sprintf("%%%s%%", keyword))).Pluck(p.ID, &planIDs)
	if err != nil {
		return nil, err
	}

	if len(planIDs) == 0 {
		return nil, nil
	}

	var reportIDs []int64
	r := dal.GetQuery().Report
	err = r.WithContext(ctx).Where(r.PlanID.In(planIDs...)).Pluck(r.ID, &reportIDs)
	if err != nil {
		return nil, err
	}

	return reportIDs, nil
}

func KeywordFindScene(ctx context.Context, teamID int64, keyword string) ([]int64, error) {
	var sceneIDs []int64

	s := dal.GetQuery().Target
	err := s.WithContext(ctx).Where(s.TeamID.Eq(teamID), s.Name.Like(fmt.Sprintf("%%%s%%", keyword))).Pluck(s.ID, &sceneIDs)
	if err != nil {
		return nil, err
	}

	if len(sceneIDs) == 0 {
		return nil, nil
	}

	var reportIDs []int64
	r := dal.GetQuery().Report
	err = r.WithContext(ctx).Where(r.SceneID.In(sceneIDs...)).Pluck(r.ID, &reportIDs)
	if err != nil {
		return nil, err
	}

	return reportIDs, nil
}

func KeywordFindUser(ctx context.Context, keyword string) ([]int64, error) {
	var userIDs []int64

	u := query.Use(dal.DB()).User
	err := u.WithContext(ctx).Where(u.Nickname.Like(fmt.Sprintf("%%%s%%", keyword))).Pluck(u.ID, &userIDs)
	if err != nil {
		return nil, err
	}

	if len(userIDs) == 0 {
		return nil, nil
	}

	var reportIDs []int64
	r := dal.GetQuery().Report
	err = r.WithContext(ctx).Where(r.RunUserID.In(userIDs...)).Pluck(r.ID, &reportIDs)
	if err != nil {
		return nil, err
	}

	return reportIDs, nil
}

func DeleteReport(ctx context.Context, teamID, reportID, userID int64) error {
	allErr := dal.GetQuery().Transaction(func(tx *query.Query) error {
		r, err := tx.Report.WithContext(ctx).Where(tx.Report.ID.Eq(reportID)).First()
		if err != nil {
			return err
		}
		if _, err := tx.Report.WithContext(ctx).Where(tx.Report.TeamID.Eq(teamID), tx.Report.ID.Eq(reportID)).Delete(); err != nil {
			return err
		}

		if err := record.InsertDelete(ctx, teamID, userID, record.OperationOperateDeleteReport, fmt.Sprintf("%s %s", r.PlanName, r.SceneName)); err != nil {
			return err
		}
		return nil
	})

	if allErr != nil {
		proof.Infof("DeleteReport：删除失败")
		return allErr
	}

	// 把mongodb库里面的报告详情数据删掉
	collection := dal.GetMongo().Database(dal.MongoDB()).Collection(consts.CollectReportData)
	reportIdString := strconv.FormatInt(reportID, 10)
	findFilter := bson.M{"reportid": reportIdString}
	_, err := collection.DeleteOne(ctx, findFilter)
	if err != nil {
		proof.Infof("DeleteReport：删除失败")
	}

	return nil
}

func GetTaskDetail(ctx context.Context, req rao.GetReportTaskDetailReq) (*rao.ReportTask, error) {
	var detail mao.ReportTask
	collection := dal.GetMongo().Database(dal.MongoDB()).Collection(consts.CollectReportTask)

	err := collection.FindOne(ctx, bson.D{{"report_id", req.ReportID}}).Decode(&detail)
	if err != nil {
		proof.Error("mongo decode err", proof.WithError(err))
		return nil, err
	}

	r := query.Use(dal.DB()).Report
	ru, err := r.WithContext(ctx).Where(r.TeamID.Eq(req.TeamID), r.ID.Eq(req.ReportID)).First()
	if err != nil {
		proof.Error("req not found err", proof.WithError(err))
		return nil, err
	}

	u := query.Use(dal.DB()).User
	user, err := u.WithContext(ctx).Where(u.ID.Eq(ru.RunUserID)).First()
	if err != nil {
		proof.Error("user not found err", proof.WithError(err))
		return nil, err
	}

	// 从mongo查出编辑报告的数据列表
	collection = dal.GetMongo().Database(dal.MongoDB()).Collection(consts.CollectChangeReportConf)
	ChangeTaskConfDetail, _ := collection.Find(ctx, bson.D{{"report_id", req.ReportID}})

	changeTaskConf := make([]*mao.ChangeTaskConf, 0, 10)
	if err := ChangeTaskConfDetail.All(ctx, &changeTaskConf); err != nil {
		proof.Infof("没有查到编辑报告列表数据", proof.WithError(err))
	}

	modeConf := &rao.ModeConf{
		ReheatTime:       detail.ModeConf.ReheatTime,
		RoundNum:         detail.ModeConf.RoundNum,
		Concurrency:      detail.ModeConf.Concurrency,
		ThresholdValue:   detail.ModeConf.ThresholdValue,
		StartConcurrency: detail.ModeConf.StartConcurrency,
		Step:             detail.ModeConf.Step,
		StepRunTime:      detail.ModeConf.StepRunTime,
		MaxConcurrency:   detail.ModeConf.MaxConcurrency,
		Duration:         detail.ModeConf.Duration,
		CreatedTimeSec:   ru.CreatedAt.Unix(),
	}

	changeTaskConfData := &rao.ModeConf{
		ReheatTime:       detail.ModeConf.ReheatTime,
		RoundNum:         detail.ModeConf.RoundNum,
		Concurrency:      detail.ModeConf.Concurrency,
		ThresholdValue:   detail.ModeConf.ThresholdValue,
		StartConcurrency: detail.ModeConf.StartConcurrency,
		Step:             detail.ModeConf.Step,
		StepRunTime:      detail.ModeConf.StepRunTime,
		MaxConcurrency:   detail.ModeConf.MaxConcurrency,
		Duration:         detail.ModeConf.Duration,
		CreatedTimeSec:   ru.CreatedAt.Unix(),
	}

	res := &rao.ReportTask{
		UserID:         user.ID,
		UserName:       user.Nickname,
		UserAvatar:     user.Avatar,
		PlanID:         detail.PlanID,
		PlanName:       detail.PlanName,
		ReportID:       detail.ReportID,
		SceneID:        ru.SceneID,
		SceneName:      ru.SceneName,
		CreatedTimeSec: ru.CreatedAt.Unix(),
		TaskType:       detail.TaskType,
		TaskMode:       detail.TaskMode,
		TaskStatus:     ru.Status,
		ModeConf:       modeConf,
	}

	res.ChangeTakeConf = append(res.ChangeTakeConf, changeTaskConfData)

	if len(changeTaskConf) > 0 {
		for _, changeTaskConfTmp := range changeTaskConf {
			tmp := &rao.ModeConf{
				ReheatTime:       changeTaskConfTmp.ModeConf.ReheatTime,
				RoundNum:         changeTaskConfTmp.ModeConf.RoundNum,
				Concurrency:      changeTaskConfTmp.ModeConf.Concurrency,
				ThresholdValue:   changeTaskConfTmp.ModeConf.ThresholdValue,
				StartConcurrency: changeTaskConfTmp.ModeConf.StartConcurrency,
				Step:             changeTaskConfTmp.ModeConf.Step,
				StepRunTime:      changeTaskConfTmp.ModeConf.StepRunTime,
				MaxConcurrency:   changeTaskConfTmp.ModeConf.MaxConcurrency,
				Duration:         changeTaskConfTmp.ModeConf.Duration,
				CreatedTimeSec:   changeTaskConfTmp.ModeConf.CreatedTimeSec,
			}
			res.ChangeTakeConf = append(res.ChangeTakeConf, tmp)
		}
	}

	return res, nil
}

func GetReportDebugStatus(ctx context.Context, report rao.GetReportReq) string {
	reportId := int(report.ReportID)
	//reportId := report.ReportID
	collection := dal.GetMongo().Database(dal.MongoDB()).Collection(consts.CollectDebugStatus)
	filter := bson.D{{"report_id", reportId}}
	//fmt.Println("filter:", filter)
	cur := collection.FindOne(ctx, filter)
	result, err := cur.DecodeBytes()
	if err != nil {
		return consts.StopDebug
	}
	list, err := result.Elements()
	if err != nil {
		return consts.StopDebug
	}
	for _, value := range list {
		if value.Key() == "debug" {
			return value.Value().StringValue()
		}
	}
	return consts.StopDebug
}

func GetReportDebugLog(ctx context.Context, report rao.GetReportReq) (err error, debugMsgList []map[string]interface{}) {
	//clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s/%s", user, password, host, db))

	reportId := strconv.FormatInt(report.ReportID, 10)
	collection := dal.GetMongo().Database(dal.MongoDB()).Collection(consts.CollectStressDebug)
	filter := bson.D{{"report_id", reportId}}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		proof.Error("debug日志查询失败", proof.WithError(err))
		return
	}
	for cur.Next(ctx) {
		debugMsg := make(map[string]interface{})
		err = cur.Decode(&debugMsg)
		if err != nil {
			proof.Error("debug日志转换失败", proof.WithError(err))
			return
		}
		if debugMsg["end"] != true {
			debugMsgList = append(debugMsgList, debugMsg)
		}
	}
	return
}

// GetReportDetail 从redis获取测试数据
func GetReportDetail(ctx context.Context, report rao.GetReportReq) (err error, resultData ResultData) {
	collection := dal.GetMongo().Database(dal.MongoDB()).Collection(consts.CollectReportData)
	filter := bson.D{{"reportid", fmt.Sprintf("%d", report.ReportID)}}
	var resultMsg SceneTestResultDataMsg
	var dataMap = make(map[string]string)
	err = collection.FindOne(ctx, filter).Decode(dataMap)
	_, ok := dataMap["data"]
	if err != nil || !ok {
		rdb := dal.GetRDB()
		key := fmt.Sprintf("%d:%d:reportData", report.PlanId, report.ReportID)
		dataList := rdb.LRange(ctx, key, 0, -1).Val()
		if len(dataList) < 1 {
			proof.Error("mongo里面没有查到报告详情数据，err:", proof.WithError(err))
			err = nil
			return
		}
		for i := len(dataList) - 1; i >= 0; i-- {
			resultMsgString := dataList[i]
			err = json.Unmarshal([]byte(resultMsgString), &resultMsg)
			if err != nil {
				proof.Error("json转换格式错误：", proof.WithError(err))
			}
			if resultData.Results == nil {
				resultData.Results = make(map[string]*ResultDataMsg)
			}
			resultData.ReportId = resultMsg.ReportId
			resultData.End = resultMsg.End
			resultData.ReportName = resultMsg.ReportName
			resultData.PlanId = resultMsg.PlanId
			resultData.PlanName = resultMsg.PlanName
			resultData.SceneId = resultMsg.SceneId
			resultData.SceneName = resultMsg.SceneName
			resultData.TimeStamp = resultMsg.TimeStamp
			if resultMsg.Results != nil && len(resultMsg.Results) > 0 {
				for k, apiResult := range resultMsg.Results {
					if resultData.Results[k] == nil {
						resultData.Results[k] = new(ResultDataMsg)
					}
					resultData.Results[k].ApiName = apiResult.Name
					resultData.Results[k].Concurrency = apiResult.Concurrency
					resultData.Results[k].TotalRequestNum = apiResult.TotalRequestNum
					resultData.Results[k].TotalRequestTime, _ = decimal.NewFromFloat(float64(apiResult.TotalRequestTime) / float64(time.Second)).Round(2).Float64()
					resultData.Results[k].SuccessNum = apiResult.SuccessNum
					resultData.Results[k].ErrorNum = apiResult.ErrorNum
					if apiResult.TotalRequestNum != 0 {
						errRate := float64(apiResult.ErrorNum) / float64(apiResult.TotalRequestNum)
						resultData.Results[k].ErrorRate, _ = strconv.ParseFloat(fmt.Sprintf("%0.2f", errRate), 64)
					}
					resultData.Results[k].PercentAge = apiResult.PercentAge
					resultData.Results[k].ErrorThreshold = apiResult.ErrorThreshold
					resultData.Results[k].ResponseThreshold = apiResult.ResponseThreshold
					resultData.Results[k].RequestThreshold = apiResult.RequestThreshold
					resultData.Results[k].AvgRequestTime, _ = decimal.NewFromFloat(apiResult.AvgRequestTime / float64(time.Millisecond)).Round(1).Float64()
					resultData.Results[k].MaxRequestTime, _ = decimal.NewFromFloat(apiResult.MaxRequestTime / float64(time.Millisecond)).Round(1).Float64()
					resultData.Results[k].MinRequestTime, _ = decimal.NewFromFloat(apiResult.MinRequestTime / float64(time.Millisecond)).Round(1).Float64()
					resultData.Results[k].CustomRequestTimeLine = apiResult.CustomRequestTimeLine
					resultData.Results[k].CustomRequestTimeLineValue, _ = decimal.NewFromFloat(apiResult.CustomRequestTimeLineValue / float64(time.Millisecond)).Round(1).Float64()
					resultData.Results[k].FiftyRequestTimelineValue, _ = decimal.NewFromFloat(apiResult.FiftyRequestTimelineValue / float64(time.Millisecond)).Round(1).Float64()
					resultData.Results[k].NinetyRequestTimeLine = apiResult.NinetyRequestTimeLine
					resultData.Results[k].NinetyRequestTimeLineValue, _ = decimal.NewFromFloat(apiResult.NinetyRequestTimeLineValue / float64(time.Millisecond)).Round(1).Float64()
					resultData.Results[k].NinetyFiveRequestTimeLine = apiResult.NinetyFiveRequestTimeLine
					resultData.Results[k].NinetyFiveRequestTimeLineValue, _ = decimal.NewFromFloat(apiResult.NinetyFiveRequestTimeLineValue / float64(time.Millisecond)).Round(1).Float64()
					resultData.Results[k].NinetyNineRequestTimeLine = apiResult.NinetyNineRequestTimeLine
					resultData.Results[k].NinetyNineRequestTimeLineValue, _ = decimal.NewFromFloat(apiResult.NinetyNineRequestTimeLineValue / float64(time.Millisecond)).Round(1).Float64()
					resultData.Results[k].SendBytes, _ = decimal.NewFromFloat(apiResult.SendBytes).Round(1).Float64()
					resultData.Results[k].ReceivedBytes, _ = decimal.NewFromFloat(apiResult.ReceivedBytes).Round(1).Float64()
					resultData.Results[k].Qps = apiResult.Qps
					resultData.Results[k].SRps = apiResult.SRps
					if resultData.Results[k].QpsList == nil {
						resultData.Results[k].QpsList = []TimeValue{}
					}
					var timeValue = TimeValue{}
					timeValue.TimeStamp = resultData.TimeStamp
					// qps列表
					timeValue.Value = resultData.Results[k].Qps
					resultData.Results[k].QpsList = append(resultData.Results[k].QpsList, timeValue)
					timeValue.Value = resultData.Results[k].ErrorNum
					if resultData.Results[k].ErrorNumList == nil {
						resultData.Results[k].ErrorNumList = []TimeValue{}
					}
					// 错误数列表
					resultData.Results[k].ErrorNumList = append(resultData.Results[k].ErrorNumList, timeValue)
					timeValue.Value = resultData.Results[k].Concurrency
					if resultData.Results[k].ConcurrencyList == nil {
						resultData.Results[k].ConcurrencyList = []TimeValue{}
					}
					// 并发数列表
					resultData.Results[k].ConcurrencyList = append(resultData.Results[k].ConcurrencyList, timeValue)

					// 平均响应时间列表
					timeValue.Value = resultData.Results[k].AvgRequestTime
					if resultData.Results[k].AvgList == nil {
						resultData.Results[k].AvgList = []TimeValue{}
					}
					resultData.Results[k].AvgList = append(resultData.Results[k].AvgList, timeValue)

					// 50响应时间列表
					timeValue.Value = resultData.Results[k].FiftyRequestTimelineValue
					if resultData.Results[k].FiftyList == nil {
						resultData.Results[k].FiftyList = []TimeValue{}
					}
					resultData.Results[k].FiftyList = append(resultData.Results[k].FiftyList, timeValue)

					// 90响应时间列表
					timeValue.Value = resultData.Results[k].NinetyNineRequestTimeLineValue
					if resultData.Results[k].NinetyList == nil {
						resultData.Results[k].NinetyList = []TimeValue{}
					}
					resultData.Results[k].NinetyList = append(resultData.Results[k].NinetyList, timeValue)

					// 95响应时间列表
					timeValue.Value = resultData.Results[k].NinetyFiveRequestTimeLineValue
					if resultData.Results[k].NinetyFiveList == nil {
						resultData.Results[k].NinetyFiveList = []TimeValue{}
					}
					resultData.Results[k].NinetyFiveList = append(resultData.Results[k].NinetyFiveList, timeValue)

					// 99响应时间列表
					timeValue.Value = resultData.Results[k].NinetyNineRequestTimeLineValue
					if resultData.Results[k].NinetyNineList == nil {
						resultData.Results[k].NinetyNineList = []TimeValue{}
					}
					resultData.Results[k].NinetyNineList = append(resultData.Results[k].NinetyNineList, timeValue)
				}
			}
			if resultMsg.End {
				var by []byte
				by, err = json.Marshal(resultData)
				if err != nil {
					proof.Error("resultData转字节失败：：    ", proof.WithError(err))
					return
				}
				var apiResultTotalMsg = make(map[string]string)
				for _, value := range resultData.Results {
					apiResultTotalMsg[value.ApiName] = fmt.Sprintf("平均响应时间为%0.1fms； 百分之五十响应时间线的值为%0.1fms; 百分之九十响应时间线的值为%0.1fms; 百分之九十五响应时间线的值为%0.1fms; 百分之九十九响应时间线的值为%0.1fms; RPS为%0.1f; SRPS为%0.1f",
						value.AvgRequestTime, value.FiftyRequestTimelineValue, value.NinetyRequestTimeLineValue, value.NinetyFiveRequestTimeLineValue, value.NinetyNineRequestTimeLineValue, value.Qps, value.SRps)
				}
				dataMap["reportid"] = resultData.ReportId
				dataMap["data"] = string(by)
				by, _ = json.Marshal(apiResultTotalMsg)
				dataMap["analysis"] = string(by)
				_, err = collection.InsertOne(ctx, dataMap)
				if err != nil {
					proof.Error("测试数据写入mongo失败：    ", proof.WithError(err))
					return
				}
				err = rdb.Del(ctx, key).Err()
				if err != nil {
					proof.Error(fmt.Sprintf("删除redis的key：%s:    ", key), proof.WithError(err))
					return
				}
			}
		}
	} else {
		data := dataMap["data"]
		err = json.Unmarshal([]byte(data), &resultData)
		return
	}
	err = nil
	return
}

type SceneTestResultDataMsg struct {
	End        bool                             `json:"end" bson:"end"`
	ReportId   string                           `json:"report_id" bson:"report_id"`
	ReportName string                           `json:"report_name" bson:"report_name"`
	PlanId     int64                            `json:"plan_id" bson:"plan_id"`     // 任务ID
	PlanName   string                           `json:"plan_name" bson:"plan_name"` //
	SceneId    int64                            `json:"scene_id" bson:"scene_id"`   // 场景
	SceneName  string                           `json:"scene_name" bson:"scene_name"`
	Results    map[string]*ApiTestResultDataMsg `json:"results" bson:"results"`
	Machine    map[string]int64                 `json:"machine" bson:"machine"`
	TimeStamp  int64                            `json:"time_stamp" bson:"time_stamp"`
}

// ApiTestResultDataMsg 接口测试数据经过计算后的测试结果
type ApiTestResultDataMsg struct {
	Name                           string  `json:"name" bson:"name"`
	Concurrency                    int64   `json:"concurrency" bson:"concurrency"`
	TotalRequestNum                uint64  `json:"total_request_num" bson:"total_request_num"`   // 总请求数
	TotalRequestTime               uint64  `json:"total_request_time" bson:"total_request_time"` // 总响应时间
	SuccessNum                     uint64  `json:"success_num" bson:"success_num"`
	ErrorNum                       uint64  `json:"error_num" bson:"error_num"`                   // 错误数
	ErrorThreshold                 float64 `json:"error_threshold" bson:"error_threshold"`       // 自定义错误率
	RequestThreshold               int64   `json:"request_threshold" bson:"request_threshold"`   // Rps（每秒请求数）阈值
	ResponseThreshold              int64   `json:"response_threshold" bson:"response_threshold"` // 响应时间阈值
	PercentAge                     int64   `json:"percent_age" bson:"percent_age"`               // 响应时间线
	AvgRequestTime                 float64 `json:"avg_request_time" bson:"avg_request_time"`     // 平均响应事件
	MaxRequestTime                 float64 `json:"max_request_time" bson:"max_request_time"`
	MinRequestTime                 float64 `json:"min_request_time" bson:"min_request_time"` // 毫秒
	CustomRequestTimeLine          int64   `json:"custom_request_time_line" bson:"custom_request_time_line"`
	FiftyRequestTimeline           int64   `json:"fifty_request_time_line" bson:"fifty_request_time_line"`
	NinetyRequestTimeLine          int64   `json:"ninety_request_time_line" bson:"ninety_request_time_line"`
	NinetyFiveRequestTimeLine      int64   `json:"ninety_five_request_time_line" bson:"ninety_five_request_time_line"`
	NinetyNineRequestTimeLine      int64   `json:"ninety_nine_request_time_line" bson:"ninety_nine_request_time_line"`
	FiftyRequestTimelineValue      float64 `json:"fifty_request_time_line_value"`
	CustomRequestTimeLineValue     float64 `json:"custom_request_time_line_value" bson:"custom_request_time_line_value"`
	NinetyRequestTimeLineValue     float64 `json:"ninety_request_time_line_value" bson:"ninety_request_time_line_value"`
	NinetyFiveRequestTimeLineValue float64 `json:"ninety_five_request_time_line_value" bson:"ninety_five_request_time_line_value"`
	NinetyNineRequestTimeLineValue float64 `json:"ninety_nine_request_time_line_value" bson:"ninety_nine_request_time_line_value"`
	SendBytes                      float64 `json:"send_bytes" bson:"send_bytes"`         // 发送字节数
	ReceivedBytes                  float64 `json:"received_bytes" bson:"received_bytes"` // 接收字节数
	Qps                            float64 `json:"qps" bson:"qps"`
	SRps                           float64 `json:"srps" bson:"srps"`
}

// ResultDataMsg 前端展示各个api数据
type ResultDataMsg struct {
	ApiName                        string      `json:"api_name" bson:"api_name"`
	Concurrency                    int64       `json:"concurrency" bson:"concurrency"`
	TotalRequestNum                uint64      `json:"total_request_num" bson:"total_request_num"`   // 总请求数
	TotalRequestTime               float64     `json:"total_request_time" bson:"total_request_time"` // 总响应时间
	SuccessNum                     uint64      `json:"success_num" bson:"success_num"`
	ErrorRate                      float64     `json:"error_rate" bson:"error_rate"`
	ErrorNum                       uint64      `json:"error_num" bson:"error_num"`               // 错误数
	AvgRequestTime                 float64     `json:"avg_request_time" bson:"avg_request_time"` // 平均响应事件
	MaxRequestTime                 float64     `json:"max_request_time" bson:"max_request_time"`
	MinRequestTime                 float64     `json:"min_request_time" bson:"min_request_time"`     // 毫秒
	PercentAge                     int64       `json:"percent_age" bson:"percent_age"`               // 响应时间线
	ErrorThreshold                 float64     `json:"error_threshold" bson:"error_threshold"`       // 自定义错误率
	RequestThreshold               int64       `json:"request_threshold" bson:"request_threshold"`   // Rps（每秒请求数）阈值
	ResponseThreshold              int64       `json:"response_threshold" bson:"response_threshold"` // 响应时间阈值
	CustomRequestTimeLine          int64       `json:"custom_request_time_line" bson:"custom_request_time_line"`
	FiftyRequestTimeline           int64       `json:"fifty_request_time_line" bson:"fifty_request_time_line"`
	NinetyRequestTimeLine          int64       `json:"ninety_request_time_line" bson:"ninety_request_time_line"`
	NinetyFiveRequestTimeLine      int64       `json:"ninety_five_request_time_line" bson:"ninety_five_request_time_line"`
	NinetyNineRequestTimeLine      int64       `json:"ninety_nine_request_time_line" bson:"ninety_nine_request_time_line"`
	CustomRequestTimeLineValue     float64     `json:"custom_request_time_line_value" bson:"custom_request_time_line_value"`
	FiftyRequestTimelineValue      float64     `json:"fifty_request_time_line_value" bson:"fifty_request_time_line_value"`
	NinetyRequestTimeLineValue     float64     `json:"ninety_request_time_line_value" bson:"ninety_request_time_line_value"`
	NinetyFiveRequestTimeLineValue float64     `json:"ninety_five_request_time_line_value" bson:"ninety_five_request_time_line_value"`
	NinetyNineRequestTimeLineValue float64     `json:"ninety_nine_request_time_line_value" bson:"ninety_nine_request_time_line_value"`
	SendBytes                      float64     `json:"send_bytes" bson:"send_bytes"`         // 发送字节数
	ReceivedBytes                  float64     `json:"received_bytes" bson:"received_bytes"` // 接收字节数
	Qps                            float64     `json:"qps" bson:"qps"`
	SRps                           float64     `json:"srps" bson:"srps"`
	ConcurrencyList                []TimeValue `json:"concurrency_list" bson:"concurrency_list"`
	QpsList                        []TimeValue `json:"qps_list" bson:"qps_list"`
	ErrorNumList                   []TimeValue `json:"error_num_list" bson:"error_num_list"`
	AvgList                        []TimeValue `json:"avg_list" bson:"avg_list"`
	FiftyList                      []TimeValue `json:"fifty_list" bson:"fifty_list"`
	NinetyList                     []TimeValue `json:"ninety_list" bson:"ninety_list"`
	NinetyFiveList                 []TimeValue `json:"ninety_five_list" bson:"ninety_five_list"`
	NinetyNineList                 []TimeValue `json:"ninety_nine_list" bson:"ninety_nine_list"`
}

type ResultData struct {
	End        bool                      `json:"end"`
	ReportId   string                    `json:"report_id"`
	ReportName string                    `json:"report_name"`
	PlanId     int64                     `json:"plan_id"`   // 任务ID
	PlanName   string                    `json:"plan_name"` //
	SceneId    int64                     `json:"scene_id"`  // 场景
	SceneName  string                    `json:"scene_name"`
	Results    map[string]*ResultDataMsg `json:"results"`
	TimeStamp  int64                     `json:"time_stamp"`
	Analysis   string                    `json:"analysis"`
	Msg        string                    `json:"msg"`
}

type TimeValue struct {
	TimeStamp int64       `json:"time_stamp" bson:"time_stamp"`
	Value     interface{} `json:"value" bson:"value"`
}

func GetCompareReportData(ctx context.Context, req rao.CompareReportReq) (*CompareReportResponse, error) {
	// 获取报告的基本信息
	reportTable := dal.GetQuery().Report
	reportBaseList, err := reportTable.WithContext(ctx).Where(reportTable.ID.In(req.ReportIDs...)).Find()
	if err != nil {
		return nil, err
	}

	reportNames := make([]string, 0, len(reportBaseList))                    // 计划和场景名字
	reportBaseData := make([]*mao.ReportTask, 0, len(reportBaseList))        // 报告基本信息
	reportCollectMap := make([][]*reportCollectData, 0, len(reportBaseList)) // 报告汇总信息

	//reportDetailMap := make([]*reportDetailData, 0, 10)
	reportDetailAllMap := make([][]*reportDetailData, 0, len(reportBaseList)) // 报告详情信息

	for _, reportBaseInfo := range reportBaseList {
		// 把报告基本信息设置到map当中
		planAndSceneName := reportBaseInfo.PlanName + "/" + reportBaseInfo.SceneName
		reportNames = append(reportNames, planAndSceneName)

	}

	// 从mg查询任务对应的配置信息
	var reportTaskConfList []*mao.ReportTask
	collection := dal.GetMongo().Database(dal.MongoDB()).Collection(consts.CollectReportTask)
	reportTaskConfListTmp, err := collection.Find(ctx, bson.D{{"report_id", bson.D{{"$in", req.ReportIDs}}}})
	if err != nil {
		// todo
		return nil, err
	}
	if err := reportTaskConfListTmp.All(ctx, &reportTaskConfList); err != nil {
		// todo
		return nil, err
	}

	for _, reportTaskConfInfo := range reportTaskConfList {
		reportBaseData = append(reportBaseData, reportTaskConfInfo)
	}

	// 从mg里面获取报告汇总信息
	var reportDataList []*ResultData
	collection = dal.GetMongo().Database(dal.MongoDB()).Collection(consts.CollectReportData)
	reportDataListTmp, err := collection.Find(ctx, bson.D{{"reportid", bson.D{{"$in", req.ReportIDs}}}})
	if err != nil {
		// todo
		return nil, err
	}
	if err := reportDataListTmp.All(ctx, &reportDataList); err != nil {
		// todo
		return nil, err
	}

	// 先组装一个事件id对应报告结果数据
	for _, reportData := range reportDataList {
		//rId, err := strconv.Atoi(reportData.ReportId)
		//if err != nil {
		//	continue
		//}

		var reportCollectMapTmp []*reportCollectData
		var reportDetailDataSlice []*reportDetailData
		for _, reportMsg := range reportData.Results {
			sceneNodeMap := &reportCollectData{
				ApiName:                   reportMsg.ApiName,
				TotalRequestNum:           reportMsg.TotalRequestNum,
				TotalRequestTime:          reportMsg.TotalRequestTime,
				MaxRequestTime:            reportMsg.MaxRequestTime,
				MinRequestTime:            reportMsg.MinRequestTime,
				AvgRequestTime:            reportMsg.AvgRequestTime,
				NinetyRequestTimeLine:     reportMsg.NinetyRequestTimeLine,
				NinetyFiveRequestTimeLine: reportMsg.NinetyFiveRequestTimeLine,
				NinetyNineRequestTimeLine: reportMsg.NinetyNineRequestTimeLine,
				Qps:                       reportMsg.Qps,
				SRps:                      reportMsg.SRps,
				ErrorRate:                 reportMsg.ErrorRate,
				ReceivedBytes:             reportMsg.ReceivedBytes,
				SendBytes:                 reportMsg.SendBytes,
			}
			reportCollectMapTmp = append(reportCollectMapTmp, sceneNodeMap)

			// 接口请求详情数据
			reportDetailDataTmp := &reportDetailData{
				AvgList:         reportMsg.AvgList,
				QpsList:         reportMsg.QpsList,
				ConcurrencyList: reportMsg.ConcurrencyList,
				ErrorNumList:    reportMsg.ErrorNumList,
				FiftyList:       reportMsg.FiftyList,
				NinetyList:      reportMsg.NinetyList,
				NinetyFiveList:  reportMsg.NinetyFiveList,
				NinetyNineList:  reportMsg.NinetyNineList,
			}
			reportDetailDataSlice = append(reportDetailDataSlice, reportDetailDataTmp)
		}

		reportCollectMap = append(reportCollectMap, reportCollectMapTmp)
		reportDetailAllMap = append(reportDetailAllMap, reportDetailDataSlice)
	}

	res := &CompareReportResponse{
		ReportNamesData:     reportNames,
		ReportBaseData:      reportBaseData,
		ReportCollectData:   reportCollectMap,
		ReportDetailAllData: reportDetailAllMap,
	}

	return res, nil
}

type reportBaseValue struct {
	ReportID      int64  `json:"report_id"`
	Name          string `json:"name"`
	RunUserID     int64  `json:"run_user_id"`
	Performer     string `json:"performer"`
	CreateTimeSec int64  `json:"create_time_sec"`
	TaskType      int32  `json:"task_type"` // 任务类型
	TaskMode      int32  `json:"task_mode"` // 压测模式
	rao.ModeConf
}

type reportCollectAllData struct {
	ReportId          int64                `json:"report_id"`
	PlanAndScene      string               `json:"plan_and_scene"` // 计划和场景名称
	ReportCollectData []*reportCollectData `json:"report_collect_data"`
}
type reportCollectData struct {
	ApiName                   string  `json:"api_name" bson:"api_name"`
	TotalRequestNum           uint64  `json:"total_request_num" bson:"total_request_num"`   // 总请求数
	TotalRequestTime          float64 `json:"total_request_time" bson:"total_request_time"` // 总响应时间
	MaxRequestTime            float64 `json:"max_request_time" bson:"max_request_time"`
	MinRequestTime            float64 `json:"min_request_time" bson:"min_request_time"` // 毫秒
	AvgRequestTime            float64 `json:"avg_request_time" bson:"avg_request_time"` // 平均响应事件
	NinetyRequestTimeLine     int64   `json:"ninety_request_time_line" bson:"ninety_request_time_line"`
	NinetyFiveRequestTimeLine int64   `json:"ninety_five_request_time_line" bson:"ninety_five_request_time_line"`
	NinetyNineRequestTimeLine int64   `json:"ninety_nine_request_time_line" bson:"ninety_nine_request_time_line"`
	Qps                       float64 `json:"qps" bson:"qps"`
	SRps                      float64 `json:"srps" bson:"srps"`
	ErrorRate                 float64 `json:"error_rate" bson:"error_rate"`
	ReceivedBytes             float64 `json:"received_bytes" bson:"received_bytes"` // 接收字节数
	SendBytes                 float64 `json:"send_bytes" bson:"send_bytes"`         // 发送字节数
}

type reportDetailData struct {
	AvgList         []TimeValue `json:"avg_list" bson:"avg_list"`
	QpsList         []TimeValue `json:"qps_list" bson:"qps_list"`
	ConcurrencyList []TimeValue `json:"concurrency_list" bson:"concurrency_list"`
	ErrorNumList    []TimeValue `json:"error_num_list" bson:"error_num_list"`
	FiftyList       []TimeValue `json:"fifty_list" bson:"fifty_list"`
	NinetyList      []TimeValue `json:"ninety_list" bson:"ninety_list"`
	NinetyFiveList  []TimeValue `json:"ninety_five_list" bson:"ninety_five_list"`
	NinetyNineList  []TimeValue `json:"ninety_nine_list" bson:"ninety_nine_list"`
}

// 对比报告接口返回值
type CompareReportResponse struct {
	ReportNamesData     []string               `json:"report_names_data"`
	ReportBaseData      []*mao.ReportTask      `json:"report_base_data"`
	ReportCollectData   [][]*reportCollectData `json:"report_collect_data"`
	ReportDetailAllData [][]*reportDetailData  `json:"report_detail_all_data"`
}
