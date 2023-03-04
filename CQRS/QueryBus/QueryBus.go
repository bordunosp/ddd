package QueryBus

import (
	"context"
	"errors"
	"github.com/bordunosp/ddd/CQRS"
	"github.com/bordunosp/ddd/CQRS/Middleware"
	"sync"
)

var ErrQueryAlreadyRegistered = errors.New("query already registered")
var ErrQueryNotRegistered = errors.New("query not registered")
var ErrQueryHandlerType = errors.New("IQueryHandler has incorrect type")

var registeredQueries = &sync.Map{}

func Register[T IQuery, K any](handler IQueryHandler[T, K]) error {
	var query T

	if _, ok := registeredQueries.Load(query.QueryConfig().Name); ok {
		return ErrQueryAlreadyRegistered
	}

	registeredQueries.Store(query.QueryConfig().Name, handler)
	return nil
}

func Handle[K any, T IQuery](ctx context.Context, query T) (value K, err error) {
	handler, ok := registeredQueries.Load(query.QueryConfig().Name)
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

	if query.QueryConfig().Sanitize {
		err = Middleware.Sanitize(ctx, &query)
		if err != nil {
			return value, err
		}
	}

	if query.QueryConfig().Validate {
		err = Middleware.Validate(query)
		if err != nil {
			return
		}
	}

	value, err = typedHandler(ctx, query)

	if query.QueryConfig().SanitizeResponse {
		err = Middleware.Sanitize(ctx, &value)
		if err != nil {
			return value, err
		}
	}

	if query.QueryConfig().SanitizeResponse {
		err = Middleware.Validate(value)
		if err != nil {
			return
		}
	}

	return
}

func HandleOrPanic[K any, T IQuery](ctx context.Context, query T) K {
	val, err := Handle[K, T](ctx, query)
	if err != nil {
		panic(err)
	}

	return val
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

func HandleAsyncAwaitOrPanic[K any, T IQuery](ctx context.Context, query T) K {
	val, err := HandleAsyncAwait[K, T](ctx, query)
	if err != nil {
		panic(err)
	}

	return val
}
