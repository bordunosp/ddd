package EventBus

import (
	"context"
	"errors"
	"github.com/bordunosp/ddd/CQRS"
	"github.com/bordunosp/ddd/CQRS/Middleware"
	"sync"
)

var ErrEventAlreadyRegistered = errors.New("event already registered")
var ErrEventHandlerType = errors.New("IEventHandler has incorrect type")

var registeredEvents = &sync.Map{}

func Register[T IEvent, Tx any](handlers []IEventHandler[T, Tx]) error {
	var event T

	if _, ok := registeredEvents.Load(event.EventConfig().Name); ok {
		return ErrEventAlreadyRegistered
	}
	registeredEvents.Store(event.EventConfig().Name, handlers)
	return nil
}

func Dispatch[T IEvent, Tx any](ctx context.Context, tx Tx, event T) (err error) {
	handlers, ok := registeredEvents.Load(event.EventConfig().Name)
	if !ok {
		return nil
	}

	defer func() {
		if _err := CQRS.RecoverToError(recover()); _err != nil {
			err = _err
		}
	}()

	typedHandlers, ok := handlers.([]IEventHandler[T, Tx])
	if !ok {
		return ErrEventHandlerType
	}

	if event.EventConfig().Sanitize {
		err = Middleware.Sanitize(ctx, &event)
		if err != nil {
			return
		}
	}

	if event.EventConfig().Validate {
		err = Middleware.Validate(event)
		if err != nil {
			return
		}
	}

	for _, handler := range typedHandlers {
		err = handler(ctx, tx, event)
		if err != nil {
			return
		}
	}

	return
}

func DispatchAsync[T IEvent, Tx any](ctx context.Context, tx Tx, event T) chan error {
	c := make(chan error)

	go func(ctx context.Context, event T) {
		defer close(c)
		c <- Dispatch[T](ctx, tx, event)
	}(ctx, event)

	return c
}

func DispatchAsyncAwait[T IEvent, Tx any](ctx context.Context, tx Tx, event T) error {
	return <-DispatchAsync[T](ctx, tx, event)
}
