package Assertion

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func Uuid(value string, msg string) {
	_, err := uuid.Parse(value)

	if err != nil {
		panic(errors.Wrap(err, msg))
	}

}
