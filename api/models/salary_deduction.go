package models

import "time"

type SalaryDeduction struct {
	Id              uint       `json:"id" gorm:"unique;primaryKey,autoIncrement"`
	UserID          uint       `json:"user_id"`
	OtherDeduction  int64      `json:"other_deduction"`
	ProfessionalTax int64      `json:"professional_tax"`
	IncomeTax       int64      `json:"IncomeTax"`
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
}
