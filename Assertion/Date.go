package Assertion

import (
	"github.com/pkg/errors"
	"time"
)

func Date(dateString, format, msg string) {
	_, err := time.Parse(format, dateString)

	if err != nil {
		panic(errors.Wrap(err, msg))
	}
}
