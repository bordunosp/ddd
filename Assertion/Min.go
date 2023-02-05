package Assertion

import (
	"errors"
	"github.com/bordunosp/ddd/Assertion/types"
)

func Min[T types.IntegerOrFloat](value, minValue T, msg string) {
	if value < minValue {
		panic(errors.New(msg))
	}
}
