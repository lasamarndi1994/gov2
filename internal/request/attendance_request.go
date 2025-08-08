package request

import "time"

type AttendanceRequest struct {
	UserId         uint      `json:"user_id" binding:"required"`
	AttendanceDate time.Time `json:"attendance_date" binding:"required"`
	CheckIn        time.Time `json:"check_in" binding:"required"`
	CheckOut       time.Time `json:"check_out" binding:"required"`
	Status         string    `json:"status"`
}
