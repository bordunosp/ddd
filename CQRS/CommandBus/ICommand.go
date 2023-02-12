package CommandBus

type ICommand interface {
	CommandConfig() CommandConfig
}
