


## Query
###### src: app/user/application/query/Info/Query.go
```golang
package Info

import (
    "github.com/bordunosp/ddd/CQRS/QueryBus"
    "github.com/google/uuid"
)

const QueryName = "InfoQuery"

func NewQuery(id uuid.UUID) QueryBus.IQuery {
    return &Query{
        Id: id,
    }
}

type Query struct {
    Id uuid.UUID
}

func (c Query) QueryName() string {
    return QueryName
}
```



## Query Handler
###### src: app/user/application/query/Info/Handler.go
```golang
package Info

import (
    "context"
    "github.com/bordunosp/ddd/CQRS/QueryBus"
    "log"
)

func Handler(ctx context.Context, queryAny QueryBus.IQuery) (any, error) {
    request, _ := queryAny.(Query)
	
    log.Print(request.Id)

    return NewResponse(
        "name",
        "email",
    ), nil
}
```

## Query Response
###### src: app/user/application/query/Info/Response.go
```golang
package Info

func NewResponse(name, email string) *Response {
    return &Response{
        Name:  name,
        Email: email,
    }
}

type Response struct {
    Name  string
    Email string
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
    err := QueryBus.RegisterQueries([]QueryBus.QueryItem{
        {Info.QueryName, Info.Handler},
    })
    if err != nil {
        log.Fatal(err)
    }
}
```

## QueryBus Usage
```golang
import (
    "github.com/google/uuid"
    "github.com/bordunosp/ddd/CQRS/QueryBus"
    "github.com/bordunosp/ddd/example/app/user/application/query/Info"
)

ctx := context.TODO()
query := Info.NewQuery(uuid.New())

// choose 1 of 3 possible ways to execute handler
dto, err = QueryBus.Handle(ctx, query)
dtoChan  = QueryBus.HandleAsync(ctx, query)
dto, err = QueryBus.HandleAsyncAwait(ctx, query)
```

