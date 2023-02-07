package event

import (
	"github.com/google/uuid"
)

type UserCreated struct {
	Id    uuid.UUID
	Name  string
	Email string
}

func (u UserCreated) EventName() string {
	return "UserCreatedEvent"
}
