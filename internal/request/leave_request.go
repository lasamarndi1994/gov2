package request

import (
	"encoding/json"
	"time"
)

type Leave struct {
	LeaveTypeId       uint         `json:"leave_type"`
	FromDate          time.Time    `json:"from_date" binding:"required"`
	ToDate            time.Time    `json:"to_date" binding:"required"`
	FromLeaveValue    string       `json:"from_leave_value" binding:"required"`
	ToLeaveValue      string       `json:"to_leave_value" binding:"required"`
	Remarks           string       `json:"remarks" binding:"required"`
	Attachement       string       `json:"attachement"`
	EmailNotification json.Encoder `json:"email_notification"`
}
