package Assertion

import (
	"errors"
	"unicode/utf8"
)

func RangeLength(value string, minLength, maxLength int, msg string) {
	cnt := utf8.RuneCountInString(value)

	if cnt < minLength || cnt > maxLength {
		panic(errors.New(msg))
	}
}
