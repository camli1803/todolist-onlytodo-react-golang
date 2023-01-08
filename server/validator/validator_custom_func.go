package validator

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

func IsTodoContent(fl validator.FieldLevel) bool {
	content := fl.Field().String()
	var hasLetter = false
	for _, c := range content {
		if unicode.IsLetter(c) {
			hasLetter = true
		}
	}

	return hasLetter
}
