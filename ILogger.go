package ddd

import "io"

type ILogger interface {
	With(key, val string) ILogger
	LoggerForDb() io.Writer

	Info(msg string)
	Infof(format string, v ...any)
	ErrorStack(err error, msg string)
	FatalStack(err error, msg string)
	PanicStack(err error, msg string)
}
