package DI

type ServiceItem struct {
	IsSingleton     bool
	ServiceName     string
	ServiceInitFunc ServiceInitFunc
}
