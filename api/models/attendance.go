package models

import "time"

type Attendance struct {
	Id             uint      `json:"id" gorm:"primarykey;autoIncrement"`
	UserId         uint      `json:"user_id"`
	AttendanceDate time.Time `json:"attendance_date" gorm:"not null;type:date"`
	CheckIn        time.Time `json:"check_in" gorm:"not null;type:time"`
	CheckOut       time.Time `json:"check_out" grom:"not null;type:time"`
	Status         string    `json:"status" gorm:"type:enum('Leave','Absent','Present')"`
	Remarks        string    `json:"remarks"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	User           User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
