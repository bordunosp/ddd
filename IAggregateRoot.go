package ddd

import "github.com/google/uuid"

type IAggregateRoot interface {
	ID() string
	UUID() uuid.UUID
	AddDomainEvent(event IEvent)
	DomainEvents() []IEvent
}
