package internalerrors

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"strings"
)

func ValidadeStruct(obj interface{}) error {
	validate := validator.New()
	err := validate.Struct(obj)
	if err == nil {
		return nil
	}
	var validationErrors validator.ValidationErrors
	errors.As(err, &validationErrors)
	validationError := validationErrors[0]

	field := strings.ToLower(validationError.StructField())
	switch validationError.Tag() {
	case "required":
		return errors.New(field + " is required")
	case "max":
		return errors.New(field + " is required with max " + validationError.Param())
	case "min":
		return errors.New(field + " is required with min " + validationError.Param())
	case "email":
		return errors.New(field + " is invalid")
	}
	return nil
}
