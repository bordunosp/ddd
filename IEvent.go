package ddd

type IEvent interface {
	Name() string
	AggregateID() string
}
