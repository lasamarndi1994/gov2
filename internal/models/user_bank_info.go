package models

import (
	"time"

	"gorm.io/gorm"
)

type BankInfo struct {
	Id            uint `json:"id" gorm:"unique;autoIncrement;primaryKey"`
	UserId        uint
	BankName      string
	AccountNumber int64
	IfscCode      string
	Branch        string
	DeletedBy     uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}
