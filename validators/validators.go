package validators

import (
	"gopkg.in/go-playground/validator.v9"
	"strings"
)

// this is the way we right custom validation in golang
// Here we are using validator.v9
func ValidateCoolTitel(field validator.FieldLevel) bool {
	return strings.Contains(field.Field().String(), "cool")
}
