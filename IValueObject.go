package ddd

type IValueObject interface {
	IsEqual(other IValueObject) bool
}
