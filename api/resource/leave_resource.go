package resource

import (
	"github.com/lasamarndi1994/gov2/api/models"
	"github.com/lasamarndi1994/gov2/utility/response"
)

type LeaveTypeResponse struct {
	Name string `json:"name"`
}

type LeaveResponse struct {
	ID        uint   `json:"id"`
	LeaveType string `json:"leave_type"`
	FromDate  string `json:"from_date"`
	ToDate    string `json:"to_date"`
	Remarks   string `json:"remarks"`
}

// Map from []models.Leave → []LeaveResponse
func FromLeaves(leaves []models.Leave) interface{} {
	res := make([]LeaveResponse, 0, len(leaves))
	for _, l := range leaves {
		res = append(res, LeaveResponse{
			ID:        l.Id,
			LeaveType: l.LeaveType.Name,
			Remarks:   l.Remarks,
		})
	}
	return response.SuccessResponse(res)

	//return respose.
}
