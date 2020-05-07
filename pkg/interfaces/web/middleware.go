package web

import (
	"hello-go/pkg/common/log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	uuid "github.com/satori/go.uuid"
)

const (
	requestIDHeader = "X-Request-ID"
	requestID       = "RequestID"
	logger          = "Logger"
)

func registerMiddleware(e *echo.Echo) {
	registerRequestIDMiddleware(e)
	registerRequestLoggerMiddleware(e)
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
			c.Set(requestID, rID)

			return next(c)
		}
	})
}

func registerRequestLoggerMiddleware(e *echo.Echo) {
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// This is used to create request context
			rID := c.Get(requestID).(string)
			l := log.GetLogger(rID)

			c.Set(logger, l)

			req := c.Request()
			res := c.Response()
			start := time.Now()

			// Log request
			l.Info().Dict("Request", zerolog.Dict().
				Str("Method", req.Method).
				Str("Path", req.URL.Path),
			).Msg("Incoming Request")

			if err := next(c); err != nil {
				c.Error(err)
				return nil
			}

			// Log response
			l.Info().Dict("Response", zerolog.Dict().
				Int("Status", res.Status).
				Int64("Duration", time.Since(start).Milliseconds()),
			).Msg("Sending Response")

			return nil
		}
	})
}
