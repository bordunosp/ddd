package Assertion

import (
	"errors"
	"unicode/utf8"
)

func MaxLength(value string, maxLength int, msg string) {
	cnt := utf8.RuneCountInString(value)

	if cnt > maxLength {
		panic(errors.New(msg))
	}
}
