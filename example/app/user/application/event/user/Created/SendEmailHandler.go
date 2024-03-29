package Created

import (
	"context"
	"github.com/bordunosp/ddd/DI"
	"github.com/bordunosp/ddd/example/app/user/domain"
	"github.com/bordunosp/ddd/example/app/user/domain/event"
	"gorm.io/gorm"
	"log"
)

func SendEmailHandler(ctx context.Context, tx *gorm.DB, event event.UserCreated) error {
	userService, err := DI.Get[domain.IUserService]()
	if err != nil {
		return err
	}

	log.Print("Print email from SendEmailHandler: ", event.Email)
	return userService.SendCreatedEmail(ctx, event.Id)
}
