package CommandBus

import (
	"errors"
	"fmt"
)

var registeredCommands = make(map[string]ICommandHandler)

func RegisterCommands(commandItems []CommandItem) error {
	for _, commandItem := range commandItems {
		if _, ok := registeredCommands[commandItem.CommandName]; ok {
			return errors.New(fmt.Sprintf("command by name '%s' already registered", commandItem.CommandName))
		}
		registeredCommands[commandItem.CommandName] = commandItem.Handler
	}
	return nil
}
