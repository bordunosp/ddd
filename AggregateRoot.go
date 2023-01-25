package ddd

import "github.com/google/uuid"

type AggregateRoot struct {
	id uuid.UUID

	domainEvents []IEvent
}

func (a *AggregateRoot) ID() string {
	return a.id.String()
}

func (a *AggregateRoot) UUID() uuid.UUID {
	return a.id
}

func (a *AggregateRoot) AddDomainEvent(event IEvent) {
	a.domainEvents = append(a.domainEvents, event)
}

func (a *AggregateRoot) DomainEvents() []IEvent {
	return a.domainEvents
}
