package QueryBus

import (
	"context"
	"errors"
	"github.com/bordunosp/ddd/CQRS"
	"sync"
)

var ErrQueryAlreadyRegistered = errors.New("query already registered")
var ErrQueryNotRegistered = errors.New("query not registered")
var ErrQueryHandlerType = errors.New("IQueryHandler has incorrect type")

var registeredQueries = &sync.Map{}

func Register[T IQuery, K any](handler IQueryHandler[T, K]) error {
	var query T

	if _, ok := registeredQueries.Load(query.QueryName()); ok {
		return ErrQueryAlreadyRegistered
	}

	registeredQueries.Store(query.QueryName(), handler)
	return nil
}

func Handle[K any, T IQuery](ctx context.Context, query T) (value K, err error) {
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

func HandleAsync[K any, T IQuery](ctx context.Context, query T) chan ReplayDTO[K] {
	replay := make(chan ReplayDTO[K])

	go func(query T) {
		defer close(replay)
		value, err := Handle[K, T](ctx, query)
		replay <- *&ReplayDTO[K]{Value: value, Err: err}
	}(query)

	return replay
}

func HandleAsyncAwait[K any, T IQuery](ctx context.Context, query T) (K, error) {
	replay := <-HandleAsync[K, T](ctx, query)
	return replay.Value, replay.Err
}
