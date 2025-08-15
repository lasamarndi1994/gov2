package request

type Leave struct {
	LeaveTypeId       uint        `json:"leave_type" binding:"required"`
	FromDate          string      `json:"from_date" binding:"required"`
	ToDate            string      `json:"to_date" binding:"required"`
	FromLeaveValue    string      `json:"from_leave_value" binding:"required"`
	ToLeaveValue      string      `json:"to_leave_value" binding:"required"`
	Remarks           string      `json:"remarks" binding:"required"`
	Attachement       string      `json:"attachement"`
	EmailNotification interface{} `json:"email_notification"`
}
