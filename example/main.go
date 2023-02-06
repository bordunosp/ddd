package main

import (
	"github.com/bordunosp/ddd/Assertion"
	"github.com/bordunosp/ddd/CQRS/CommandBus"
	"github.com/bordunosp/ddd/CQRS/EventBus"
	"github.com/bordunosp/ddd/CQRS/QueryBus"
	"github.com/bordunosp/ddd/DI"
	"github.com/bordunosp/ddd/example/app/user/application/command/CreateNew"
	"github.com/bordunosp/ddd/example/app/user/application/command/UpdateEmail"
	"github.com/bordunosp/ddd/example/app/user/application/event/user/Created"
	"github.com/bordunosp/ddd/example/app/user/application/query/Info"
	"github.com/bordunosp/ddd/example/app/user/domain/event"
	"github.com/bordunosp/ddd/example/app/user/infrastructure"
	"log"
	"os"
)

func registerDI() {
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
	Assertion.ErrorIsNull(err, "Cant register DI services")
}

func registerCommands() {
	err := CommandBus.RegisterCommand[CreateNew.Command](CommandBus.CommandItem[CreateNew.Command]{
		CommandName: CreateNew.CommandName,
		Handler:     CreateNew.Handler[CreateNew.Command],
	})
	Assertion.ErrorIsNull(err, "Cant register command "+CreateNew.CommandName)

	err = CommandBus.RegisterCommand[UpdateEmail.Command](CommandBus.CommandItem[UpdateEmail.Command]{
		CommandName: UpdateEmail.CommandName,
		Handler:     UpdateEmail.Handler[UpdateEmail.Command],
	})
	Assertion.ErrorIsNull(err, "Cant register command "+UpdateEmail.CommandName)
}

func registerQueries() {
	err := QueryBus.RegisterQuery[Info.Query, Info.Response](QueryBus.QueryItem[Info.Query, Info.Response]{
		QueryName: Info.QueryName,
		Handler:   Info.Handler[Info.Query, Info.Response],
	})
	Assertion.ErrorIsNull(err, "Cant register query "+Info.QueryName)
}

func registerEvents() {
	err := EventBus.RegisterEvent[EventBus.IEvent](EventBus.EventItem[EventBus.IEvent]{
		EventName: event.UserUpdatedEvent,
		Handlers:  []EventBus.IEventHandler[EventBus.IEvent]{
			// event may not have handlers
			//
			// you never know when it might be really useful
			// that is why events are created long before handlers are created.
		},
	})
	Assertion.ErrorIsNull(err, "Cant register Event "+event.UserUpdatedEvent)

	err = EventBus.RegisterEvent[event.UserCreated](EventBus.EventItem[event.UserCreated]{
		EventName: event.UserCreatedEvent,
		Handlers: []EventBus.IEventHandler[event.UserCreated]{
			Created.SaveLogHandler,
			Created.SendEmailHandler,
		},
	})
	Assertion.ErrorIsNull(err, "Cant register Event "+event.UserCreatedEvent)
}

func init() {
	registerDI()
	registerCommands()
	registerQueries()
	registerEvents()
}

func main() {
	// Use service from DI
	// it can be used anywhere in your project (after registered)
	logger, _ := DI.Get[*log.Logger]("logger")
	logger.Println("logger.Println called")
}
