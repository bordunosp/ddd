


## AggregateRoot
```golang
// src: app/user/domain/user.go

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
```golang
// src: app/user/domain/entity/profile.go

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

## Event
```golang
// src: app/user/domain/event/created.go

import (
    "github.com/google/uuid"
    "github.com/bordunosp/ddd"
)

func NewCreatedEvent(userID uuid.UUID) ddd.IEvent {
    return &createdDomainEvent{
        userID: userID,
    }
}

type createdEvent struct {
    userID uuid.UUID
}

func (event createdEvent) Name() string {
    return "user_created"
}

func (event createdEvent) AggregateID() uuid.UUID {
    return event.userID
}
```

## Event vs AggregateRoot
```golang
user, _ := NewUser(uuid.New(), "name", "email@test.com")
event := NewCreatedEvent(user.UUID())
_ := user.AddDomainEvent(event)

// DomainEvents must be executed before the transaction is committed

// 1) open transaction
// 2) Save user to DB
// 3) Execute DomainEvents user.DomainEvents()
// 4) commit transaction 
```