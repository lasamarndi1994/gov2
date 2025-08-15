package request

import "mime/multipart"

type LeaveReuest struct {
	LeaveTypeId       uint                  `form:"leave_type" binding:"required"`
	FromDate          string                `form:"from_date" binding:"required"`
	ToDate            string                `form:"to_date" binding:"required"`
	FromLeaveValue    string                `form:"from_leave_value" binding:"required"`
	ToLeaveValue      string                `form:"to_leave_value" binding:"required"`
	Remarks           string                `form:"remarks" binding:"required"`
	Attachement       *multipart.FileHeader `form:"attachement"`
	LeaveStatus       string                `form:"leave_status"`
	EmailNotification string                `form:"email_notification"`
}
