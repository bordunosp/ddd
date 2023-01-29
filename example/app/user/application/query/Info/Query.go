package Info

import (
	"github.com/bordunosp/ddd/CQRS/QueryBus"
	"github.com/google/uuid"
)

const QueryName = "InfoQuery"

func NewQuery(id uuid.UUID) QueryBus.IQuery {
	return &Query{
		Id: id,
	}
}

type Query struct {
	Id uuid.UUID
}

func (c Query) QueryName() string {
	return QueryName
}
