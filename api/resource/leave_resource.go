package resource

import (
	"time"

	"github.com/lasamarndi1994/gov2/api/models"
	"github.com/lasamarndi1994/gov2/utility/response"
)

type LeaveTypeResponse struct {
	Name string `json:"name"`
}

type LeaveResponse struct {
	ID             uint      `json:"id"`
	LeaveType      string    `json:"leave_type"`
	FromDate       time.Time `json:"from_date"`
	ToDate         time.Time `json:"to_date"`
	FromLeaveValue string    `json:"from_leave_value"`
	ToLeaveValue   string    `json:"to_leave_value"`
	Remarks        string    `json:"remarks"`
}

// Map from []models.Leave → []LeaveResponse
func FromLeaves(leaves []models.Leave) interface{} {
	res := make([]LeaveResponse, 0, len(leaves))
	for _, l := range leaves {
		res = append(res, LeaveResponse{
			ID:             l.Id,
			LeaveType:      l.LeaveType.Name,
			FromDate:       l.FromDate,
			ToDate:         l.ToDate,
			FromLeaveValue: l.FromLeaveValue,
			ToLeaveValue:   l.ToLeaveValue,
			Remarks:        l.Remarks,
		})
	}
	return response.SuccessResponse(res)
}
