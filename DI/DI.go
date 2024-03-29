package DI

import (
	"errors"
	"sync"
)

var ErrServiceNotRegistered = errors.New("service not registered")
var ErrServiceHasIncorrectType = errors.New("service has incorrect type")
var ErrServiceAlreadyRegistered = errors.New("service already registered")

var registeredServices = &sync.Map{}

func Get[T any]() (T, error) {
	var serviceT T
	err := ErrServiceNotRegistered

	registeredServices.Range(func(_, value any) bool {
		if serviceTT, ok := value.(T); ok {
			serviceT = serviceTT
			err = nil
			return true
		}
		return true
	})

	return serviceT, err
}

func GetOrPanic[T any]() T {
	service, err := Get[T]()
	if err != nil {
		panic(err)
	}

	return service
}

func GetByName[T any](serviceName string) (T, error) {
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

func GetByNameOrPanic[T any](serviceName string) T {
	service, err := GetByName[T](serviceName)
	if err != nil {
		panic(err)
	}

	return service
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

func Dispose() {
	registeredServices.Range(func(key, service any) bool {
		if disposableService, ok := service.(IDisposable); ok {
			disposableService.Dispose()
		}

		registeredServices.Delete(key)
		return true
	})
}
