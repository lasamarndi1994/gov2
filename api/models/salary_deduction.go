package models

import "time"

type SalaryDeduction struct {
	Id              uint       `json:"id" gorm:"unique;primaryKey,autoIncrement"`
	UserId          uint       `json:"user_id"`
	ProfessionalTax float64    `json:"professional_tax"`
	IncomeTax       float64    `json:"income_tax"`
	Leave           float64    `json:"leave"`
	ProvidentFund   float64    `json:"ProvidentFund"`
	OtherDeduction  float64    `json:"other_deduction"`
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
}
