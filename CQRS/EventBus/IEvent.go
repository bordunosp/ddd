package EventBus

type IEvent interface {
	EventConfig() EventConfig
}
