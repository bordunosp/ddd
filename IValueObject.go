package ddd

type IValueObject[T IValueObject[T]] interface {
	IsEqual(other T) bool
}
