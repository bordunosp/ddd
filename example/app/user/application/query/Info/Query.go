package Info

import (
	"github.com/google/uuid"
)

type Query struct {
	Id uuid.UUID
}

func (c Query) QueryName() string {
	return "InfoQuery"
}
