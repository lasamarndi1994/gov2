package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id           uint   `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	UserId       string `json:"emp_id" gotm:"unique"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email" gorm:"unique"`
	MobileNumber int    `json:"mobile_number" gorm:"unique"`
	Password     string `json:"password"`
	Status       string `json:"status"`
	JoiningDate  string `json:"joining_date"`
	DateofBirth  time.Time
	AboutMe      *string `json:"about_me"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}
