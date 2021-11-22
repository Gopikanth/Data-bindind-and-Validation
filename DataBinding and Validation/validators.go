package validators

import (
	"strings"

	"gopkg.in/go-playground/validator.v9" //this package is imported for the validation
)

// ValidateCoolTitle returns true when the field value contains the word "celebrity".
func ValidateCoolTitle(fl validator.FieldLevel) bool {
	return strings.Contains(fl.Field().String(), "celebrity")
}

//we can return error when there is no proper validation
