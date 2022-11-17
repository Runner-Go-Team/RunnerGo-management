// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTimedTaskConf = "timed_task_conf"

// TimedTaskConf mapped from table <timed_task_conf>
type TimedTaskConf struct {
	ID            int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                      // 表id
	PlanID        int64          `gorm:"column:plan_id;not null" json:"plan_id"`                                 // 计划id
	SenceID       int64          `gorm:"column:sence_id;not null" json:"sence_id"`                               // 场景id
	TeamID        int64          `gorm:"column:team_id;not null" json:"team_id"`                                 // 团队id
	UserID        int64          `gorm:"column:user_id;not null" json:"user_id"`                                 // 用户ID
	Frequency     int32          `gorm:"column:frequency;not null" json:"frequency"`                             // 任务执行频次
	TaskExecTime  int64          `gorm:"column:task_exec_time;not null" json:"task_exec_time"`                   // 任务执行时间
	TaskCloseTime int64          `gorm:"column:task_close_time;not null" json:"task_close_time"`                 // 任务结束时间
	TaskType      int32          `gorm:"column:task_type;not null;default:2" json:"task_type"`                   // 任务类型：1-普通任务，2-定时任务
	TaskMode      int32          `gorm:"column:task_mode;not null;default:1" json:"task_mode"`                   // 压测模式：1-并发模式，2-阶梯模式，3-错误率模式，4-响应时间模式，5-每秒请求数模式，  6 //每秒事务数模式，
	ModeConf      string         `gorm:"column:mode_conf;not null" json:"mode_conf"`                             // 压测详细配置
	Status        int32          `gorm:"column:status;not null" json:"status"`                                   // 任务状态：0-未启用，1-运行中，3-已过期
	CreatedAt     time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"` // 创建时间
	UpdatedAt     time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"` // 更新时间
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`                                    // 删除时间
}

// TableName TimedTaskConf's table name
func (*TimedTaskConf) TableName() string {
	return TableNameTimedTaskConf
}
