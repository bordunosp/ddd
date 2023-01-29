package CommandBus

import (
	"errors"
)

var ErrCommandAlreadyRegistered = errors.New("command already registered")

var registeredCommands = make(map[string]ICommandHandler)

func RegisterCommands(commandItems []CommandItem) error {
	for _, commandItem := range commandItems {
		if _, ok := registeredCommands[commandItem.CommandName]; ok {
			return ErrCommandAlreadyRegistered
		}
		registeredCommands[commandItem.CommandName] = commandItem.Handler
	}
	return nil
}
