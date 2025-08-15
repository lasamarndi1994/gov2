package models

import (
	"time"
)

type Leave struct {
	Id                uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	LeaveTypeId       uint       `json:"leave_type_id"`
	UserId            uint       `json:"user_id"`
	FromDate          time.Time  `json:"from_date" gorm:"type:date"`
	ToDate            time.Time  `json:"to_date" gorm:"type:date"`
	FromLeaveValue    string     `json:"from_leave_value"`
	ToLeaveValue      string     `json:"to_leave_value"`
	Remarks           string     `json:"remarks"`
	Attachement       string     `json:"attachement"`
	EmailNotification string     `gorm:"type:json;nullable"`
	LeaveStatus       string     `json:"leave_status" gorm:"type:enum('Pending','Aprroved','Cancel');default:'Pending'"`
	CreatedAt         *time.Time `json:"created_at"`
	UpdatedAt         *time.Time `json:"updated_at"`
}
