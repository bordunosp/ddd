package Assertion

import "errors"

func NotNull(value any, msg string) {
	if value == nil {
		panic(errors.New(msg))
	}
}
