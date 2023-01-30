package EventBus

import (
	"context"
	"errors"
	"github.com/bordunosp/ddd/CQRS"
)

var ErrEventAlreadyRegistered = errors.New("event already registered")

var registeredEvents = make(map[string][]IEventHandler)

func RegisterEvents(eventItems []EventItem) error {
	for _, eventItem := range eventItems {
		if _, ok := registeredEvents[eventItem.EventName]; ok {
			return ErrEventAlreadyRegistered
		}
		registeredEvents[eventItem.EventName] = eventItem.Handlers
	}
	return nil
}

func Dispatch(ctx context.Context, event IEvent) (err error) {
	handlers, ok := registeredEvents[event.EventName()]
	if !ok {
		return nil
	}

	defer func() {
		if _err := CQRS.RecoverToError(recover()); _err != nil {
			err = _err
		}
	}()

	for _, handler := range handlers {
		err = handler(ctx, event)
		if err != nil {
			return
		}
	}

	return
}

func DispatchAsync(ctx context.Context, event IEvent) chan error {
	c := make(chan error)

	go func(ctx context.Context, event IEvent) {
		defer close(c)
		c <- Dispatch(ctx, event)
	}(ctx, event)

	return c
}

func DispatchAsyncAwait(ctx context.Context, event IEvent) error {
	return <-DispatchAsync(ctx, event)
}
