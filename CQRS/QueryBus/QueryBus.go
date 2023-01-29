package QueryBus

import (
	"errors"
)

var ErrQueryAlreadyRegistered = errors.New("query already registered")

var registeredQueries = make(map[string]IQueryHandler)

func RegisterQueries(queryItems []QueryItem) error {
	for _, queryItem := range queryItems {
		if _, ok := registeredQueries[queryItem.QueryName]; ok {
			return ErrQueryAlreadyRegistered
		}
		registeredQueries[queryItem.QueryName] = queryItem.Handler
	}
	return nil
}
