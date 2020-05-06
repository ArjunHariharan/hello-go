package handler

import (
	"hello-go/pkg/common/context"

	"github.com/labstack/echo/v4"
)

// ToReqContext converts echo context to request context
func ToReqContext(e echo.Context) *context.ReqContext {
	return context.New("1234")
}
