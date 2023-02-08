package ddd

import "github.com/google/uuid"

type Entity interface {
	ID() string
	UUID() uuid.UUID
}

func NewEntity(id, aggregateID uuid.UUID) Entity {
	return &entity{
		id:          id,
		aggregateID: aggregateID,
	}
}

type entity struct {
	id          uuid.UUID
	aggregateID uuid.UUID
}

func (e entity) ID() string {
	return e.id.String()
}

func (e entity) UUID() uuid.UUID {
	return e.id
}
