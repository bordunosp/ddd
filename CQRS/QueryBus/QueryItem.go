package QueryBus

type QueryItem struct {
	QueryName string
	Handler   IQueryHandler
}
