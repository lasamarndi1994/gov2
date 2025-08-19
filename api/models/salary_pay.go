package models

import "time"

type SalaryPay struct {
	Id                  uint       `json:"id" gorm:"unique;primaryKey,autoIncrement"`
	UserID              uint       `json:"user_id"`
	BasicSalary         int64      `json:"basic_slary"`
	HouseRentAllowance  int64      `json:"house_rent_allowance"`
	ConveyanceAllowance int64      `json:"conveyance_allowance"`
	MedicalAllowance    int64      `json:"medical_allowance"`
	OtherAllowances     int64      `json:"other_allowance"`
	CreatedAt           *time.Time `json:"created_at"`
	UpdatedAt           *time.Time `json:"updated_at"`
}
