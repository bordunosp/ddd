package EventBus

import (
	"errors"
)

var ErrEventAlreadyRegistered = errors.New("event already registered")

var registeredEvents = make(map[string][]IEventHandler)

func RegisterEvents(eventItems []EventItem) error {
	for _, eventItem := range eventItems {
		if _, ok := registeredEvents[eventItem.EventName]; ok {
			return ErrEventAlreadyRegistered
		}
		registeredEvents[eventItem.EventName] = eventItem.Handlers
	}
	return nil
}
