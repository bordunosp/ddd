package event

import (
	"github.com/bordunosp/ddd/CQRS/EventBus"
	"github.com/google/uuid"
)

type UserUpdated struct {
	Id    uuid.UUID
	Name  string
	Email string
}

func (e UserUpdated) EventConfig() EventBus.EventConfig {
	return EventBus.EventConfig{
		Name:     "UserUpdatedEvent",
		Sanitize: true,
		Validate: true}
}
