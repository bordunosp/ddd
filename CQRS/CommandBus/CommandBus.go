package CommandBus

import (
	"context"
	"errors"
	"github.com/bordunosp/ddd/CQRS"
	"github.com/bordunosp/ddd/CQRS/Middleware"
	"sync"
)

var ErrCommandAlreadyRegistered = errors.New("command already registered")
var ErrCommandNotRegistered = errors.New("command not registered")
var ErrCommandHandlerType = errors.New("ICommandHandler has incorrect type")
var ErrCommandMiddlewareType = errors.New("MiddlewareType has incorrect type")

var middlewares []any
var middlewaresMutex sync.Mutex
var registeredCommands = &sync.Map{}

func Register[T ICommand](handler ICommandHandler[T]) error {
	var command T

	if _, ok := registeredCommands.Load(command.CommandConfig().Name); ok {
		return ErrCommandAlreadyRegistered
	}

	registeredCommands.Store(command.CommandConfig().Name, handler)
	return nil
}

func RegisterMiddleware[T ICommand](middleware ICommandMiddleWare[T]) {
	middlewaresMutex.Lock()
	defer middlewaresMutex.Unlock()

	middlewares = append(middlewares, middleware)
}

func Execute[T ICommand](ctx context.Context, command T) (err error) {
	handler, ok := registeredCommands.Load(command.CommandConfig().Name)
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

	if command.CommandConfig().Sanitize {
		err = Middleware.Sanitize(ctx, &command)
		if err != nil {
			return
		}
	}

	if command.CommandConfig().Validate {
		err = Middleware.Validate(command)
		if err != nil {
			return
		}
	}

	for _, middleware := range middlewares {
		typedMiddleware, ok := middleware.(ICommandMiddleWare[T])
		if !ok {
			return ErrCommandMiddlewareType
		}

		typedHandler = typedMiddleware(typedHandler)
	}

	err = typedHandler(ctx, command)
	return
}

func ExecuteOrPanic[T ICommand](ctx context.Context, command T) {
	if err := Execute[T](ctx, command); err != nil {
		panic(err)
	}
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

func ExecuteAsyncAwaitOrPanic[T ICommand](ctx context.Context, command T) {
	if err := ExecuteAsyncAwait[T](ctx, command); err != nil {
		panic(err)
	}
}
