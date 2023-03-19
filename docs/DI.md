## DI
```golang
package main

import (
    "github.com/bordunosp/ddd/DI"
    "log"
    "os"
)

func init()  {
    err := DI.RegisterServices([]DI.ServiceItem{
        {
            // will be initialized immediately (once)
            IsSingleton: true,
            ServiceName: "logger",
            ServiceInitFunc: func() (any, error) {
                // initialize simple log as Example
                return log.New(os.Stderr, "\t", log.Ldate|log.Ltime|log.Lshortfile), nil
            },
        },
        {
            // will be initialized many times (per each call)
            IsSingleton: false,
            ServiceName: "UserService",
            ServiceInitFunc: func() (any, error) {
                logger, err := DI.Get[*log.Logger]()
                if err != nil {
                    return nil, err
                }

                // creating new UserService which use logger from DI
                return infrastructure.NewUserService(logger)
            },
        },
    })

    if err != nil {
        log.Fatal(err)
    }
}

func main() {
    // if service (registered as singleton) implemented ddd.IDisposable interface
    // it will be called for dispose
    // can be used for close connections etc... 
    defer DI.Dispose()

    // Use service from DI
    // it can be used anywhere in your project (after registered)
    logger, _ := DI.GetByName[*log.Logger]("logger")
    logger.Println("logger.Println called")
}
```
