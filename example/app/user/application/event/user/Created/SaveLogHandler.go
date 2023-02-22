package Created

import (
	"context"
	"github.com/bordunosp/ddd/example/app/user/domain/event"
	"gorm.io/gorm"
	"log"
)

func SaveLogHandler(ctx context.Context, tx *gorm.DB, event event.UserCreated) error {
	log.Print("Print name from SaveLogHandler: ", event.Name)
	return nil
}
