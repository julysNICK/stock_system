package api

import (
	"github.com/go-playground/validator"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {

	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "This field must be a valid email address"
	case "gte":
		return "This field must be greater than or equal to 0"
	case "lte":
		return "This field must be less than or equal to 100"
	case "len":
		return "This field must be 6 characters long"

	default:
		return "This field is invalid"
	}
}
