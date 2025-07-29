package request

import "github.com/go-playground/validator/v10"

type ResetPasswordRequest struct {
	Token           string `json:"token" binding:"required"`
	Password        string `json:"password" binding:"required,gte=6"`
	ConfirmPassword string `json:"confirm_password" binding:"required,gte=6"`
}

var Validate *validator.Validate

func InitValidator() {
	Validate = validator.New()
	Validate.RegisterStructValidation(ResetPasswordStructLevelValidation, ResetPasswordRequest{})
}

func ResetPasswordStructLevelValidation(sl validator.StructLevel) {
	req := sl.Current().Interface().(ResetPasswordRequest)
	if req.Password != req.ConfirmPassword {
		sl.ReportError(req.ConfirmPassword, "ConfirmPassword", "confirm_password", "passwordmatch", "")
	}
}
