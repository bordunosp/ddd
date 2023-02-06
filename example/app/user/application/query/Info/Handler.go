package Info

import (
	"context"
	"log"
)

func Handler(ctx context.Context, query Query) (Response, error) {
	log.Print(query.Id)

	return NewResponse(
		"name",
		"email",
	), nil
}
