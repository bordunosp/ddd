package DI

import (
	"errors"
	"fmt"
)

var registeredServices = make(map[string]any)

func Get(serviceName string) (any, error) {
	service, ok := registeredServices[serviceName]
	if !ok {
		return nil, errors.New(fmt.Sprintf("service by name '%s' not registered", serviceName))
	}

	if diInitFunc, ok := service.(ServiceInitFunc); ok {
		return diInitFunc()
	}

	return service, nil
}

func RegisterServices(services []ServiceItem) error {
	for _, service := range services {
		if _, ok := registeredServices[service.ServiceName]; ok {
			return errors.New(fmt.Sprintf("service by name '%s' already registered", service.ServiceName))
		}

		if service.IsSingleton {
			instance, err := service.ServiceInitFunc()
			if err != nil {
				return err
			}
			registeredServices[service.ServiceName] = instance
		} else {
			registeredServices[service.ServiceName] = service.ServiceInitFunc
		}
	}

	return nil
}
