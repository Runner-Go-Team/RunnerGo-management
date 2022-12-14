package handler

import (
	"context"
	"github.com/go-omnibus/omnibus"
	"github.com/go-omnibus/proof"
	"strconv"
	"strings"
	"time"

	"kp-management/internal/pkg/biz/consts"
	"kp-management/internal/pkg/biz/errno"
	"kp-management/internal/pkg/biz/jwt"
	"kp-management/internal/pkg/biz/mail"
	"kp-management/internal/pkg/biz/record"
	"kp-management/internal/pkg/biz/response"
	"kp-management/internal/pkg/dal"
	"kp-management/internal/pkg/dal/model"
	"kp-management/internal/pkg/dal/rao"
	"kp-management/internal/pkg/logic/plan"
	"kp-management/internal/pkg/logic/stress"

	"github.com/gin-gonic/gin"

	"kp-management/internal/pkg/dal/query"
)

type RunStressReq struct {
	PlanID  int64   `json:"plan_id"`
	TeamID  int64   `json:"team_id"`
	SceneID []int64 `json:"scene_id"`
	UserID  int64   `json:"user_id"`
}

// RunPlan 启动计划
func RunPlan(ctx *gin.Context) {
	var req rao.RunPlanReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		return
	}

	// 调用controller方法改成本地
	runStressParams := RunStressReq{
		PlanID:  req.PlanID,
		TeamID:  req.TeamID,
		SceneID: req.SceneID,
		UserID:  jwt.GetUserIDByCtx(ctx),
	}

	errnoNum, runErr := RunStress(ctx, runStressParams)
	if runErr != nil {
		response.ErrorWithMsg(ctx, errnoNum, runErr.Error())
		return
	}

	px := dal.GetQuery().Plan
	p, err := px.WithContext(ctx).Where(px.ID.Eq(req.PlanID)).First()
	if err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	if err := record.InsertRun(ctx, req.TeamID, jwt.GetUserIDByCtx(ctx), record.OperationOperateRunPlan, p.Name); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	tx := dal.GetQuery().PlanEmail
	emails, err := tx.WithContext(ctx).Where(tx.PlanID.Eq(req.PlanID)).Find()
	if err != nil {
		response.ErrorWithMsg(ctx, errno.ErrHttpFailed, err.Error())
		return
	}

	if len(emails) > 0 {
		px := dal.GetQuery().Plan
		planInfo, err := px.WithContext(ctx).Where(px.ID.Eq(req.PlanID)).First()
		if err != nil {
			response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
			return
		}

		ttx := dal.GetQuery().Team
		team, err := ttx.WithContext(ctx).Where(ttx.ID.Eq(req.TeamID)).First()
		if err != nil {
			response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
			return
		}

		rx := dal.GetQuery().Report
		reports, err := rx.WithContext(ctx).Where(rx.PlanID.Eq(req.PlanID), rx.CreatedAt.Gt(emails[0].CreatedAt)).Find()
		if err != nil {
			response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
			return
		}

		ux := dal.GetQuery().User
		user, err := ux.WithContext(ctx).Where(ux.ID.Eq(jwt.GetUserIDByCtx(ctx))).First()
		if err != nil {
			response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
			return
		}

		var userIDs []int64
		for _, report := range reports {
			userIDs = append(userIDs, report.RunUserID)
		}
		runUsers, err := ux.WithContext(ctx).Where(ux.ID.In(userIDs...)).Find()
		if err != nil {
			response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
			return
		}

		for _, email := range emails {
			if err := mail.SendPlanEmail(ctx, email.Email, planInfo.Name, team.Name, user.Nickname, reports, runUsers); err != nil {
				response.ErrorWithMsg(ctx, errno.ErrHttpFailed, err.Error())
				return
			}
		}
	}

	response.Success(ctx)
	return
}

// StopPlan 停止计划
func StopPlan(ctx *gin.Context) {
	var req rao.StopPlanReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		return
	}

	tx := dal.GetQuery().Report
	reports, err := tx.WithContext(ctx).Where(tx.PlanID.In(req.PlanIDs...)).Find()
	if err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	var reportIDs []int64
	for _, report := range reports {
		reportIDs = append(reportIDs, report.ID)
	}

	// 停止计划的时候，往redis里面写一条数据
	reportIDsString := omnibus.Int64sToStrings(reportIDs)
	teamIDString := strconv.Itoa(int(req.TeamID))
	planIDString := strconv.Itoa(int(req.PlanIDs[0]))
	for _, reportID := range reportIDsString {
		stopPlanKey := consts.StopPlanPrefix + teamIDString + ":" + planIDString + ":" + reportID
		_, err := dal.GetRDB().Set(ctx, stopPlanKey, "stop", 0).Result()
		if err != nil {
			proof.Errorf("停止计划--写入redis数据失败，err:", err)
			response.ErrorWithMsg(ctx, errno.ErrRedisFailed, err.Error())
			return
		}
	}

	//_, err = resty.New().R().
	//	SetBody(runner.StopRunnerReq{ReportIds: omnibus.Int64sToStrings(reportIDs), TeamID: req.TeamID, PlanID: req.PlanIDs[0]}).
	//	Post(conf.Conf.Clients.Runner.StopPlan)
	//
	//if err != nil {
	//	response.ErrorWithMsg(ctx, errno.ErrHttpFailed, err.Error())
	//	return
	//}

	px := dal.GetQuery().Plan
	_, err = px.WithContext(ctx).Where(px.ID.In(req.PlanIDs...)).UpdateColumn(px.Status, consts.PlanStatusNormal)
	if err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	//_, err = tx.WithContext(ctx).Where(tx.PlanID.In(req.PlanIDs...)).UpdateColumn(tx.Status, consts.ReportStatusFinish)
	//if err != nil {
	//	response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
	//	return
	//}

	response.Success(ctx)
	return
}

// ClonePlan 克隆计划
func ClonePlan(ctx *gin.Context) {
	var req rao.ClonePlanReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		return
	}

	if err := plan.ClonePlan(ctx, req.PlanID, req.TeamID, jwt.GetUserIDByCtx(ctx)); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	response.Success(ctx)
	return
}

// ListUnderwayPlan 运行中的计划
func ListUnderwayPlan(ctx *gin.Context) {
	var req rao.ListUnderwayPlanReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		return
	}

	plans, total, err := plan.ListByStatus(ctx, req.TeamID, consts.PlanStatusUnderway, req.Size, (req.Page-1)*req.Size)
	if err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	response.SuccessWithData(ctx, rao.ListUnderwayPlanResp{
		Plans: plans,
		Total: total,
	})
	return
}

// ListPlans 测试计划列表
func ListPlans(ctx *gin.Context) {
	var req rao.ListPlansReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		return
	}

	plans, total, err := plan.ListByTeamID(ctx, req.TeamID, req.Size, (req.Page-1)*req.Size,
		req.Keyword, req.StartTimeSec, req.EndTimeSec, req.TaskType, req.TaskMode, req.Status, req.Sort)
	if err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	response.SuccessWithData(ctx, rao.ListPlansResp{
		Plans: plans,
		Total: total,
	})
	return
}

// SavePlan 创建修改计划
func SavePlan(ctx *gin.Context) {
	var req rao.SavePlanReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		return
	}

	planID, err := plan.Save(ctx, &req, jwt.GetUserIDByCtx(ctx))
	if err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	response.SuccessWithData(ctx, rao.SavePlanResp{PlanID: planID})
	return
}

// SavePlanTask 创建/修改计划配置
func SavePlanTask(ctx *gin.Context) {
	var req rao.SavePlanConfReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		return
	}

	// 必填项判断
	if (req.ModeConf.Duration == 0 && req.ModeConf.RoundNum == 0) || (req.Mode == 1 && req.ModeConf.Concurrency == 0) {
		response.ErrorWithMsg(ctx, errno.ErrMustTaskInit, "必填项不能为空")
		return
	}

	if req.Mode != 1 { // 非并发模式参数校验
		if req.ModeConf.StartConcurrency == 0 || req.ModeConf.Step == 0 || req.ModeConf.StepRunTime == 0 || req.ModeConf.MaxConcurrency == 0 || req.ModeConf.Duration == 0 {
			response.ErrorWithMsg(ctx, errno.ErrMustTaskInit, "必填项不能为空")
			return
		}
	}

	if req.TaskType == 2 {
		if req.TimedTaskConf.Frequency == 0 {
			if req.TimedTaskConf.TaskExecTime == 0 {
				response.ErrorWithMsg(ctx, errno.ErrMustTaskInit, "必填项不能为空")
				return
			}
		} else {
			if req.TimedTaskConf.TaskExecTime == 0 || req.TimedTaskConf.TaskCloseTime == 0 {
				response.ErrorWithMsg(ctx, errno.ErrMustTaskInit, "必填项不能为空")
				return
			}
		}
	}

	if err := plan.SaveTask(ctx, &req, jwt.GetUserIDByCtx(ctx)); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	response.Success(ctx)
	return
}

func GetPlanTask(ctx *gin.Context) {
	var req rao.GetPlanTaskReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		return
	}

	pt, err := plan.GetPlanTask(ctx, req.PlanID, req.SceneID)
	if err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	response.SuccessWithData(ctx, rao.GetPlanTaskResp{PlanTask: pt})
	return
}

// GetPlan 获取计划
func GetPlan(ctx *gin.Context) {
	var req rao.GetPlanConfReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		return
	}

	p, err := plan.GetByPlanID(ctx, req.TeamID, req.PlanID)
	if err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	response.SuccessWithData(ctx, rao.GetPlanResp{Plan: p})
	return
}

// DeletePlan 删除计划
func DeletePlan(ctx *gin.Context) {
	var req rao.DeletePlanReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		return
	}

	if err := plan.DeleteByPlanID(ctx, req.TeamID, req.PlanID, jwt.GetUserIDByCtx(ctx)); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	response.Success(ctx)
	return
}

// ImportScene 导入场景
func ImportScene(ctx *gin.Context) {
	var req rao.ImportSceneReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		return
	}

	scenes, err := plan.ImportScene(ctx, jwt.GetUserIDByCtx(ctx), req.PlanID, req.TargetIDList)
	if err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	response.SuccessWithData(ctx, rao.ImportSceneResp{
		Scenes: scenes,
	})
	return
}

// PlanAddEmail 添加计划收件人
func PlanAddEmail(ctx *gin.Context) {
	var req rao.PlanEmailReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		return
	}

	var planEmails []*model.PlanEmail
	for _, email := range req.Emails {
		planEmails = append(planEmails, &model.PlanEmail{
			PlanID: req.PlanID,
			Email:  email,
		})
	}

	tx := dal.GetQuery().PlanEmail
	cnt, err := tx.WithContext(ctx).Where(tx.PlanID.Eq(req.PlanID), tx.Email.In(req.Emails...)).Count()
	if err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	if cnt > 0 {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, "email exists")
		return
	}

	if err := tx.WithContext(ctx).CreateInBatches(planEmails, 5); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	response.Success(ctx)
	return
}

// PlanListEmail 计划收件人列表
func PlanListEmail(ctx *gin.Context) {
	var req rao.PlanListEmailReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		return
	}

	tx := dal.GetQuery().PlanEmail
	emails, err := tx.WithContext(ctx).Where(tx.PlanID.Eq(req.PlanID)).Order(tx.CreatedAt).Find()
	if err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	ret := make([]*rao.PlanEmail, 0)
	for _, e := range emails {
		ret = append(ret, &rao.PlanEmail{
			ID:            e.ID,
			PlanID:        e.PlanID,
			Email:         e.Email,
			CreateTimeSec: e.CreatedAt.Unix(),
		})
	}

	response.SuccessWithData(ctx, rao.PlanListEmailResp{Emails: ret})
	return
}

// PlanDeleteEmail 删除计划收件人
func PlanDeleteEmail(ctx *gin.Context) {
	var req rao.PlanDeleteEmailReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		return
	}

	tx := dal.GetQuery().PlanEmail
	_, err := tx.WithContext(ctx).Where(tx.ID.Eq(req.EmailID)).Delete()
	if err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	response.Success(ctx)
	return
}

// RunStress 调度压力测试机进行压测的方法
func RunStress(ctx context.Context, req RunStressReq) (int, error) {
	rms := &stress.RunMachineStress{}

	//siv := &stress.SplitImportVariable{}
	//siv.SetNext(rms)

	ss := &stress.SplitStress{}
	ss.SetNext(rms)

	ms := &stress.MakeStress{}
	ms.SetNext(ss)

	mr := &stress.MakeReport{}
	mr.SetNext(ms)

	iv := &stress.AssembleImportVariables{}
	iv.SetNext(mr)

	sv := &stress.AssembleSceneVariables{}
	sv.SetNext(iv)

	f := &stress.AssembleFlows{}
	f.SetNext(sv)

	v := &stress.AssembleGlobalVariables{}
	v.SetNext(f)

	t := &stress.AssembleTask{}
	t.SetNext(v)

	s := &stress.AssembleScenes{}
	s.SetNext(t)

	p := &stress.AssemblePlan{}
	p.SetNext(s)

	m := &stress.CheckIdleMachine{}
	m.SetNext(p)

	errnoNum, err := m.Execute(&stress.Baton{
		Ctx:      ctx,
		PlanID:   req.PlanID,
		TeamID:   req.TeamID,
		SceneIDs: req.SceneID,
		UserID:   req.UserID,
	})

	return errnoNum, err
}

type notifyStopStressReq struct {
	ReportID int64    `json:"report_id"`
	Machines []string `json:"machines"`
}

// NotifyStopStress 压力机回调压测状态和结果
func NotifyStopStress(ctx *gin.Context) {
	var req notifyStopStressReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		return
	}

	reportInfo := new(model.Report)

	err := query.Use(dal.DB()).Transaction(func(tx *query.Query) error {
		r := tx.Report
		// 修改报告状态
		_, err := r.WithContext(ctx).Where(r.ID.Eq(req.ReportID)).UpdateSimple(r.Status.Value(consts.ReportStatusFinish), r.UpdatedAt.Value(time.Now()))
		if err != nil {
			proof.Errorf("NotifyStopStress--修改报告状态失败")
			return err
		}

		// 查找报告对应计划
		report, err := r.WithContext(ctx).Where(r.ID.Eq(req.ReportID)).First()
		if err != nil {
			proof.Errorf("NotifyStopStress--查找报告对应计划失败")
			return err
		}
		reportInfo = report

		// 统计报告是否全部完成
		reportCnt, err := r.WithContext(ctx).Where(r.PlanID.Eq(report.PlanID)).Count()
		if err != nil {
			proof.Errorf("NotifyStopStress--统计当前计划下所有的报告数量--失败")
			return err
		}
		finishReportCnt, err := r.WithContext(ctx).Where(r.PlanID.Eq(report.PlanID), r.Status.Eq(consts.ReportStatusFinish)).Count()
		if err != nil {
			proof.Errorf("NotifyStopStress--统计当前计划下所有成功的报告--失败")
			return err
		}

		if finishReportCnt == reportCnt { // 报告全部完成则计划也完成
			p := tx.Plan
			_, err := p.WithContext(ctx).Where(p.ID.Eq(report.PlanID)).UpdateSimple(p.Status.Value(consts.PlanStatusNormal), p.UpdatedAt.Value(time.Now()))
			if err != nil {
				proof.Infof("NotifyStopStress计划下所有报告并未全部完成")
				return err
			}
		}

		return nil
	})

	for _, machine := range req.Machines {
		mInfo := strings.Split(machine, "_")
		if len(mInfo) != 2 {
			continue
		}
		machineUseStateKey := consts.MachineUseStatePrefix + mInfo[0] + ":" + mInfo[1]
		dal.RDB.Del(machineUseStateKey)
	}

	if err != nil {
		proof.Errorf("NotifyStopStress整体事务失败")
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	} else {
		// 判断当前计划下是否有定时任务
		ttc := dal.GetQuery().TimedTaskConf
		TimedTaskConfInfo, err2 := ttc.WithContext(ctx).Where(ttc.SenceID.Eq(reportInfo.SceneID)).First()
		if err2 == nil { // 查到定时任务了
			if TimedTaskConfInfo.Status == 1 { // 定时任务还没有过期
				p := dal.GetQuery().Plan
				_, err := p.WithContext(ctx).Where(p.ID.Eq(reportInfo.PlanID)).UpdateSimple(p.Status.Value(consts.PlanStatusUnderway), p.UpdatedAt.Value(time.Now()))
				if err != nil {
					proof.Infof("NotifyStopStress计划下--有定时任务没有完成")
					response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
					return
				}
			}

		}
	}
	response.Success(ctx)
	return
}
