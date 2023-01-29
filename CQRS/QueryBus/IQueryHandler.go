package QueryBus

import "context"

type IQueryHandler func(ctx context.Context, queryAny IQuery) (any, error)
