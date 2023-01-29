package EventBus

import "context"

type IEventHandler func(ctx context.Context, eventAny IEvent) error
