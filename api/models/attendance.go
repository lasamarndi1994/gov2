package models

import "time"

type Attendance struct {
	UserId         uint      `json:"user_id"`
	AttendanceDate time.Time `json:"attendance_date" gorm:"not null;type:date"`
	CheckIn        time.Time `json:"check_in" gorm:"not null;type:time"`
	CheckOut       time.Time `json:"check_out" grom:"not null;type:time"`
	Status         string    `json:"status" gorm:"type:enum('Leave','Absent','Present')"`
}
