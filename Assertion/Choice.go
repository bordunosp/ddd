package Assertion

import "errors"

func Choice[T comparable](value T, values []T, msg string) {
	contains := false

	for _, a := range values {
		if a == value {
			contains = true
			break
		}
	}

	if !contains {
		panic(errors.New(msg))
	}
}
