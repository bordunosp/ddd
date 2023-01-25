


## AggregateRoot
```golang
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