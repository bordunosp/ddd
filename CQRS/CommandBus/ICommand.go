package CommandBus

type ICommand interface {
	CommandName() string
}
