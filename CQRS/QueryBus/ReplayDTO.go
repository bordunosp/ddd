package QueryBus

type ReplayDTO[K any] struct {
	Value K
	Err   error
}
