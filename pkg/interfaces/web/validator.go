package web

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// CustomValidator struct definition
type CustomValidator struct {
	validator *validator.Validate
}

// Validate method for Customvalidator
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func registerValidator(e *echo.Echo) {
	validator := validator.New()
	validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	e.Validator = &CustomValidator{validator: validator}
}
