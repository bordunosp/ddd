package Assertion

import (
	"errors"
	"github.com/bordunosp/ddd/Assertion/types"
)

func Max[T types.IntegerOrFloat](value, maxValue T, msg string) {
	if value > maxValue {
		panic(errors.New(msg))
	}
}
