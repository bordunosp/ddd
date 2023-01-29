package DI

type ServiceItem[T any] struct {
	IsSingleton     bool
	ServiceName     string
	ServiceInitFunc ServiceInitFunc
}
