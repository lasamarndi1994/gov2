package request

type SalaryRequest struct {
	BasicSalary         float64 `json:"basic_slary" binding:"required"`
	HouseRentAllowance  float64 `json:"house_rent_allowance" binding:"required"`
	ConveyanceAllowance float64 `json:"conveyance_allowance" binding:"required"`
	MedicalAllowance    float64 `json:"medical_allowance" binding:"required"`
	DearnessAllowance   float64 `json:"dearness_allowance" binding:"required"`
	Allowance           float64 `json:"allowance" binding:"required"`
	OtherAllowances     float64 `json:"other_allowance" binding:"required"`
	ProfessionalTax     float64 `json:"professional_tax"  binding:"required"`
	IncomeTax           float64 `json:"income_tax"  binding:"required"`
	Leave               float64 `json:"leave"  binding:"required"`
	ProvidentFund       float64 `json:"ProvidentFund" binding:"required"`
	OtherDeduction      float64 `json:"other_deduction"  binding:"required"`
}

type Slarly interface {
}
