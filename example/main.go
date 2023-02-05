package main

import (
	"github.com/bordunosp/ddd/CQRS/CommandBus"
	"github.com/bordunosp/ddd/CQRS/EventBus"
	"github.com/bordunosp/ddd/CQRS/QueryBus"
	"github.com/bordunosp/ddd/DI"
	"github.com/bordunosp/ddd/example/app/user/application/command/CreateNew"
	"github.com/bordunosp/ddd/example/app/user/application/command/UpdateEmail"
	"github.com/bordunosp/ddd/example/app/user/application/event/user/created"
	"github.com/bordunosp/ddd/example/app/user/application/query/Info"
	"github.com/bordunosp/ddd/example/app/user/domain/event"
	"github.com/bordunosp/ddd/example/app/user/infrastructure"
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

	err = CommandBus.RegisterCommands([]CommandBus.CommandItem{
		{CreateNew.CommandName, CreateNew.Handler},
		{UpdateEmail.CommandName, UpdateEmail.Handler},
	})
	if err != nil {
		log.Fatal(err)
	}

	err = QueryBus.RegisterQueries([]QueryBus.QueryItem{
		{Info.QueryName, Info.Handler},
	})
	if err != nil {
		log.Fatal(err)
	}

	err = EventBus.RegisterEvents([]EventBus.EventItem{
		{
			EventName: event.UserCreatedEvent,
			Handlers: []EventBus.IEventHandler{
				Created.SaveLogHandler,
				Created.SendEmailHandler,
			},
		},

		{event.UserUpdatedEvent, []EventBus.IEventHandler{
			// event may not have handlers
			//
			// you never know when it might be really useful
			// that is why events are created long before handlers are created.
		}},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Use service from DI
	// it can be used anywhere in your project (after registered)
	logger, _ := DI.Get[*log.Logger]("logger")
	logger.Println("logger.Println called")
}
