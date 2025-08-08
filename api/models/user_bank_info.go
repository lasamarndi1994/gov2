package models

import (
	"time"

	"gorm.io/gorm"
)

type BankInfo struct {
	Id            uint      `json:"id" gorm:"unique;autoIncrement;primaryKey"`
	UserId        uint      `json:"user_id"`
	BankName      string    `json:"bank_name"`
	AccountNumber int64     `json:"AccountNumber"`
	IfscCode      string    `json:"ifsc_code"`
	Branch        string    `json:"branch"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     gorm.DeletedAt
	User          User `gorm:"constraint:onUpdate:CASCADE,onDelete:CASCADE"`
}
