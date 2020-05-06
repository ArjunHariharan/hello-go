package web

import (
	"hello-go/pkg/common/context"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

const (
	requestIDHeader = "X-Request-ID"
)

func registerMiddleware(e *echo.Echo) {
	registerRequestIDMiddleware(e)
}

func registerRequestIDMiddleware(e *echo.Echo) {
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()
			rID := req.Header.Get(echo.HeaderXRequestID)

			if rID == "" {
				id := uuid.NewV4()
				rID = id.String()
			}

			res.Header().Set(requestIDHeader, rID)

			// This is used to create request context
			c.Set(context.RequestID, rID)

			return next(c)
		}
	})
}
