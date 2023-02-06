package Created

import (
	"context"
	"github.com/bordunosp/ddd/example/app/user/domain/event"
	"log"
)

func SaveLogHandler(ctx context.Context, event event.UserCreated) error {
	log.Print(event.Id.String())

	return nil
}
