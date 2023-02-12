
## Query (DTO)
###### src: app/user/application/query/Info/Query.go
```golang
package Info

import "github.com/google/uuid"

type Query struct {
    Id uuid.UUID
	Name string `mod:"trim" validate:"required"`
}

// implemented interface QueryBus.IQuery 
// (need uniq name for registration)
func (c Query) QueryConfig() QueryBus.QueryConfig {
    return QueryBus.QueryConfig{
        Name:             "InfoQuery",
        Sanitize:         true,
        Validate:         true,
        SanitizeResponse: true,
        ValidateResponse: false}
}

```


## Query Response (DTO)
###### src: app/user/application/query/Info/Response.go
```golang
package Info

type Response struct {
    Name  string `mod:"trim"`
    Email string `mod:"trim,lcase"`
}
```


## Query Handler
###### src: app/user/application/query/Info/Handler.go
```golang
package Info

import (
    "context"
    "log"
)

// type of QueryBus.IEventHandler 
func Handler(ctx context.Context, query Query) (Response, error) {
    log.Print(query.Id)

    return Response{
        Name: "res name",
        Email: "res@email.com",
    }, nil
}

```

## QueryBus Register
```golang
package main

import (
    "github.com/bordunosp/ddd/CQRS/QueryBus"
    "github.com/bordunosp/ddd/example/app/user/application/query/Info"
    "log"
)

func main() {
    err := QueryBus.Register(Info.Handler)
    if err != nil {
        log.Fatal(err)
    }
}
```

## QueryBus Usage
```golang
import (
    "github.com/bordunosp/ddd/CQRS/QueryBus"
    "github.com/bordunosp/ddd/example/app/user/application/query/Info"
    "github.com/google/uuid"
    "log"
)

func main() {
    ctx := context.TODO()

    query := Info.Query{
        Id: uuid.New(),
    }

    // choose 1 of 3 possible ways to execute handler
    res, err = QueryBus.Handle[Info.Response](ctx, query)
    resChan  = QueryBus.HandleAsync[Info.Response](ctx, query)
    res, err = QueryBus.HandleAsyncAwait[Info.Response](ctx, query)

    log.Println("Name from query response: ", res.Name)
}

```

