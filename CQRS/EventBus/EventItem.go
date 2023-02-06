package EventBus

type EventItem[T IEvent] struct {
	EventName string
	Handlers  []IEventHandler[T]
}
