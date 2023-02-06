package CreateNew

import (
	"context"
	"log"
)

func Handler(ctx context.Context, command Command) error {
	log.Print(command.Name)

	return nil
}
