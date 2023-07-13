package Saga

import (
	"context"
	"fmt"
	"github.com/sethvargo/go-retry"
	"time"
)

func main() {
	err := ExecuteTransactional(&ExampleSaga{
		Id:    "Some id",
		Email: "some@email.com",
	})

	if err != nil {
		panic(err)
	}
}

type FirstStep struct {
	Email string
}

func (f FirstStep) Name() string {
	return "FirstStep"
}

func (f FirstStep) RetryBackoff() retry.Backoff {
	return retry.WithMaxRetries(3, retry.NewFibonacci(1*time.Second))
}

func (f FirstStep) Forward() error {
	fmt.Println("call Forward " + f.Email)
	return nil
}

func (f FirstStep) Backward() error {
	fmt.Println("call Backward")
	return nil
}

type SecondStep struct {
	Id string
}

func (f SecondStep) Name() string {
	return "FirstStep"
}

func (f SecondStep) RetryBackoff() retry.Backoff {
	return retry.WithMaxRetries(3, retry.NewFibonacci(1*time.Second))
}

func (f SecondStep) Forward() error {
	fmt.Println("call Forward " + f.Id)
	return nil
}

func (f SecondStep) Backward() error {
	fmt.Println("call Backward")
	return nil
}

type ExampleSaga struct {
	Id    string
	Email string
}

func (s *ExampleSaga) GetName() string {
	return "ExampleSaga"
}
func (s *ExampleSaga) GetSteps() []ISagaStep {
	return []ISagaStep{
		&FirstStep{Email: s.Email},
		&SecondStep{Id: s.Id},
	}
}
func (s *ExampleSaga) onRollbackError(err error, keyStep int) {
	panic(err)
}

func ExecuteChoreography(saga ISaga) error {
	done := make(chan error)
	var err error
	steps := saga.GetSteps()

	for i := 0; i < len(steps); i++ {
		go func(step ISagaStep, keyStep int) {
			if err := retry.Do(context.Background(), steps[keyStep].RetryBackoff(), func(ctx context.Context) error {
				if err := steps[keyStep].Forward(); err != nil {
					return retry.RetryableError(err)
				}

				return nil
			}); err != nil {
				done <- err
				return
			}

			done <- nil
		}(steps[i], i)
	}

	for range steps {
		if err == nil {
			err = <-done
		} else {
			<-done
		}
	}

	return err
}

func ExecuteTransactional(saga ISaga) error {
	steps := saga.GetSteps()

	for i := 0; i < len(steps); i++ {
		if err := retry.Do(context.Background(), steps[i].RetryBackoff(), func(ctx context.Context) error {
			if err := steps[i].Forward(); err != nil {
				return retry.RetryableError(err)
			}

			return nil
		}); err != nil {
			rollbackTransactional(saga, i)
			return err
		}
	}

	return nil
}

func rollbackChoreography(saga ISaga, keyStep int) {
	steps := saga.GetSteps()

	if err := retry.Do(context.Background(), steps[keyStep].RetryBackoff(), func(ctx context.Context) error {
		if err := steps[keyStep].Backward(); err != nil {
			return retry.RetryableError(err)
		}

		return nil
	}); err != nil {
		saga.onRollbackError(err, keyStep)
	}
}

func rollbackTransactional(saga ISaga, keyStep int) {
	steps := saga.GetSteps()

	for i := keyStep; i >= 0; i-- {
		if err := retry.Do(context.Background(), steps[i].RetryBackoff(), func(ctx context.Context) error {
			if err := steps[i].Backward(); err != nil {
				return retry.RetryableError(err)
			}

			return nil
		}); err != nil {
			saga.onRollbackError(err, i)
		}
	}
}
