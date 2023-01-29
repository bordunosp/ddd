package CreateNew

import (
	"github.com/bordunosp/ddd/CQRS/CommandBus"
	"github.com/google/uuid"
)

const CommandName = "CreateNewCommand"

func NewCommand(id uuid.UUID, name string, email string) CommandBus.ICommand {
	return &Command{
		Id:    id,
		Name:  name,
		Email: email,
	}
}

type Command struct {
	Id    uuid.UUID
	Name  string
	Email string
}

func (c Command) CommandName() string {
	return CommandName
}
