package CreateNew

import (
	"context"
	"log"
)

func Handler(ctx context.Context, command Command) error {
	log.Print("Print from command Handler: ", command.Name)

	return nil
}
