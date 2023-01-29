package Info

import (
	"context"
	"errors"
	"github.com/bordunosp/ddd/CQRS/QueryBus"
	"log"
)

func Handler(ctx context.Context, queryAny QueryBus.IQuery) (any, error) {
	request, ok := queryAny.(Query)
	if !ok {
		return nil, errors.New("Incorrect QueryType: " + queryAny.QueryName())
	}

	log.Print(request.Id)

	return NewResponse(
		"name",
		"email",
	), nil
}
