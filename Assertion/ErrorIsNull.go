package Assertion

import "github.com/pkg/errors"

func ErrorIsNull(err error, msg string) {
	if err != nil {
		panic(errors.Wrap(err, msg))
	}
}
