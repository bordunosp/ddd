## Classic way
```golang
package main

func checkError(err error, msg string) {
    if err != nil {
        panic(errors.Join(err, errors.New(msg)))
    }
}


func main()  {
    err := errors.New("same error")
	
    if err != nil {
        panic(errors.Join(err, errors.New("more information of error")))
    }
	
    // or
    checkError(err, "more information of error")
}
```


## Assertion way

```golang
package main

import (
    "github.com/bordunosp/ddd/Assertion"
    "errors"
)


func main() {
    err := errors.New("same error")
    Assertion.ErrorIsNull(err, "more information of error")
	
    // or make a new error if validation is failed
	
    someStr := "someStr"
    Assertion.OneOf(someStr, []string{"green", "yellow", "black"}, "someStr has unexpected value")
    Assertion.MaxLength(someStr, 10, "someStr must be less then 10")
    Assertion.MinLength(someStr, 10, "someStr must be greater then 10")
    Assertion.RangeLength(someStr, 10, 20, "someStr length mut be between 10 and 20")
    Assertion.Same(someStr, "some other string", "values not equals")
    Assertion.NotSame(someStr, someStr, "Values cant be equals")
    Assertion.Uuid(someStr, "someStr has incorrect uuid format")
    Assertion.Date(someStr, "02 Jan 06 15:04 MST", "someStr has incorrect date format")
    
    someBool := false
    Assertion.False(someBool, "someBool must be false")
    Assertion.True(someBool, "someBool must be true")
    Assertion.Same(someBool, someBool, "values not equals")
    Assertion.NotSame(someBool, someBool, "Values cant be equals")
    
    someInt := 40
    Assertion.OneOf(someInt, []int{10,50,100}, "someInt has unexpected value")
    Assertion.Max(someInt, 100, "someInt must be less then 100")
    Assertion.Min(someInt, 10, "someInt has value less then 10")
    Assertion.Range(someInt, 10, 20, "someInt not in range 10-20")
    Assertion.Same(someInt, someInt, "values not equals")
    Assertion.NotSame(someInt, someInt, "Values cant be equals")
    
    var service any
    Assertion.NotNull(service, "service cant be null")
}
```