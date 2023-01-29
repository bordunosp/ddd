package EventBus

type IEvent interface {
	EventName() string
}
