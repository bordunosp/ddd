package ddd

import "github.com/google/uuid"

type IEvent interface {
	Name() string
	AggregateID() uuid.UUID
}
