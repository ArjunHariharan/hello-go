package handler

import (
	"fmt"
	"hello-go/pkg/application"
	"net/http"

	"github.com/labstack/echo/v4"
)

//UserHandler implementation
type UserHandler struct {
	a *application.UserApplication
}

// NewUserHandler creates an instance of user handler
func NewUserHandler(a *application.UserApplication) *UserHandler {
	return &UserHandler{a: a}
}

// CreateUser is used to create a user
func (h *UserHandler) CreateUser(c echo.Context) error {
	d := application.CreateUserDto{}

	if err := RestDtoTransformer(c, &d); err != nil {
		fmt.Println(err.Error())
		return err
	}

	rc := ToReqContext(c)

	h.a.SaveUser(rc)
	fmt.Println(d)

	return c.JSON(http.StatusOK, d)
}
