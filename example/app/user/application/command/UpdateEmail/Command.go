package UpdateEmail

import (
	"github.com/bordunosp/ddd/CQRS/CommandBus"
	"github.com/google/uuid"
)

type Command struct {
	Id    uuid.UUID
	Email string
}

func (c Command) CommandConfig() CommandBus.CommandConfig {
	return CommandBus.CommandConfig{
		Name:     "UpdateEmailCommand",
		Sanitize: true,
		Validate: true}
}
