package models

import "time"

type UserDesignation struct {
	Id             uint      `json:"id"`
	UserId         uint      `json:"user_id" gorm:"not null"`
	DesiginationId uint      `json:"designation_id" gorm:"not null"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
	User           User      `gorm:"constraint:onDelete:CASCADE,onUpdate:CASCADE"`
}
