// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUserTeam = "user_team"

// UserTeam mapped from table <user_team>
type UserTeam struct {
	ID     int64 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID int64 `gorm:"column:user_id;not null" json:"user_id"`
	TeamID int64 `gorm:"column:team_id;not null" json:"team_id"`
	Sort   int32 `gorm:"column:sort;not null" json:"sort"`
}

// TableName UserTeam's table name
func (*UserTeam) TableName() string {
	return TableNameUserTeam
}
