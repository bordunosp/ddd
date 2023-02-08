package ddd

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"os"
	"strconv"
	"strings"
	"time"
)

func NewLoggerFromZerolog(isProd bool, microServiceName string) (ILogger, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	zerolog.LevelFieldName = "lvl"
	zerolog.MessageFieldName = "msg"
	zerolog.TimestampFieldName = "utc"
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05"
	zerolog.TimestampFunc = time.Now().UTC
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return strings.Replace(file, pwd, "", 1) + ":" + strconv.Itoa(line)
	}

	if isProd {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log := zerolog.
		New(os.Stderr).
		With().
		Str("srv", microServiceName).
		Timestamp().
		Logger()

	return &loggerZerolog{
		log: log,
	}, nil
}

type loggerZerolog struct {
	log zerolog.Logger
}

// Printf usage for logs from Gorm ORM
func (l *loggerZerolog) Printf(s string, a ...any) {
	l.log.Printf(s, a...)
}

func (l *loggerZerolog) With(key, val string) ILogger {
	newLogger := *l
	newLogger.log = newLogger.log.With().Str(key, val).Logger()
	return &newLogger
}

func (l *loggerZerolog) Info(msg string) {
	l.log.Info().Msg(msg)
}

func (l *loggerZerolog) Infof(format string, v ...any) {
	l.log.Info().Msgf(format, v...)
}

func (l *loggerZerolog) ErrorStack(err error, msg string) {
	l.log.Error().Stack().Err(errors.WithStack(err)).Msg(msg)
}

func (l *loggerZerolog) FatalStack(err error, msg string) {
	l.log.Fatal().Stack().Err(errors.WithStack(err)).Msg(msg)
}

func (l *loggerZerolog) PanicStack(err error, msg string) {
	l.log.Panic().Stack().Err(errors.WithStack(err)).Msg(msg)
}
