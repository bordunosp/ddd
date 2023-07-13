## Saga
### The Saga architecture pattern provides transaction management using a sequence of local transactions.

A local transaction is the unit of work performed by a Saga participant. Every operation that is part of the Saga can be rolled back by a compensating transaction. Further, the Saga pattern guarantees that either all operations complete successfully or the corresponding compensation transactions are run to undo the work previously completed.

In the Saga pattern, a compensating transaction must be idempotent and retryable. These two principles ensure that we can manage transactions without any manual intervention.

---

## Choreography Saga

For Choreography Saga, all steps are executed asynchronously (not sequentially).

In the event of an error, the compensation function will only be executed for the step in the saga where the error occurred.

---

## Transactional Saga

All steps in the saga are executed in a specific order, one after the other.

If an error occurs in any step, the compensation function will be executed for all steps that have successfully executed before it.


---

## Example

### First Step
```go
package main

import (
    "fmt"
    "github.com/sethvargo/go-retry"
)


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
```

### Second Step
```go
package main

import (
    "fmt"
    "github.com/sethvargo/go-retry"
)

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
```

### Example Saga
```go
package main

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

```

### Call func
```go
package main

import (
    "github.com/bordunosp/ddd/Saga"
)

func main() {
    if err := Saga.ExecuteTransactional(&ExampleSaga{
        Id:    "Some id",
        Email: "some@email.com",
    }); err != nil {
        panic(err)
    }
	
    // or

    if err := Saga.ExecuteChoreography(&ExampleSaga{
        Id:    "Some id",
        Email: "some@email.com",
    }); err != nil {
        panic(err)
    }
}
```

