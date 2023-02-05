package ddd

type iAggregateRootType IAggregateRoot
type iEntityType IEntity

type IAggregateOrEntity interface {
	iAggregateRootType | iEntityType
}

type IModelGorm[E IAggregateOrEntity] interface {
	FromModelGorm() (E, error)
	ToModelGorm(entity E) IModelGorm[E]
}
