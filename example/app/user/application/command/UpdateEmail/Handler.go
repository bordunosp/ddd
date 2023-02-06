package UpdateEmail

import (
	"context"
	"log"
)

func Handler(ctx context.Context, command Command) error {
	log.Print(command.Email)

	return nil
}
