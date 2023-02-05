package DI

import (
	"errors"
	"sync"
)

var ErrServiceNotRegistered = errors.New("service not registered")
var ErrServiceHasIncorrectType = errors.New("service has incorrect type")
var ErrServiceAlreadyRegistered = errors.New("service already registered")

var registeredServices = sync.Map{}

func Get[T any](serviceName string) (T, error) {
	var serviceT T

	service, ok := registeredServices.Load(serviceName)
	if !ok {
		return serviceT, ErrServiceNotRegistered
	}

	if diInitFunc, ok := service.(ServiceInitFunc); ok {
		obj, err := diInitFunc()
		if err != nil {
			return serviceT, err
		}

		convertableType, ok := obj.(T)
		if !ok {
			return serviceT, ErrServiceHasIncorrectType
		}

		return convertableType, nil
	}

	convertableType, ok := service.(T)
	if !ok {
		return serviceT, ErrServiceHasIncorrectType
	}

	return convertableType, nil
}

func RegisterServices(services []ServiceItem) error {
	for _, service := range services {
		if _, ok := registeredServices.Load(service.ServiceName); ok {
			return ErrServiceAlreadyRegistered
		}

		if service.IsSingleton {
			instance, err := service.ServiceInitFunc()
			if err != nil {
				return err
			}
			registeredServices.Store(service.ServiceName, instance)
		} else {
			registeredServices.Store(service.ServiceName, service.ServiceInitFunc)
		}
	}

	return nil
}
