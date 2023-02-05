package Assertion

import "errors"

func True(val bool, msg string) {
	if !val {
		panic(errors.New(msg))
	}
}
