package DI

import (
	"errors"
	"github.com/bordunosp/ddd"
	"sync"
)

var ErrServiceNotRegistered = errors.New("service not registered")
var ErrServiceHasIncorrectType = errors.New("service has incorrect type")
var ErrServiceAlreadyRegistered = errors.New("service already registered")

var registeredServices = struct {
	mu    sync.RWMutex
	items map[string]any
}{
	mu:    sync.RWMutex{},
	items: make(map[string]any),
}

func Get[T any](serviceName string) (T, error) {
	registeredServices.mu.RLocker()
	defer registeredServices.mu.RUnlock()

	var serviceT T

	service, ok := registeredServices.items[serviceName]
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
	registeredServices.mu.Lock()
	defer registeredServices.mu.Unlock()

	for _, service := range services {
		if _, ok := registeredServices.items[service.ServiceName]; ok {
			return ErrServiceAlreadyRegistered
		}

		if service.IsSingleton {
			instance, err := service.ServiceInitFunc()
			if err != nil {
				return err
			}
			registeredServices.items[service.ServiceName] = instance
		} else {
			registeredServices.items[service.ServiceName] = service.ServiceInitFunc
		}
	}

	return nil
}

func Dispose() {
	registeredServices.mu.Lock()
	defer registeredServices.mu.Unlock()

	for _, service := range registeredServices.items {
		if _, ok := service.(ServiceInitFunc); ok {
			continue
		}

		if disposableService, ok := service.(ddd.IDisposable); ok {
			disposableService.Dispose()
		}
	}

	registeredServices.items = make(map[string]any)
}
