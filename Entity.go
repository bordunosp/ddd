package ddd

import "github.com/google/uuid"

type Entity struct {
	id          uuid.UUID
	aggregateID uuid.UUID
}

func (e *Entity) ID() string {
	return e.id.String()
}

func (e *Entity) UUID() uuid.UUID {
	return e.id
}

func (e *Entity) AggregateID() string {
	return e.aggregateID.String()
}

func (e *Entity) AggregateUUID() uuid.UUID {
	return e.aggregateID
}
