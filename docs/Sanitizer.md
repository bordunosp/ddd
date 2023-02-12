## Sanitizer
### for sanitize we use most used [Sanitizer](https://github.com/go-playground/mold)
###### for use custom sanitize item u can use Middleware.SanitizeRegisterTag:

```golang
package main

import (
    "context"
    "github.com/bordunosp/ddd/CQRS/Middleware"
    "github.com/go-playground/mold/v4"
)

func main() {
	Middleware.SanitizeRegisterTag("custom_trim", customTrim)
}

// customTrim implements mold.Func
func customTrim(ctx context.Context, fl mold.FieldLevel) error {
    // do stuff 
    return nil
}
```

