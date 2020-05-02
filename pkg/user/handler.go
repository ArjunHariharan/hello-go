package user

import (
	"fmt"
	"hello-go/pkg/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

func create(c echo.Context) error {
	d := new(CreateUserDto)
	if err := util.RestDtoTransformer(c, d); err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(d)

	return c.JSON(http.StatusOK, d)
}
