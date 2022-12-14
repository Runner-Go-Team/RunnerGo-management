package rao

import (
	"kp-management/internal/pkg/dal/model"
)

type Plan struct {
	PlanID   int64  `json:"plan_id"`
	Rank     int64  `json:"rank"`
	TeamID   int64  `json:"team_id"`
	Name     string `json:"name"`
	TaskType int32  `json:"task_type"`
	Mode     int32  `json:"mode"`
	Status   int32  `json:"status"`
	//RunUserID         int64     `json:"run_user_id"`
	//RunUserName       string    `json:"run_user_name"`
	CreatedUserID     int64     `json:"created_user_id"`
	CreatedUserName   string    `json:"created_user_name"`
	CreatedUserAvatar string    `json:"created_user_avatar"`
	Remark            string    `json:"remark"`
	CreatedTimeSec    int64     `json:"created_time_sec"`
	UpdatedTimeSec    int64     `json:"updated_time_sec"`
	CronExpr          string    `json:"cron_expr"`
	ModeConf          *ModeConf `json:"mode_conf"`
}

type PlanTask struct {
	PlanID        int64          `json:"plan_id"`
	SceneID       int64          `json:"scene_id"`
	TaskType      int32          `json:"task_type"`
	Mode          int32          `json:"mode"`
	ModeConf      *ModeConf      `json:"mode_conf"`
	TimedTaskConf *TimedTaskConf `json:"timed_task_conf"`
}

type ModeConf struct {
	ReheatTime       int64 `json:"reheat_time"`       // 预热时长
	RoundNum         int64 `json:"round_num"`         // 轮次
	Concurrency      int64 `json:"concurrency"`       // 并发数
	ThresholdValue   int64 `json:"threshold_value"`   // 阈值
	StartConcurrency int64 `json:"start_concurrency"` // 起始并发数
	Step             int64 `json:"step"`              // 步长
	StepRunTime      int64 `json:"step_run_time"`     // 步长执行时长
	MaxConcurrency   int64 `json:"max_concurrency"`   // 最大并发数
	Duration         int64 `json:"duration"`          // 稳定持续时长，持续时长
	CreatedTimeSec   int64 `json:"created_time_sec"`  // 创建时间
}

type RunPlanReq struct {
	PlanID  int64   `json:"plan_id"`
	TeamID  int64   `json:"team_id"`
	SceneID []int64 `json:"scene_id"`
}

type RunPlanResp struct {
}

type StopPlanReq struct {
	//ReportIds []int64 `json:"report_ids"`
	TeamID  int64   `json:"team_id"`
	PlanIDs []int64 `json:"plan_ids"`
}

type StopPlanResp struct {
}

type ListUnderwayPlanReq struct {
	TeamID int64 `form:"team_id" binding:"required,gt=0"`
	Page   int   `form:"page,default=1"`
	Size   int   `form:"size,default=10"`
}

type ListUnderwayPlanResp struct {
	Plans []*Plan `json:"plans"`
	Total int64   `json:"total"`
}

type ClonePlanReq struct {
	TeamID int64 `json:"team_id"`
	PlanID int64 `json:"plan_id"`
}

type ClonePlanResp struct {
}

type ListPlansReq struct {
	TeamID int64 `form:"team_id" binding:"required,gt=0"`
	Page   int   `form:"page,default=1"`
	Size   int   `form:"size,default=10"`

	Keyword      string `form:"keyword"`
	StartTimeSec int64  `form:"start_time_sec"`
	EndTimeSec   int64  `form:"end_time_sec"`

	TaskType int32 `form:"task_type"`
	TaskMode int32 `form:"task_mode"`
	Status   int32 `form:"status"`
	Sort     int32 `form:"sort"`
}

type ListPlansResp struct {
	Plans []*Plan `json:"plans"`
	Total int64   `json:"total"`
}

type SavePlanReq struct {
	PlanID int64  `json:"plan_id"`
	TeamID int64  `json:"team_id" binding:"required,gt=0"`
	Name   string `json:"name" binding:"required"`
	Remark string `json:"remark"`
}

type SavePlanConfReq struct {
	PlanID        int64          `json:"plan_id" binding:"required,gt=0"`
	SceneID       int64          `json:"scene_id" binding:"required,gt=0"`
	TeamID        int64          `json:"team_id" binding:"required,gt=0"`
	Name          string         `json:"name" binding:"required"`
	TaskType      int32          `json:"task_type" binding:"required,gt=0"`
	Mode          int32          `json:"mode" binding:"required,gt=0"`
	Remark        string         `json:"remark"`
	ModeConf      *ModeConf      `json:"mode_conf"`
	CronExpr      string         `json:"cron_expr"`
	TimedTaskConf *TimedTaskConf `json:"timed_task_conf"`

	//Nodes           []*Node `json:"nodes"`
	//Edges           []*Edge `json:"edges"`
	//MultiLevelNodes string  `json:"multi_level_nodes"`
}

type TimedTaskConf struct {
	Frequency     int32 `json:"frequency"`       // 频次: 0-一次，1-每天，2-每周，3-每月
	TaskExecTime  int64 `json:"task_exec_time"`  // 任执行时间
	TaskCloseTime int64 `json:"task_close_time"` // 结束时间
}

type SavePlanResp struct {
	PlanID int64 `json:"plan_id"`
}

type GetPlanTaskReq struct {
	PlanID  int64 `form:"plan_id" binding:"required,gt=0"`
	SceneID int64 `form:"scene_id" binding:"required,gt=0"`
	TeamID  int64 `form:"team_id" binding:"required,gt=0"`
}

type GetPlanTaskResp struct {
	PlanTask *PlanTask `json:"plan_task"`
}

type GetPlanConfReq struct {
	TeamID int64 `form:"team_id" binding:"required,gt=0"`
	PlanID int64 `form:"plan_id" binding:"required,gt=0"`
}

type GetPlanResp struct {
	Plan *Plan `json:"plan"`
}

type DeletePlanReq struct {
	PlanID int64 `json:"plan_id"`
	TeamID int64 `json:"team_id"`
}

type DeletePlanResp struct {
}

type PlanEmailReq struct {
	PlanID int64    `json:"plan_id"`
	Emails []string `json:"emails"`
}

type PlanEmailResp struct {
}

type PlanListEmailReq struct {
	TeamID int64 `form:"team_id"`
	PlanID int64 `form:"plan_id"`
}

type PlanListEmailResp struct {
	Emails []*PlanEmail `json:"emails"`
}

type PlanEmail struct {
	ID            int64  `json:"id"`
	PlanID        int64  `json:"plan_id"`
	Email         string `json:"email"`
	CreateTimeSec int64  `json:"create_time_sec"`
}

type PlanDeleteEmailReq struct {
	PlanID  int64 `json:"plan_id"`
	TeamID  int64 `json:"team_id"`
	EmailID int64 `json:"email_id"`
}

type PlanDeleteEmailResp struct {
}

type ImportSceneReq struct {
	PlanID       int64   `json:"plan_id"`
	TargetIDList []int64 `json:"target_id_list"`
}

type ImportSceneResp struct {
	Scenes []*model.Target `json:"scenes"`
}
