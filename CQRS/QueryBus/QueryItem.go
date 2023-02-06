package QueryBus

type QueryItem[T IQuery, K any] struct {
	QueryName string
	Handler   IQueryHandler[T, K]
}
