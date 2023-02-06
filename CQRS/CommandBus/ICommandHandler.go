package CommandBus

import "context"

type ICommandHandler[T ICommand] func(ctx context.Context, commandAny T) error
