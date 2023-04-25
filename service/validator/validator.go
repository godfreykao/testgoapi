package validator

import (
	"strings"
	"testgoapi/resource"

	govalidator "github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Code    int
	Message string
}

func Fails(params interface{}) any {
	err := govalidator.New().Struct(params)

	if errors, ok := err.(govalidator.ValidationErrors); ok {
		e := errors[0]

		field := e.Field()
		field = strings.ToLower(field)

		return ErrorResponse{
			Code:    resource.ErrorCode[field],
			Message: "Invalid " + field,
		}
	}

	return nil
}
