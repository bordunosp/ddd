package event

import (
	"github.com/bordunosp/ddd/CQRS/EventBus"
	"github.com/google/uuid"
)

type UserCreated struct {
	Id    uuid.UUID
	Name  string
	Email string
}

func (e UserCreated) EventConfig() EventBus.EventConfig {
	return EventBus.EventConfig{
		Name:     "UserCreatedEvent",
		Sanitize: true,
		Validate: true}
}
