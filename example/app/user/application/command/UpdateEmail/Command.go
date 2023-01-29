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
