// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID          int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Email       string         `gorm:"column:email;not null" json:"email"`
	Password    string         `gorm:"column:password;not null" json:"password"`
	Nickname    string         `gorm:"column:nickname;not null" json:"nickname"`
	Avatar      string         `gorm:"column:avatar" json:"avatar"`
	LastLoginAt time.Time      `gorm:"column:last_login_at" json:"last_login_at"`
	CreatedAt   time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
