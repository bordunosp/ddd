package CommandBus

type ICommandMiddleWare[T ICommand] func(handler ICommandHandler[T]) ICommandHandler[T]
