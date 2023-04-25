package validator

import (
	"fmt"
	"strings"
	"testgoapi/resource"

	govalidator "github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Code    int
	Message string
}

var messageFormat = map[string]string{
	"required": "%s is required",
}

func Fails(params interface{}) any {
	err := govalidator.New().Struct(params)

	if errors, ok := err.(govalidator.ValidationErrors); ok {
		e := errors[0]

		field := e.Field()
		field = strings.ToLower(field)
		tag := e.Tag()

		message := fmt.Sprintf(messageFormat[tag], field)

		return ErrorResponse{
			Code:    resource.ErrorCode[field],
			Message: message,
		}
	}

	return nil
}
