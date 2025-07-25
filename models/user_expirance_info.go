package models

import (
	"time"

	"gorm.io/gorm"
)

type ExpiranceInfo struct {
	Id             uint `json:"id" gorm:"unique;autoIncrement;primaryKey"`
	UserId         uint
	CompanyName    string
	Designation    string
	JoiningDate    string
	ExistDate      string
	AboutExpirance *string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}
