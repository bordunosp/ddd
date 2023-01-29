package event

import (
	"github.com/bordunosp/ddd/CQRS/EventBus"
	"github.com/google/uuid"
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
