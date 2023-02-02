package ddd

type ILogger interface {
	Info(msg string)
	Infof(format string, v ...any)
	ErrorStack(err error, msg string)
	FatalStack(err error, msg string)
	PanicStack(err error, msg string)

	With(key, val string) ILogger

	AssertErrNotNil(err error, msg string)
	AssertErrNotNilFatal(err error, msg string)
	AssertErrNotNilPanic(err error, msg string)
}