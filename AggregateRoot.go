package ddd

import (
	"errors"
	"github.com/google/uuid"
)

func NewAggregateRoot(id uuid.UUID) IAggregateRoot {
	return &aggregateRoot{
		id: id,
	}
}

type aggregateRoot struct {
	id uuid.UUID

	domainEvents []IEvent
}

func (a *aggregateRoot) ID() string {
	return a.id.String()
}

func (a *aggregateRoot) UUID() uuid.UUID {
	return a.id
}

func (a *aggregateRoot) AddDomainEvent(event IEvent) error {
	if event.AggregateID() != a.ID() {
		return errors.New("unexpected AggregateID")
	}

	a.domainEvents = append(a.domainEvents, event)

	return nil
}

func (a *aggregateRoot) DomainEvents() []IEvent {
	return a.domainEvents
}
