package handler

import (
	"hello-go/pkg/common/context"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

// ToReqContext converts echo context to request context
func ToReqContext(e echo.Context) *context.ReqContext {
	requestID := e.Get("RequestID").(string)
	logger := e.Get("Logger").(*zerolog.Logger)

	return context.New(requestID, logger)
}
