package validator

import (
	govalidator "github.com/go-playground/validator/v10"
	"github.com/lgbya/go-dump"
)

type ErrorResponse struct {
	Code    int
	Message string
}

func Fails(params interface{}) any {
	err := govalidator.New().Struct(params)

	if validationErrors, ok := err.(govalidator.ValidationErrors); ok {
		if len(validationErrors) > 0 {
			dump.Printf(validationErrors[0])
		}
	}

	return ErrorResponse{
		Code:    100,
		Message: "Invalid brandid",
	}
}
