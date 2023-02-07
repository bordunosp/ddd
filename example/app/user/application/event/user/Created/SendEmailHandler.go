package Created

import (
	"context"
	"github.com/bordunosp/ddd/DI"
	"github.com/bordunosp/ddd/example/app/user/domain"
	"github.com/bordunosp/ddd/example/app/user/domain/event"
	"log"
)

func SendEmailHandler(ctx context.Context, event event.UserCreated) error {
	userService, err := DI.Get[domain.IUserService]("UserService")
	if err != nil {
		return err
	}

	log.Print("Print name from SendEmailHandler: ", event.Email)
	return userService.SendCreatedEmail(ctx, event.Id)
}
