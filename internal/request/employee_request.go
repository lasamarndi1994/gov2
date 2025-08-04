package request

type EmployeeRequest struct {
	FirstName     string `json:"first_name" binding:"required"`
	LastName      string `json:"last_name" binding:"required"`
	Email         string `json:"email" binding:"required,email"`
	Password      string `json:"password"`
	MobileNumber  int    `json:"mobile_number" binding:"required,gte=9"`
	EmployeeId    string `json:"employee_id"`
	DepartmentId  uint   `json:"department_id" binding:"required"`
	DesignationId uint   `json:"designation_id"  binding:"required"`
	AboutMe       string `json:"about_me"`
	JoiningDate   string `json:"joining_date"  time_format:"2006-01-02"`
	DateofBirth   string `json:"dateof_birth" binding:"required" time_format:"2006-01-02 15:04:05"`
}

type UserRequest struct {
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password"`
	MobileNumber int    `json:"mobile_number" binding:"required,gte=9"`
}
