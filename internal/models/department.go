package models

import (
	"time"

	"gorm.io/gorm"
)

type Department struct {
	Id        uint   `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Nmae      string `json:"name" gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
