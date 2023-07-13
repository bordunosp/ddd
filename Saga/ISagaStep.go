package Saga

import "github.com/sethvargo/go-retry"

type ISagaStep interface {
	Name() string
	RetryBackoff() retry.Backoff
	Forward() error
	Backward() error
}
