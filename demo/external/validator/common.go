package common

import (
	"regexp"
	"strconv"

	"github.com/go-playground/validator/v10"
)

func NewValidator() *validator.Validate {
	val := validator.New()
	registerValidation(val)
	return val
}

// 参数校验
func registerValidation(validate *validator.Validate) {
	_ = validate.RegisterValidation("studentType", studentTypeValidation)
	_ = validate.RegisterValidation("twoDecimal", twoDecimal)
}

const STUDENT_TYPE = 1 // "pupil"
func studentTypeValidation(fl validator.FieldLevel) bool {
	if val, ok := fl.Field().Interface().(int32); ok {
		return val == STUDENT_TYPE
	}
	return false
}

func twoDecimal(fl validator.FieldLevel) bool {
	if val, ok := fl.Field().Interface().(float64); ok {
		str := strconv.FormatFloat(val, 'f', -1, 64)
		matched, err := regexp.Match("^-?[0-9]+(.[0-9]{1,2})?$", []byte(str))
		if err != nil {
			return false
		}
		return matched
	}
	return false
}
