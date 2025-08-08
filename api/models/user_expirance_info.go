package models

import (
	"time"

	"gorm.io/gorm"
)

type ExpiranceInfo struct {
	Id             uint   `json:"id" gorm:"unique;autoIncrement;primaryKey"`
	UserId         uint   `json:"user_id"`
	CompanyName    string `json:"company_name" gorm:"size:256"`
	Designation    string `json:"designation" gorm:"size:256"`
	JoiningDate    string `json:"joining_date" gorm:"size:256;type:date"`
	ExistDate      string `json:"exist_date" gorm:"size:256;type:date"`
	AboutExpirance string `json:"about_expirance" gorm:"size:256;type:date"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
	User           User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
