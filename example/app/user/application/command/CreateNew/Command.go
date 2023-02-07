package CreateNew

import (
	"github.com/google/uuid"
)

type Command struct {
	Id    uuid.UUID
	Name  string
	Email string
}

func (c Command) CommandName() string {
	return "CreateNewCommand"
}
