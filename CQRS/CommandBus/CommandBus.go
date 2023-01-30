package CommandBus

import (
	"context"
	"errors"
	"github.com/bordunosp/ddd/CQRS"
)

var ErrCommandAlreadyRegistered = errors.New("command already registered")
var ErrCommandNotRegistered = errors.New("command not registered")

var registeredCommands = make(map[string]ICommandHandler)

func RegisterCommands(commandItems []CommandItem) error {
	for _, commandItem := range commandItems {
		if _, ok := registeredCommands[commandItem.CommandName]; ok {
			return ErrCommandAlreadyRegistered
		}
		registeredCommands[commandItem.CommandName] = commandItem.Handler
	}
	return nil
}

func Execute(ctx context.Context, command ICommand) (err error) {
	handler, ok := registeredCommands[command.CommandName()]
	if !ok {
		return ErrCommandNotRegistered
	}

	defer func() {
		if _err := CQRS.RecoverToError(recover()); _err != nil {
			err = _err
		}
	}()

	err = handler(ctx, command)
	return
}

func ExecuteAsync(ctx context.Context, command ICommand) chan error {
	c := make(chan error)

	go func(ctx context.Context, command ICommand) {
		defer close(c)
		c <- Execute(ctx, command)
	}(ctx, command)

	return c
}

func ExecuteAsyncAwait(ctx context.Context, command ICommand) error {
	return <-ExecuteAsync(ctx, command)
}
