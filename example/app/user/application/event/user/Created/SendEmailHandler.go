package Created

import (
	"context"
	"errors"
	"github.com/bordunosp/ddd/CQRS/EventBus"
	"github.com/bordunosp/ddd/DI"
	"github.com/bordunosp/ddd/example/app/user/domain"
	"github.com/bordunosp/ddd/example/app/user/domain/event"
)

func SendEmailHandler(ctx context.Context, eventAny EventBus.IEvent) error {
	request, ok := eventAny.(event.UserCreated)
	if !ok {
		return errors.New("Incorrect EventType: " + eventAny.EventName())
	}

	var userService domain.IUserService
	userService, err := DI.Get("UserService", userService)
	if err != nil {
		return err
	}

	return userService.SendCreatedEmail(ctx, request.Id)
}
