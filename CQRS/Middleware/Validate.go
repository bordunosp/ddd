package Middleware

import (
	"github.com/go-playground/validator/v10"
)

// @url https://github.com/go-playground/validator
var _validator = validator.New()

func ValidateRegisterCustomTypeFunc(fn validator.CustomTypeFunc, types ...any) {
	_validator.RegisterCustomTypeFunc(fn, types...)

}

func Validate(val any) error {
	return _validator.Struct(val)
}
