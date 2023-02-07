package ddd

type ILogger interface {
	With(key, val string) ILogger
	Info(msg string)
	Infof(format string, v ...any)
	ErrorStack(err error, msg string)
	FatalStack(err error, msg string)
	PanicStack(err error, msg string)
	Printf(string, ...any)
}
