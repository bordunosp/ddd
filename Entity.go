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

func (e *entity) IsEqual(other IEntity) bool {
	return other.ID() == e.ID()
}
