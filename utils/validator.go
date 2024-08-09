package utils

import "github.com/go-playground/validator"

// Validate variable to hold *validator.Validate
var Validate *validator.Validate

func init() {
	Validate = validator.New()
}
