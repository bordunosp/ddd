package QueryBus

import (
	"context"
	"errors"
	"github.com/bordunosp/ddd/CQRS"
)

var ErrQueryAlreadyRegistered = errors.New("query already registered")
var ErrQueryNotRegistered = errors.New("query not registered")

var registeredQueries = make(map[string]IQueryHandler)

func RegisterQueries(queryItems []QueryItem) error {
	for _, queryItem := range queryItems {
		if _, ok := registeredQueries[queryItem.QueryName]; ok {
			return ErrQueryAlreadyRegistered
		}
		registeredQueries[queryItem.QueryName] = queryItem.Handler
	}
	return nil
}

func Handle(ctx context.Context, query IQuery) (value any, err error) {
	handler, ok := registeredQueries[query.QueryName()]
	if !ok {
		err = ErrQueryNotRegistered
		return
	}

	defer func() {
		if _err := CQRS.RecoverToError(recover()); _err != nil {
			err = _err
		}
	}()

	value, err = handler(ctx, query)
	return
}

func HandleAsync(ctx context.Context, query IQuery) chan ReplayDTO {
	replay := make(chan ReplayDTO)

	go func(query IQuery) {
		defer close(replay)
		value, err := Handle(ctx, query)
		replay <- *&ReplayDTO{Value: value, Err: err}
	}(query)

	return replay
}

func HandleAsyncAwait(ctx context.Context, query IQuery) (any, error) {
	replay := <-HandleAsync(ctx, query)
	return replay.Value, replay.Err
}
