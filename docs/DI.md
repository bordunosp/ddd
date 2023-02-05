## DI
```golang
package main

import (
    "github.com/bordunosp/ddd/DI"
    "log"
    "os"
)

func main() {
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
                logger, err := DI.Get[*log.Logger]("logger")
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

    // Use service from DI
    // it can be used anywhere in your project (after registered)
    logger, _ := DI.Get[*log.Logger]("logger")
    logger.Println("logger.Println called")
}
```
