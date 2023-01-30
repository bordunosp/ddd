package CQRS

import (
	"errors"
	"fmt"
)

func RecoverToError(recover any) (err error) {
	if recover == nil {
		return
	}

	switch x := recover.(type) {
	case string:
		err = errors.New(x)
	case error:
		err = x
	default:
		err = errors.New(fmt.Sprintf("unknown panic. Type: %s. Err: %s", x, err))
	}

	return
}
