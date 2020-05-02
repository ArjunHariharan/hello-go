package user

import (
	"fmt"
	"hello-go/pkg/infra"
	"net/http"

	"github.com/labstack/echo/v4"
)

func create(c echo.Context) error {
	d := new(CreateUserDto)
	if err := infra.TransformDto(c, d); err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(d)

	return c.JSON(http.StatusOK, d)
}
