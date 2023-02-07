package ddd

import (
	"github.com/google/uuid"
)

func NewEntity(id, aggregateID uuid.UUID) IEntity {
	return &entity{
		id:          id,
		aggregateID: aggregateID,
	}
}

type entity struct {
	id          uuid.UUID
	aggregateID uuid.UUID
}

func (e *entity) ID() string {
	return e.id.String()
}

func (e *entity) UUID() uuid.UUID {
	return e.id
}

func (e *entity) AggregateID() string {
	return e.aggregateID.String()
}

func (e *entity) AggregateUUID() uuid.UUID {
	return e.aggregateID
}
