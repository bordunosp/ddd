package EventBus

import (
	"context"
	"errors"
	"github.com/bordunosp/ddd/CQRS"
	"sync"
)

var ErrEventAlreadyRegistered = errors.New("event already registered")
var ErrEventHandlerType = errors.New("IEventHandler has incorrect type")

var registeredEvents = &sync.Map{}

func Register[T IEvent](handlers []IEventHandler[T]) error {
	var event T

	if _, ok := registeredEvents.Load(event.EventName()); ok {
		return ErrEventAlreadyRegistered
	}
	registeredEvents.Store(event.EventName(), handlers)
	return nil
}

func Dispatch[T IEvent](ctx context.Context, event T) (err error) {
	handlers, ok := registeredEvents.Load(event.EventName())
	if !ok {
		return nil
	}

	defer func() {
		if _err := CQRS.RecoverToError(recover()); _err != nil {
			err = _err
		}
	}()

	typedHandlers, ok := handlers.([]IEventHandler[T])
	if !ok {
		return ErrEventHandlerType
	}

	for _, handler := range typedHandlers {
		err = handler(ctx, event)
		if err != nil {
			return
		}
	}

	return
}

func DispatchAsync[T IEvent](ctx context.Context, event T) chan error {
	c := make(chan error)

	go func(ctx context.Context, event T) {
		defer close(c)
		c <- Dispatch[T](ctx, event)
	}(ctx, event)

	return c
}

func DispatchAsyncAwait[T IEvent](ctx context.Context, event T) error {
	return <-DispatchAsync[T](ctx, event)
}
