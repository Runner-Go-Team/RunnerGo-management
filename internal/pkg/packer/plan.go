package packer

import (
	"kp-management/internal/pkg/biz/consts"
	"kp-management/internal/pkg/dal/mao"
	"kp-management/internal/pkg/dal/model"
	"kp-management/internal/pkg/dal/rao"
)

func TransPlansToRaoPlanList(plans []*model.Plan, users []*model.User) []*rao.Plan {
	ret := make([]*rao.Plan, 0)
	for _, p := range plans {
		for _, u := range users {
			if p.RunUserID == u.ID {
				ret = append(ret, &rao.Plan{
					PlanID:         p.ID,
					TeamID:         p.TeamID,
					Name:           p.Name,
					TaskType:       p.TaskType,
					Mode:           p.Mode,
					Status:         p.Status,
					RunUserID:      p.RunUserID,
					RunUserName:    u.Nickname,
					Remark:         p.Remark,
					CreatedTimeSec: p.CreatedAt.Unix(),
					UpdatedTimeSec: p.UpdatedAt.Unix(),
				})
			}
		}
	}

	return ret
}

func TransSavePlanReqToPlanModel(req *rao.SavePlanConfReq, userID int64) *model.Plan {
	return &model.Plan{
		ID:           req.PlanID,
		TeamID:       req.TeamID,
		Name:         req.Name,
		TaskType:     req.TaskType,
		Mode:         req.Mode,
		Status:       consts.PlanStatusNormal,
		CreateUserID: userID,
		Remark:       req.Remark,
		CronExpr:     req.CronExpr,
	}
}

func TransSavePlanReqToMaoTask(req *rao.SavePlanConfReq) *mao.Task {
	mc := req.ModeConf

	return &mao.Task{
		PlanID: req.PlanID,
		ModeConf: &mao.ModeConf{
			ReheatTime:       mc.ReheatTime,
			RoundNum:         mc.RoundNum,
			Concurrency:      mc.Concurrency,
			ThresholdValue:   mc.ThresholdValue,
			StartConcurrency: mc.StartConcurrency,
			Step:             mc.Step,
			StepRunTime:      mc.StepRunTime,
			MaxConcurrency:   mc.MaxConcurrency,
			Duration:         mc.Duration,
		},
	}

}

func TransTaskToRaoPlan(p *model.Plan, t *mao.Task) *rao.Plan {

	var mc rao.ModeConf
	if t != nil {
		mc = rao.ModeConf{
			ReheatTime:       t.ModeConf.ReheatTime,
			RoundNum:         t.ModeConf.RoundNum,
			Concurrency:      t.ModeConf.Concurrency,
			ThresholdValue:   t.ModeConf.ThresholdValue,
			StartConcurrency: t.ModeConf.StartConcurrency,
			Step:             t.ModeConf.Step,
			StepRunTime:      t.ModeConf.StepRunTime,
			MaxConcurrency:   t.ModeConf.MaxConcurrency,
			Duration:         t.ModeConf.Duration,
		}
	}

	return &rao.Plan{
		PlanID:         p.ID,
		TeamID:         p.TeamID,
		Name:           p.Name,
		TaskType:       p.TaskType,
		Mode:           p.Mode,
		Status:         p.Status,
		RunUserID:      p.RunUserID,
		RunUserName:    "",
		Remark:         p.Remark,
		CreatedTimeSec: p.CreatedAt.Unix(),
		UpdatedTimeSec: p.UpdatedAt.Unix(),
		CronExpr:       p.CronExpr,
		ModeConf:       &mc,
	}
}

func TransSetPreinstallReqToMaoPreinstall(req *rao.SetPreinstallReq) *mao.Preinstall {
	return &mao.Preinstall{
		TeamID:   req.TeamID,
		TaskType: req.TaskType,
		CronExpr: req.CronExpr,
		Mode:     req.Mode,
		ModeConf: &mao.ModeConf{
			ReheatTime:       req.ModeConf.ReheatTime,
			RoundNum:         req.ModeConf.RoundNum,
			Concurrency:      req.ModeConf.Concurrency,
			ThresholdValue:   req.ModeConf.ThresholdValue,
			StartConcurrency: req.ModeConf.StartConcurrency,
			Step:             req.ModeConf.Step,
			StepRunTime:      req.ModeConf.StepRunTime,
			MaxConcurrency:   req.ModeConf.MaxConcurrency,
			Duration:         req.ModeConf.Duration,
		},
	}
}

func TransMaoPreinstallToRaoPreinstall(p *mao.Preinstall) *rao.Preinstall {
	return &rao.Preinstall{
		TeamID:   p.TeamID,
		TaskType: p.TaskType,
		CronExpr: p.CronExpr,
		Mode:     p.Mode,
		ModeConf: &rao.ModeConf{
			ReheatTime:       p.ModeConf.ReheatTime,
			RoundNum:         p.ModeConf.RoundNum,
			Concurrency:      p.ModeConf.Concurrency,
			ThresholdValue:   p.ModeConf.ThresholdValue,
			StartConcurrency: p.ModeConf.StartConcurrency,
			Step:             p.ModeConf.Step,
			StepRunTime:      p.ModeConf.StepRunTime,
			MaxConcurrency:   p.ModeConf.MaxConcurrency,
			Duration:         p.ModeConf.Duration,
		},
	}
}
