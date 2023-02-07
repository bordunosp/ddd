package UpdateEmail

import (
	"github.com/google/uuid"
)

type Command struct {
	Id    uuid.UUID
	Email string
}

func (c Command) CommandName() string {
	return "UpdateEmailCommand"
}
