
## Command (DTO)
###### src: app/user/application/command/UpdateEmail/Command.go
```golang
package UpdateEmail

import "github.com/google/uuid"

type Command struct {
    Id    uuid.UUID
    Email string `mod:"trim,lcase" validate:"required,email"`
}

// implemented interface CommandBus.ICommand 
// (need uniq name for registration)
func (c Command) CommandConfig() CommandBus.CommandConfig {
    return CommandBus.CommandConfig{
        Name:     "UpdateEmailCommand",
        Sanitize: true,
        Validate: true}
}

```

## Command Handler
###### src: app/user/application/command/UpdateEmail/Handler.go
```golang
package UpdateEmail

import (
    "context"
    "log"
)

func Handler(ctx context.Context, command Command) error {
    // todo: update email
    log.Print(command.Email)
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
    err := CommandBus.Register(UpdateEmail.Handler)
    if err != nil {
        log.Fatal(err)
    }
	
    err = CommandBus.Register(CreateNew.Handler)
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

func main() {
    ctx := context.TODO()
	
    command := CreateNew.Command{
        Id:    uuid.New(),
        Name:  "some commandDTO Name",
        Email: "some@commandDTO.name",
    }

    // choose 1 of 3 possible ways to execute handler
    _ = CommandBus.Execute(ctx, command)
    _ = CommandBus.ExecuteAsync(ctx, command)
    _ = CommandBus.ExecuteAsyncAwait(ctx, command)
}
```
