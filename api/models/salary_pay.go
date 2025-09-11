package models

import "time"

type SalaryPay struct {
	Id                  uint       `json:"id" gorm:"unique;primaryKey,autoIncrement"`
	UserId              uint       `json:"user_id"`
	BasicSalary         float64    `json:"basic_slary"`
	HouseRentAllowance  float64    `json:"house_rent_allowance"`
	ConveyanceAllowance float64    `json:"conveyance_allowance"`
	MedicalAllowance    float64    `json:"medical_allowance"`
	DearnessAllowance   float64    `json:"dearness_allowance"`
	Allowance           float64    `json:"allowance"`
	OtherAllowances     float64    `json:"other_allowance"`
	CreatedAt           *time.Time `json:"created_at"`
	UpdatedAt           *time.Time `json:"updated_at"`
}
