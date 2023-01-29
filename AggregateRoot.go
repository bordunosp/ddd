package ddd

import (
	"github.com/google/uuid"
)

func NewAggregateRoot(id uuid.UUID) IAggregateRoot {
	return &aggregateRoot{
		id: id,
	}
}

type aggregateRoot struct {
	id uuid.UUID

	domainEvents []IDomainEvent
}

func (a *aggregateRoot) ID() string {
	return a.id.String()
}

func (a *aggregateRoot) UUID() uuid.UUID {
	return a.id
}

func (a *aggregateRoot) AddDomainEvent(event IDomainEvent) {
	a.domainEvents = append(a.domainEvents, event)
}

func (a *aggregateRoot) DomainEvents() []IDomainEvent {
	return a.domainEvents
}
