package Assertion

import "errors"

func False(val bool, msg string) {
	if val {
		panic(errors.New(msg))
	}
}
