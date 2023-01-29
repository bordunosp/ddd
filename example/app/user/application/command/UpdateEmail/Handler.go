package UpdateEmail

import (
	"context"
	"errors"
	"github.com/bordunosp/ddd/CQRS/CommandBus"
	"log"
)

func Handler(ctx context.Context, commandAny CommandBus.ICommand) error {
	request, ok := commandAny.(Command)
	if !ok {
		return errors.New("Incorrect CommandType: " + commandAny.CommandName())
	}

	log.Print(request.Email)

	return nil
}
