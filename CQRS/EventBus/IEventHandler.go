package EventBus

import "context"

type IEventHandler[T IEvent, Tx any] func(ctx context.Context, tx Tx, eventAny T) error
