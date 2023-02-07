package CommandBus

import (
	"context"
	"errors"
	"github.com/bordunosp/ddd/CQRS"
	"sync"
)

var ErrCommandAlreadyRegistered = errors.New("command already registered")
var ErrCommandNotRegistered = errors.New("command not registered")
var ErrCommandHandlerType = errors.New("ICommandHandler has incorrect types")

var registeredCommands = &sync.Map{}

func Register[T ICommand](handler ICommandHandler[T]) error {
	var command T

	if _, ok := registeredCommands.Load(command.CommandName()); ok {
		return ErrCommandAlreadyRegistered
	}

	registeredCommands.Store(command.CommandName(), handler)
	return nil
}

func Execute[T ICommand](ctx context.Context, command T) (err error) {
	handler, ok := registeredCommands.Load(command.CommandName())
	if !ok {
		return ErrCommandNotRegistered
	}

	defer func() {
		if _err := CQRS.RecoverToError(recover()); _err != nil {
			err = _err
		}
	}()

	typedHandler, ok := handler.(ICommandHandler[T])
	if !ok {
		return ErrCommandHandlerType
	}

	err = typedHandler(ctx, command)
	return
}

func ExecuteAsync[T ICommand](ctx context.Context, command T) chan error {
	c := make(chan error)

	go func(ctx context.Context, command T) {
		defer close(c)
		c <- Execute[T](ctx, command)
	}(ctx, command)

	return c
}

func ExecuteAsyncAwait[T ICommand](ctx context.Context, command T) error {
	return <-ExecuteAsync[T](ctx, command)
}
