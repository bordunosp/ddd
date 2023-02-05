package Assertion

import (
	"errors"
	"unicode/utf8"
)

func MinLength(value string, minLength int, msg string) {
	cnt := utf8.RuneCountInString(value)

	if cnt < minLength {
		panic(errors.New(msg))
	}
}
