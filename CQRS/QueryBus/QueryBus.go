package QueryBus

import (
	"context"
	"errors"
	"github.com/bordunosp/ddd/CQRS"
	"sync"
)

var ErrQueryAlreadyRegistered = errors.New("query already registered")
var ErrQueryNotRegistered = errors.New("query not registered")
var ErrQueryHandlerType = errors.New("IQueryHandler has incorrect types")

var registeredQueries = &sync.Map{}

func RegisterQuery[T IQuery, K any](queryItem QueryItem[T, K]) error {
	if _, ok := registeredQueries.Load(queryItem.QueryName); ok {
		return ErrQueryAlreadyRegistered
	}
	registeredQueries.Store(queryItem.QueryName, queryItem.Handler)
	return nil
}

func RegisterQueries[T IQuery, K any](queryItems []QueryItem[T, K]) error {
	for _, queryItem := range queryItems {
		if _, ok := registeredQueries.Load(queryItem.QueryName); ok {
			return ErrQueryAlreadyRegistered
		}
		registeredQueries.Store(queryItem.QueryName, queryItem.Handler)
	}
	return nil
}

func Handle[T IQuery, K any](ctx context.Context, query T) (value K, err error) {
	handler, ok := registeredQueries.Load(query.QueryName())
	if !ok {
		return value, ErrQueryNotRegistered
	}

	defer func() {
		if _err := CQRS.RecoverToError(recover()); _err != nil {
			err = _err
		}
	}()

	typedHandler, ok := handler.(IQueryHandler[T, K])
	if !ok {
		return value, ErrQueryHandlerType
	}

	value, err = typedHandler(ctx, query)
	return
}

func HandleAsync[T IQuery, K any](ctx context.Context, query T) chan ReplayDTO[K] {
	replay := make(chan ReplayDTO[K])

	go func(query T) {
		defer close(replay)
		value, err := Handle[T, K](ctx, query)
		replay <- *&ReplayDTO[K]{Value: value, Err: err}
	}(query)

	return replay
}

func HandleAsyncAwait[T IQuery, K any](ctx context.Context, query T) (K, error) {
	replay := <-HandleAsync[T, K](ctx, query)
	return replay.Value, replay.Err
}
