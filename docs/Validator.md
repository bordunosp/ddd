## Validation 
### for validation we use most used [Validator](https://github.com/go-playground/validator)
###### for validate custom types u can use Middleware.ValidateRegisterCustomTypeFunc:

```golang
package main

import (
    "database/sql"
    "github.com/bordunosp/ddd/CQRS/Middleware"
    "github.com/go-playground/validator/v10"
    "reflect"
)

func main() {
    Middleware.ValidateRegisterCustomTypeFunc(ValidateValuer, sql.NullString{}, sql.NullInt64{})
}

// ValidateValuer implements validator.CustomTypeFunc
func ValidateValuer(field reflect.Value) interface{} {
    if valuer, ok := field.Interface().(driver.Valuer); ok {
        val, err := valuer.Value()
        if err == nil {
            return val
        }
        // handle the error how you want
    }

    return nil
}
```

