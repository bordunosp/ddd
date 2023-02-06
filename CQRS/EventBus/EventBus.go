package EventBus

import (
	"context"
	"errors"
	"github.com/bordunosp/ddd/CQRS"
	"sync"
)

var ErrEventAlreadyRegistered = errors.New("event already registered")
var ErrEventHandlerType = errors.New("IEventHandler has incorrect types")

var registeredEvents = &sync.Map{}

func RegisterEvent[T IEvent](eventItem EventItem[T]) error {
	if _, ok := registeredEvents.Load(eventItem.EventName); ok {
		return ErrEventAlreadyRegistered
	}
	registeredEvents.Store(eventItem.EventName, eventItem.Handlers)
	return nil
}

func RegisterEvents[T IEvent](eventItems []EventItem[T]) error {
	for _, eventItem := range eventItems {
		if _, ok := registeredEvents.Load(eventItem.EventName); ok {
			return ErrEventAlreadyRegistered
		}
		registeredEvents.Store(eventItem.EventName, eventItem.Handlers)
	}
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
