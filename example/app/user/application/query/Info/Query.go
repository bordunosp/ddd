package Info

import (
	"github.com/bordunosp/ddd/CQRS/QueryBus"
	"github.com/google/uuid"
)

type Query struct {
	Id   uuid.UUID
	Name string `mod:"trim" validate:"required"`
}

func (c Query) QueryConfig() QueryBus.QueryConfig {
	return QueryBus.QueryConfig{
		Name:             "InfoQuery",
		Sanitize:         true,
		Validate:         true,
		SanitizeResponse: false,
		ValidateResponse: false}
}
