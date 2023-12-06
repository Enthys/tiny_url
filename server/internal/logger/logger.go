package logger

import (
	"io"

	"github.com/rs/zerolog"
)

type Logger interface {
	io.Writer

	Info(message string, data interface{})
	Error(message string, data interface{})
	Fatal(message string, data interface{})
}

type ZeroLogLogger struct {
	logger *zerolog.Logger
}

func NewZeroLogLogger(writer io.Writer) *ZeroLogLogger {
	logger := zerolog.New(writer).With().Timestamp().Logger()

	return &ZeroLogLogger{
		logger: &logger,
	}
}

func (l *ZeroLogLogger) Write(p []byte) (n int, err error) {
	l.logger.Error().Msg(string(p))
	return len(p), nil
}

func (l *ZeroLogLogger) Info(message string, data interface{}) {
	l.logger.Info().Interface("data", data).Msg(message)
}

func (l *ZeroLogLogger) Error(message string, data interface{}) {
	l.logger.Error().Interface("data", data).Msg(message)
}

func (l *ZeroLogLogger) Fatal(message string, data interface{}) {
	l.logger.Fatal().Interface("data", data).Msg(message)
}
