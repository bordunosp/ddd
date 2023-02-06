package EventBus

import "context"

type IEventHandler[T IEvent] func(ctx context.Context, eventAny T) error
