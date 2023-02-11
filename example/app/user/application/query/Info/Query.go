package Info

import (
	"github.com/google/uuid"
)

type Query struct {
	Id   uuid.UUID
	Name string `mod:"trim" validate:"required"`
}

func (c Query) QueryName() string {
	return "InfoQuery"
}
