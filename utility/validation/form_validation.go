package validation

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) map[string]string {
	errors := map[string]string{}

	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range ve {
			field := strings.ToLower(fe.Field()) // field name like 'email'
			switch fe.Tag() {
			case "required":
				errors[field] = field + " is required"
			case "email":
				errors[field] = "Invalid email address"
			case "gte":
				errors[field] = field + " must be at least " + fe.Param()
			case "lte":
				errors[field] = field + " must be at most " + fe.Param()
			default:
				errors[field] = "invalid value for " + field
			}
		}
	}

	return errors
}
