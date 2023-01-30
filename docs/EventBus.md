
## Event
###### src: app/user/domain/event/UserCreated.go
```golang
package event

import (
    "github.com/google/uuid"
    "github.com/bordunosp/ddd"
    "github.com/bordunosp/ddd/CQRS/EventBus"
)

const UserCreatedEvent = "UserCreatedEvent"

func NewUserCreated(id uuid.UUID, name, email string) EventBus.IEvent {
    return &UserCreated{
        Id:    id,
        Name:  name,
        Email: email,
    }
}

type UserCreated struct {
    Id    uuid.UUID
    Name  string
    Email string
}

func (u UserCreated) EventName() string {
    return UserCreatedEvent
}
```

## Event Handlers
###### src: app/user/application/event/user/Created/SendEmailHandler.go
```golang
package Created

import (
    "context"
    "github.com/bordunosp/ddd/CQRS/EventBus"
    "github.com/bordunosp/ddd/DI"
    "github.com/bordunosp/ddd/example/app/user/domain"
    "github.com/bordunosp/ddd/example/app/user/domain/event"
)

func SendEmailHandler(ctx context.Context, eventAny EventBus.IEvent) error {
    request, _ := eventAny.(event.UserCreated)

	var userService domain.IUserService
	userService, err := DI.Get("UserService", userService)
	if err != nil {
		return err
	}

    return userService.SendCreatedEmail(ctx, request.Id)
}
```

## Event vs AggregateRoot
```golang
user, _ := NewUser(uuid.New(), "name", "email@test.com")
event := NewCreatedEvent(user.UUID())
_ := user.AddDomainEvent(event)

// DomainEvents must be executed before the transaction is committed

// 1) open transaction
// 2) Execute DomainEvents user.DomainEvents()
// 3) Save user to DB
// 4) commit transaction 
```


## EventBus Register
```golang
package main

import (
    "github.com/bordunosp/ddd/CQRS/EventBus"
    "github.com/bordunosp/ddd/example/app/user/application/event/user/created"
    "github.com/bordunosp/ddd/example/app/user/domain/event"
    "log"
)

func main() {
    err := EventBus.RegisterEvents([]EventBus.EventItem{
        {
            EventName: event.UserCreatedEvent,
            Handlers: []EventBus.IEventHandler{
                Created.SaveLogHandler,
                Created.SendEmailHandler,
            },
        },

        {event.UserUpdatedEvent, []EventBus.IEventHandler{
            // event may not have handlers
            //
            // you never know when it might be really useful
            // that is why events are created long before handlers are created.
        }},
    })
    if err != nil {
        log.Fatal(err)
    }
}
```

## EventBus Usage
```golang
import (
    "github.com/google/uuid"
    "github.com/bordunosp/ddd/CQRS/EventBus"
    domainEvent "github.com/bordunosp/ddd/example/app/user/domain/event"
)

ctx := context.TODO()
event := domainEvent.NewUserCreated(uuid.New(), "John Doe", "john-doe@email.com")

// choose 1 of 3 possible ways to execute handlers
_ = EventBus.Dispatch(ctx, event)
_ = EventBus.DispatchAsync(ctx, event)
_ = EventBus.DispatchAsyncAwait(ctx, event)
```
