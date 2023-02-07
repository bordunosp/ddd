package ddd

import (
	"github.com/bordunosp/ddd/CQRS/EventBus"
	"github.com/google/uuid"
)

func NewAggregateRoot(id uuid.UUID) IAggregateRoot {
	return &aggregateRoot{
		id: id,
	}
}

type aggregateRoot struct {
	id uuid.UUID

	domainEvents []EventBus.IEvent
}

func (a *aggregateRoot) ID() string {
	return a.id.String()
}

func (a *aggregateRoot) UUID() uuid.UUID {
	return a.id
}

func (a *aggregateRoot) IsEqual(other IAggregateRoot) bool {
	return a.ID() == other.ID()
}

func (a *aggregateRoot) DomainEvents() []EventBus.IEvent {
	return a.domainEvents
}

func (a *aggregateRoot) DomainEventsAdd(event EventBus.IEvent) {
	a.domainEvents = append(a.domainEvents, event)
}

func (a *aggregateRoot) DomainEventsClear() {
	a.domainEvents = []EventBus.IEvent{}
}
