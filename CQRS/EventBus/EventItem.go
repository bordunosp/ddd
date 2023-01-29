package EventBus

type EventItem struct {
	EventName string
	Handlers  []IEventHandler
}
