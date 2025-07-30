package request

type EmployeeRequest struct {
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password" binding:"required,gte=6"`
	MobileNumber int    `json:"mobile_number" binding:"required,gte=9"`
}
