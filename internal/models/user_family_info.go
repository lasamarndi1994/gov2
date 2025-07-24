package models

import (
	"time"

	"gorm.io/gorm"
)

type FamilyInfo struct {
	Id                uint `json:"id" gorm:"unique;autoIncrement;primaryKey"`
	UserId            uint
	RelativeName      string
	RelationshipeName string
	ContactNumber     string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt
}
