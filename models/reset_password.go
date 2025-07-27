package models

import "time"

type PasswordReset struct {
	Id        uint   `gorm:"primaryKey;autoIncrement"`
	UserID    uint   `gorm:"index"`
	Token     string `gorm:"idx_password_resets_token"`
	ExpiresAt time.Time
	CreatedAt time.Time
}
