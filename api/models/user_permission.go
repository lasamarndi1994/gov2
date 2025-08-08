package models

import (
	"time"

	"gorm.io/gorm"
)

type UserPremission struct {
	Id        uint `json:"id" gorm:"unique;autoIncrement;primaryKey"`
	UserId    uint `json:"user_id"`
	ModuleId  uint32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	User      User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
