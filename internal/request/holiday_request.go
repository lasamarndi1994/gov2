package request

type HolidayRequest struct {
	Name   string `json:"name"`
	Date   string `json:"date"`
	Status bool   `json:"status"`
}
