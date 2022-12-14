// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameReport = "report"

// Report mapped from table <report>
type Report struct {
	ID              int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	TeamID          int64          `gorm:"column:team_id;not null" json:"team_id"`                         // 团队ID
	Rank            int64          `gorm:"column:rank;not null" json:"rank"`                               // 团队内份数
	PlanID          int64          `gorm:"column:plan_id;not null" json:"plan_id"`                         // 计划ID
	PlanName        string         `gorm:"column:plan_name;not null" json:"plan_name"`                     // 计划名称
	SceneID         int64          `gorm:"column:scene_id;not null" json:"scene_id"`                       // 场景ID
	SceneName       string         `gorm:"column:scene_name;not null" json:"scene_name"`                   // 场景名称
	TaskType        int32          `gorm:"column:task_type;not null" json:"task_type"`                     // 任务类型
	TaskMode        int32          `gorm:"column:task_mode;not null" json:"task_mode"`                     // 压测模式
	Status          int32          `gorm:"column:status;not null" json:"status"`                           // 报告状态1:进行中，2:已完成
	RanAt           time.Time      `gorm:"column:ran_at;not null;default:CURRENT_TIMESTAMP" json:"ran_at"` // 启动时间
	RunUserIdentify string         `gorm:"column:run_user_identify;not null" json:"run_user_identify"`
	RunUserID       int64          `gorm:"column:run_user_id;not null" json:"run_user_id"` // 启动人id
	CreatedAt       time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName Report's table name
func (*Report) TableName() string {
	return TableNameReport
}
