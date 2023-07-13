package Saga

type ISaga interface {
	GetName() string
	GetSteps() []ISagaStep
	onRollbackError(err error, keyStep int)
}
