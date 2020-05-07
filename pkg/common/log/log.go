package log

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// GetLogger returns a contextual logger
func GetLogger(rID string) *zerolog.Logger {
	sublogger := log.With().
		Str("RequestID", rID).
		Caller().
		Logger()

	return &sublogger
}
