package DI

import "errors"

var ErrServiceNotRegistered = errors.New("service not registered")
var ErrServiceHasIncorrectType = errors.New("service has incorrect type")
var ErrServiceAlreadyRegistered = errors.New("service already registered")

var registeredServices = make(map[string]any)

func Get[T any](serviceName string, serviceT T) (T, error) {
	service, ok := registeredServices[serviceName]
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

func RegisterServices(services []ServiceItem[any]) error {
	for _, service := range services {
		if _, ok := registeredServices[service.ServiceName]; ok {
			return ErrServiceAlreadyRegistered
		}

		instance, err := service.ServiceInitFunc()
		if err != nil {
			return err
		}

		if service.IsSingleton {
			registeredServices[service.ServiceName] = instance
		} else {
			registeredServices[service.ServiceName] = service.ServiceInitFunc
		}
	}

	return nil
}
