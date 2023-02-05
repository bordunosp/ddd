package Assertion

import (
	"errors"
)

func Same[T comparable](val1, val2 T, msg string) {
	if val1 == val2 {
		panic(errors.New(msg))
	}
}
