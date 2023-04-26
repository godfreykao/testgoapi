package validator

import (
	"encoding/json"
	"regexp"
	"strings"
	"testgoapi/resource"

	"github.com/davecgh/go-spew/spew"
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

func BodyParserFails(params interface{}) any {
	spew.Dump(params)
	if value, ok := params.(*json.UnmarshalTypeError); ok {
		rgx := regexp.MustCompile(`json: cannot unmarshal .+ into Go struct field .([a-zA-z]+)`)
		matches := rgx.FindStringSubmatch(value.Error())

		field := matches[1]

		return ErrorResponse{
			Code:    resource.ErrorCode[field],
			Message: "Invalid " + field,
		}
	}

	return nil
}

func FieldFails(field string) ErrorResponse {
	return ErrorResponse{
		Code:    resource.ErrorCode[field],
		Message: "Invalid " + field,
	}
}
