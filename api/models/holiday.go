package models

import "time"

type Holiday struct {
	Id        uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string     `json:"name" gorm:"unique"`
	Date      time.Time  `json:"date" gorm:"type:date"`
	Status    bool       `json:"status" gorm:"default: 1"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"-"`
	CreatedBy uint       `json:"-"`
}
