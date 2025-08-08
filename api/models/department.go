package models

import (
	"time"

	"gorm.io/gorm"
)

type Department struct {
	Id        uint           `json:"id" gorm:"unique;autoIncrement;primaryKey"`
	Name      string         `json:"name" gorm:"unique;not null;size:256"`
	Status    string         `json:"status" gorm:"type:enum('Active','Deactive');default:Active;not null"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"upd-ted_at"`
}
