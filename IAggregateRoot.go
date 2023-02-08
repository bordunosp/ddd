package ddd

import (
	"github.com/bordunosp/ddd/CQRS/EventBus"
	"github.com/google/uuid"
)

type IAggregateRoot[T IAggregateRoot[T]] interface {
	ID() string
	UUID() uuid.UUID
	DomainEvents() []EventBus.IEvent
	DomainEventsAdd(event EventBus.IEvent)
	DomainEventsClear()

	IsEqual(other T) bool
}
