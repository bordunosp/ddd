
## DI
```golang
package main

import (
    "github.com/bordunosp/ddd/DI"
    "log"
    "os"
)

func main() {
    err := DI.RegisterServices([]DI.ServiceItem{
        {
            // will be initialized immediately (once)
            IsSingleton: true,
            ServiceName: "logger",
            ServiceInitFunc: func() (any, error) {
                // create new simple logger
                return log.New(os.Stderr, "\t", log.Ldate|log.Ltime|log.Lshortfile), nil
            },
        },
        {
            // will be initialized many times (per each call)
            IsSingleton: false,
            ServiceName: "UserService",
            ServiceInitFunc: func() (any, error) {
                loggerAny, err := DI.Get("logger")
                if err != nil {
                    return nil, err
                }

                // creating new any other service which use logger from DI
                return infrastructure.NewUserService(loggerAny.(*log.Logger))
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    // Use service from DI
    // it can be used anywhere in your project (after registered)
    loggerAny, _ := DI.Get("logger")
    loggerAny.(*log.Logger).Println("loggerAny call")
}
```

---

## AggregateRoot
###### src: app/user/domain/user.go
```golang
package domain

import (
    "github.com/google/uuid"
    "github.com/bordunosp/ddd"
)

type User interface {
    ddd.IAggregateRoot

    Name() string
    Email() string
}

type user struct {
    ddd.IAggregateRoot
    
    name  string
    email string
}

func (u *user) Name() string {
    return u.name
}

func (u *user) Email() string {
    return u.email
}

func NewUser(id uuid.UUID, name, email string) (User, error) {
    user := &account{
        IAggregateRoot: ddd.NewAggregateRoot(id),
        name:           name,
        email:          email,
    }
    return user, nil
}
```

## Entity
###### src: app/user/domain/entity/profile.go
```golang
package entity

import (
    "github.com/google/uuid"
    "github.com/bordunosp/ddd"
)

type Profile interface {
    ddd.IEntity

    Name() string
    Email() string
}

type profile struct {
    ddd.IEntity
    
    name  string
    email string
}

func (p *profile) Name() string {
    return p.name
}

func (p *profile) Email() string {
    return p.email
}

func NewProfile(id, aggregateID uuid.UUID, name, email string) (Profile, error) {
    profile := &account{
        IEntity: ddd.NewEntity(id, aggregateID),
        name:    name,
        email:   email,
    }
    return profile, nil
}
```

---

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

    userServiceAny, err := DI.Get("UserService")
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

---

## Command
###### src: app/user/application/command/UpdateEmail/Command.go
```golang
package UpdateEmail

import (
    "github.com/bordunosp/ddd/CQRS/CommandBus"
    "github.com/google/uuid"
)

const CommandName = "UpdateEmailCommand"

func NewCommand(id uuid.UUID, email string) CommandBus.ICommand {
    return &Command{
        Id:    id,
        Email: email,
    }
}

type Command struct {
    Id    uuid.UUID
    Email string
}

func (c Command) CommandName() string {
    return CommandName
}
```

## Command Handler
###### src: app/user/application/command/UpdateEmail/Handler.go
```golang
package UpdateEmail

import (
    "context"
    "github.com/bordunosp/ddd/DI"
    "github.com/bordunosp/ddd/CQRS/CommandBus"
)

func Handler(ctx context.Context, commandAny CommandBus.ICommand) error {
    request, _ := commandAny.(Command)

    loggerAny, _ := DI.Get("logger")
    loggerAny.(*log.Logger).Println(request.Email)
	
    return nil
}
```

---

## Query
###### src: app/user/application/query/Info/Query.go
```golang
package Info

import (
    "github.com/bordunosp/ddd/CQRS/QueryBus"
    "github.com/google/uuid"
)

const QueryName = "InfoQuery"

func NewQuery(id uuid.UUID) QueryBus.IQuery {
    return &Query{
        Id: id,
    }
}

type Query struct {
    Id uuid.UUID
}

func (c Query) QueryName() string {
    return QueryName
}
```

## Query Handler
###### src: app/user/application/query/Info/Handler.go
```golang
package Info

import (
    "context"
    "github.com/bordunosp/ddd/CQRS/QueryBus"
    "log"
)

func Handler(ctx context.Context, queryAny QueryBus.IQuery) (any, error) {
    request, _ := queryAny.(Query)
	
    log.Print(request.Id)

    return NewResponse(
        "name",
        "email",
    ), nil
}
```

## Query Response
###### src: app/user/application/query/Info/Response.go
```golang
package Info

func NewResponse(name, email string) *Response {
    return &Response{
        Name:  name,
        Email: email,
    }
}

type Response struct {
    Name  string
    Email string
}
```

---


## CommandBus Register
```golang
package main

import (
    "github.com/bordunosp/ddd/CQRS/CommandBus"
    "github.com/bordunosp/ddd/example/app/user/application/command/CreateNew"
    "github.com/bordunosp/ddd/example/app/user/application/command/UpdateEmail"
    "log"
)

func main() {
    err := CommandBus.RegisterCommands([]CommandBus.CommandItem{
        {CreateNew.CommandName, CreateNew.Handler},
        {UpdateEmail.CommandName, UpdateEmail.Handler},
    })
    if err != nil {
        log.Fatal(err)
    }
}
```

## CommandBus Usage
```golang
```

---

## QueryBus Register
```golang
package main

import (
    "github.com/bordunosp/ddd/CQRS/QueryBus"
    "github.com/bordunosp/ddd/example/app/user/application/query/Info"
    "log"
)

func main() {
    err := QueryBus.RegisterQueries([]QueryBus.QueryItem{
        {Info.QueryName, Info.Handler},
    })
    if err != nil {
        log.Fatal(err)
    }
}
```

## QueryBus Usage
```golang
```

---


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
```


