package event

import (
	"github.com/bordunosp/ddd/CQRS/EventBus"
	"github.com/google/uuid"
)

const UserUpdatedEvent = "UserUpdatedEvent"

func NewUserUpdated(id uuid.UUID, name, email string) EventBus.IEvent {
	return &UserUpdated{
		Id:    id,
		Name:  name,
		Email: email,
	}
}

type UserUpdated struct {
	Id    uuid.UUID
	Name  string
	Email string
}

func (u UserUpdated) EventName() string {
	return UserUpdatedEvent
}
