package models

import (
	"time"

	"gorm.io/gorm"
)

type EductionDetails struct {
	Id              uint   `json:"id" gorm:"unique;autoIncrement;primaryKey"`
	UserId          uint   `json:"department_id" gorm:"not null"`
	InstitutionName string `json:"institution_name" gorm:"size:256"`
	DegreeName      string `json:"degree_name"  gorm:"size:256"`
	PassoutYear     string `json:"passout_year" gorm:"size:256"`
	TotalMark       uint16 `json:"total_mark" `
	Percentage      uint16 `json:"percentage" `
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
	User            User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
