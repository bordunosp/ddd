package ddd

import "github.com/google/uuid"

type IEntity interface {
	ID() string
	UUID() uuid.UUID
	AggregateID() string
	AggregateUUID() uuid.UUID
}
