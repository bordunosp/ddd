package Created

import (
	"context"
	"errors"
	"github.com/bordunosp/ddd/CQRS/EventBus"
	"github.com/bordunosp/ddd/example/app/user/domain/event"
	"log"
)

func SaveLogHandler(ctx context.Context, eventAny EventBus.IEvent) error {
	request, ok := eventAny.(event.UserCreated)
	if !ok {
		return errors.New("Incorrect EventType: " + eventAny.EventName())
	}

	log.Print(request.Id.String())

	return nil
}
