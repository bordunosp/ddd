package CommandBus

type CommandItem struct {
	CommandName string
	Handler     ICommandHandler
}
