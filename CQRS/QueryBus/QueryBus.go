package QueryBus

import (
	"errors"
	"fmt"
)

var registeredQueries = make(map[string]IQueryHandler)

func RegisterQueries(queryItems []QueryItem) error {
	for _, queryItem := range queryItems {
		if _, ok := registeredQueries[queryItem.QueryName]; ok {
			return errors.New(fmt.Sprintf("query by name '%s' already registered", queryItem.QueryName))
		}
		registeredQueries[queryItem.QueryName] = queryItem.Handler
	}
	return nil
}
