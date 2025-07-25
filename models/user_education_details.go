package models

import (
	"time"

	"gorm.io/gorm"
)

type EductionDetails struct {
	Id              uint `json:"id" gorm:"unique;autoIncrement;primaryKey"`
	UserId          uint
	InstitutionName string
	DegreeName      string
	PassoutYear     string
	TotalMark       uint16
	Percentage      uint16
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}
