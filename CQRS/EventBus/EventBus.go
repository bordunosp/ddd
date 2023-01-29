package EventBus

import (
	"errors"
	"fmt"
)

var registeredEvents = make(map[string][]IEventHandler)

func RegisterEvents(eventItems []EventItem) error {
	for _, eventItem := range eventItems {
		if _, ok := registeredEvents[eventItem.EventName]; ok {
			return errors.New(fmt.Sprintf("event by name '%s' already registered", eventItem.EventName))
		}
		registeredEvents[eventItem.EventName] = eventItem.Handlers
	}
	return nil
}
