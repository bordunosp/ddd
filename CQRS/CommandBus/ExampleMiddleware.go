package CommandBus

import (
	"context"
)

func exampleMiddleware[T ICommand](handler ICommandHandler[T]) ICommandHandler[T] {
	return func(ctx context.Context, command T) error {
		return handler(ctx, command)
	}
}
