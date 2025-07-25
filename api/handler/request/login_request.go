package request

type LoginReuest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password"`
}
