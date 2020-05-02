package infra

import (
	"fmt"
	"hello-go/pkg/app"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// TransformDto from json to dto
func TransformDto(c echo.Context, d interface{}) error {
	if err := c.Bind(d); err != nil {
		return app.NewError(app.InternalServerError, "Failed to bind Dto")
	}

	if err := c.Validate(d); err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			s := fmt.Sprintf("key %s failed requirement %s", err.Field(), err.Tag())
			errors = append(errors, s)
		}

		errStr := strings.Join(errors, "\n")

		return app.NewError(app.ValidationError, errStr)
	}

	return nil
}
