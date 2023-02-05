package Assertion

import (
	"errors"
	"github.com/bordunosp/ddd/Assertion/types"
)

func Range[T types.IntegerOrFloat](value, minValue, maxValue T, msg string) {
	if value < minValue || value > maxValue {
		panic(errors.New(msg))
	}
}
