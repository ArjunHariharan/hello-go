package handler

import (
	"hello-go/pkg/common/context"

	"github.com/labstack/echo/v4"
)

// ToReqContext converts echo context to request context
func ToReqContext(e echo.Context) *context.ReqContext {
	requestID := e.Get(context.RequestID).(string)

	return context.New(requestID)
}
