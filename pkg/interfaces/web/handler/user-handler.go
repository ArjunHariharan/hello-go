package handler

import (
	"fmt"
	"hello-go/pkg/application"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CreateUser is used to create a user
func CreateUser(c echo.Context) error {
	d := application.CreateUserDto{}

	if err := RestDtoTransformer(c, &d); err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(d)

	return c.JSON(http.StatusOK, d)
}
