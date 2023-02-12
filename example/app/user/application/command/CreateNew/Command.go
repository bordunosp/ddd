package CreateNew

import (
	"github.com/bordunosp/ddd/CQRS/CommandBus"
	"github.com/google/uuid"
)

type Command struct {
	Id    uuid.UUID
	Name  string
	Email string
}

func (c Command) CommandConfig() CommandBus.CommandConfig {
	return CommandBus.CommandConfig{
		Name:     "CreateNewCommand",
		Sanitize: true,
		Validate: true}
}
