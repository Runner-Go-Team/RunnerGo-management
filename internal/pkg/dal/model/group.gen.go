// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameGroup = "group"

// Group mapped from table <group>
type Group struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	TargetID  int64          `gorm:"column:target_id" json:"target_id"` // 目标id
	Request   string         `gorm:"column:request;not null" json:"request"`
	Script    string         `gorm:"column:script;not null" json:"script"`
	CreatedAt time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName Group's table name
func (*Group) TableName() string {
	return TableNameGroup
}