package CommandBus

import "context"

type ICommandHandler func(ctx context.Context, commandAny ICommand) error
