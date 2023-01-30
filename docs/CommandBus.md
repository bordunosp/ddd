
## Command
###### src: app/user/application/command/UpdateEmail/Command.go
```golang
package UpdateEmail

import (
    "github.com/bordunosp/ddd/CQRS/CommandBus"
    "github.com/google/uuid"
)

const CommandName = "UpdateEmailCommand"

func NewCommand(id uuid.UUID, email string) CommandBus.ICommand {
    return &Command{
        Id:    id,
        Email: email,
    }
}

type Command struct {
    Id    uuid.UUID
    Email string
}

func (c Command) CommandName() string {
    return CommandName
}
```

## Command Handler
###### src: app/user/application/command/UpdateEmail/Handler.go
```golang
package UpdateEmail

import (
    "context"
    "github.com/bordunosp/ddd/DI"
    "github.com/bordunosp/ddd/CQRS/CommandBus"
)

func Handler(ctx context.Context, commandAny CommandBus.ICommand) error {
    request, _ := commandAny.(Command)

    loggerAny, _ := DI.Get("logger")
    loggerAny.(*log.Logger).Println(request.Email)
	
    return nil
}
```


## CommandBus Register
```golang
package main

import (
    "github.com/bordunosp/ddd/CQRS/CommandBus"
    "github.com/bordunosp/ddd/example/app/user/application/command/CreateNew"
    "github.com/bordunosp/ddd/example/app/user/application/command/UpdateEmail"
    "log"
)

func main() {
    err := CommandBus.RegisterCommands([]CommandBus.CommandItem{
        {CreateNew.CommandName, CreateNew.Handler},
        {UpdateEmail.CommandName, UpdateEmail.Handler},
    })
    if err != nil {
        log.Fatal(err)
    }
}
```

## CommandBus Usage
```golang
import (
    "github.com/google/uuid"
    "github.com/bordunosp/ddd/CQRS/CommandBus"
    "github.com/bordunosp/ddd/example/app/user/application/command/UpdateEmail"
)

ctx := context.TODO()
command := UpdateEmail.NewCommand(uuid.New(), "some@new.email")

// choose 1 of 3 possible ways to execute handler
_ = CommandBus.Execute(ctx, command)
_ = CommandBus.ExecuteAsync(ctx, command)
_ = CommandBus.ExecuteAsyncAwait(ctx, command)
```
