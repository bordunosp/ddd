package ddd

import (
	"github.com/bordunosp/ddd/CQRS/EventBus"
	"github.com/google/uuid"
)

type IAggregateRoot interface {
	ID() string
	UUID() uuid.UUID
	AddDomainEvent(event EventBus.IEvent)
	DomainEvents() []EventBus.IEvent
}
