
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
