package models

import "time"

type LevaeType struct {
	Id        uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"unique;size:256;not null"`
	Status    string    `json:"status" gorm:"type:enum('Active','Deactive');default:Active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
