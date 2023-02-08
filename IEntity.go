package ddd

import "github.com/google/uuid"

type IEntity[T IEntity[T]] interface {
	ID() string
	UUID() uuid.UUID
	IsEqual(other *T) bool
}
