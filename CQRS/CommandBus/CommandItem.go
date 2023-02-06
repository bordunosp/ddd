package CommandBus

type CommandItem[T ICommand] struct {
	CommandName string
	Handler     ICommandHandler[T]
}
