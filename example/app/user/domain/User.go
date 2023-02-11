package domain

import (
	"github.com/bordunosp/ddd"
	"github.com/google/uuid"
)

func NewUser(id uuid.UUID, name, email string) User {
	return &user{
		AggregateRoot: ddd.NewAggregateRoot(id),
		name:          name,
		email:         email,
	}
}

type User interface {
	ddd.IAggregateRoot

	Name() string
	Email() string
}

type user struct {
	ddd.AggregateRoot

	name  string
	email string
}

func (u *user) IsEqual(other ddd.IAggregateRoot) bool {
	otherObj, ok := other.(User)
	return ok && u.ID() == otherObj.ID()
}

func (u *user) Name() string {
	return u.name
}

func (u *user) Email() string {
	return u.email
}
