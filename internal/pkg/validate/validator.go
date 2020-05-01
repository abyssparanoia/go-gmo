package validate

import (
	"github.com/go-playground/validator"
)

var v = validator.New()

// Struct ... validate
func Struct(s interface{}) error {
	return v.Struct(s)
}
