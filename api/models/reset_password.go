package models

import "time"

type PasswordReset struct {
	Id        uint   `gorm:"primaryKey;autoIncrement"`
	UserId    uint   `json:"user_id" gorm:"index"`
	Token     string `gorm:"idx_password_resets_token"`
	ExpiresAt time.Time
	CreatedAt time.Time
}
