
## Event (DTO)
###### src: app/user/domain/event/UserCreated.go
```golang
package event

import (
    "github.com/google/uuid"
    "github.com/bordunosp/ddd"
    "github.com/bordunosp/ddd/CQRS/EventBus"
)

type UserCreated struct {
    Id    uuid.UUID
    Name  string `mod:"trim" validate:"required"`
    Email string `mod:"trim" validate:"required,email"`
}

// implemented interface EventBus.IEvent 
// (need uniq name for registration)
func (e UserCreated) EventConfig() EventBus.EventConfig {
    return EventBus.EventConfig{
        Name:     "UserCreatedEvent",
        Sanitize: true,
        Validate: true}
}
```

## Event Handlers
###### src: app/user/application/event/user/Created/SendEmailHandler.go
```golang
package Created

import (
    "context"
    "github.com/bordunosp/ddd/DI"
    "github.com/bordunosp/ddd/example/app/user/domain"
    "github.com/bordunosp/ddd/example/app/user/domain/event"
    "log"
)

// For One event can be many handles

func SendEmailHandler(ctx context.Context, event event.UserCreated) error {
    // todo: send email
    log.Print("Print email from SendEmailHandler: ", event.Email)
    return nil
}

func SaveLogHandler(ctx context.Context, event event.UserCreated) error {
    // todo: save logs
    log.Print("Print name from SaveLogHandler: ", event.Name)
    return nil
}


```

## Event vs AggregateRoot
```golang
user, _ := NewUser(uuid.New(), "name", "email@test.com")
event := NewCreatedEvent(user.UUID())
_ := user.AddDomainEvent(event)

// DomainEvents must be executed before or after the transaction is committed
// but always in one transaction

// 1) open transaction
// 2) Save user to DB
// 3) Execute DomainEvents user.DomainEvents()
// 4) commit transaction 
```


## Event vs EventBus (Register)
```golang
package main

import (
    "github.com/bordunosp/ddd/CQRS/EventBus"
    "github.com/bordunosp/ddd/example/app/user/application/event/user/created"
    "github.com/bordunosp/ddd/example/app/user/domain/event"
    "log"
)

func main() {
    err := EventBus.Register([]EventBus.IEventHandler[event.UserCreated]{
        Created.SaveLogHandler,
        Created.SendEmailHandler,
    })
    if err != nil {
        log.Fatal(err)
    }

    err = EventBus.Register([]EventBus.IEventHandler[event.UserUpdated]{
        // event may not have handlers
        //
        // you never know when it might be really useful
        // that is why events are created long before handlers are created.
    })
    if err != nil {
        log.Fatal(err)
    }
}
```

## Event vs EventBus (Usage)
```golang
import (
    "github.com/google/uuid"
    "github.com/bordunosp/ddd/CQRS/EventBus"
    domainEvent "github.com/bordunosp/ddd/example/app/user/domain/event"
)

func main() {
    ctx := context.TODO()
	
    event = event.UserCreated{
        Id:    uuid.New(),
        Name:  "some eventDTO name",
        Email: "some@event.email",
    }
	
    // choose 1 of 3 possible ways to execute handlers
    _ = EventBus.Dispatch(ctx, event)
    _ = EventBus.DispatchAsync(ctx, event)
    _ = EventBus.DispatchAsyncAwait(ctx, event)
}
```
