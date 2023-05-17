package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

func validatorErrorParserInParams(ctx *gin.Context, err error) {
	var verr validator.ValidationErrors

	erroAs := errors.As(err, &verr)

	if erroAs {
		out := make([]ErrorMsg, len(verr))
		for i, fe := range verr {
			out[i] = ErrorMsg{
				Field:   fe.Field(),
				Message: getErrorMsg(fe),
			}

		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
	}
}
