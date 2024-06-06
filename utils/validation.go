package utils

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// custom validation for birth-date field
func IsBirthDate(fl validator.FieldLevel) bool {
	_, err := time.Parse(time.DateOnly, fl.Field().String())
	if err != nil {
		return false
	}
	return true
}
