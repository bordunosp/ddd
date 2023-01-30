

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

