package QueryBus

import "context"

type IQueryHandler[T IQuery, K any] func(ctx context.Context, queryAny T) (K, error)
