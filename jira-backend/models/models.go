package models

import (
	"database/sql"
	"time"
)

type User struct {
	UserId uint `gorm:"primaryKey;auto_increment;not_null"`
	//RoleId    uint
	Username  string
	Firstname string
	Lastname  string
	IsDeleted sql.NullBool `gorm:default:false"`
	EmailId   string
	CreatedAt time.Time `gorm:"autoUpdateTime:milli"`
}
type Project struct {
	ProjectId   uint `gorm:"primaryKey;auto_increment;not_null"`
	ProjectName string
	IsDeleted   bool
	CreatedAt   time.Time `gorm:"autoUpdateTime:milli"`
	UpdatedAt   time.Time
}

type Sprint struct {
	SprintId   uint `gorm:"primaryKey;auto_increment;"`
	SprintName string
	ProjectRef uint
	CreatedAt  time.Time `gorm:"autoUpdateTime:milli"`
	UpdatedAt  time.Time
	StartDate  time.Time
	EndDate    time.Time
	Status     uint         `gorm:"default:1"`
	Project    Project      `gorm:"foreignKey:ProjectRef"`
	IsDeleted  sql.NullBool `gorm:default:false"`
}
type Issue struct {
	IssueId     uint `gorm:"primaryKey;auto_increment;not_null"`
	Status      uint `gorm:"default:1"`
	Type        uint `gorm:"default:1"` //`epic/task/subtask/bug/`
	Title       string
	CreatedBy   uint //user ref
	SprintRef   uint
	AssigneeId  uint //user ref
	ProjectRef  uint
	IsDeleted   sql.NullBool
	UpdatedAt   time.Time
	Description string
	CreatedAt   time.Time `gorm:"autoUpdateTime:milli"`
	Sprint      Sprint    `gorm:"foreignKey:SprintRef"`
	AssignedTo  User      `gorm:"foreignKey:AssigneeId"`
	Creator     User      `gorm:"foreignKey:CreatedBy"`
	Project     Project   `gorm:"foreignKey:ProjectRef"`
}

type UserRole struct {
	UserId     uint `gorm:"primaryKey"`
	RoleId     uint `gorm:"primaryKey"`
	ProjectId  uint `gorm:"primaryKey"`
	Membership uint
	Project    Project `gorm:"foreignKey:ProjectId"`
	User       User    `gorm:"foreignKey:UserId"`
	Role       Role    `gorm:"foreignKey:RoleId"`
}

type Role struct {
	RoleId        uint `gorm:"primaryKey"`
	RoleName      string
	PermissionRef uint
	Permission    Permission `gorm:"foreignKey:PermissionRef"`
}

type Permission struct {
	PermissionId   uint `gorm:"primaryKey"`
	PermissionName string
}
