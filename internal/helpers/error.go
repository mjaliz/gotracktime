package helpers

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	var errMsg string
	switch fe.Tag() {
	case "required":
		errMsg = "This field is required"
	case "lte":
		errMsg = fmt.Sprintf("Should be less than %s", fe.Param())
	case "gte":
		errMsg = fmt.Sprintf("Should be greater than %s", fe.Param())
	default:
		errMsg = "Unknown error"
	}
	return errMsg
}

func ParseValidationError(err error) []ErrorMsg {
	var ve validator.ValidationErrors
	out := make([]ErrorMsg, len(ve))
	if errors.As(err, &ve) {
		for _, fe := range ve {
			out = append(out, ErrorMsg{fe.Field(), getErrorMsg(fe)})
		}
	}
	return out
}
