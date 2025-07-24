package models

import (
	"time"

	"gorm.io/gorm"
)

type Designation struct {
	Id        uint   `json:"id" gorm:"unique;autoIncrement;primaryKey"`
	Name      string `json:"name" gorm:"unique"`
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
