package main

import (
	"context"
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
	"github.com/google/uuid"
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
	Assertion.ErrorIsNull(CommandBus.Register(CreateNew.Handler), "CreateNew.Handler register")
	Assertion.ErrorIsNull(CommandBus.Register(UpdateEmail.Handler), "UpdateEmail.Handler register")
}

func registerQueries() {
	Assertion.ErrorIsNull(QueryBus.Register(Info.Handler), "Info.Handler register")
}

func registerEvents() {
	err := EventBus.Register([]EventBus.IEventHandler[event.UserCreated]{
		Created.SaveLogHandler,
		Created.SendEmailHandler,
	})
	Assertion.ErrorIsNull(err, "event.UserCreated register")

	err = EventBus.Register([]EventBus.IEventHandler[event.UserUpdated]{
		// event may not have handlers
		//
		// you never know when it might be really useful
		// that is why events are created long before handlers are created.
	})
	Assertion.ErrorIsNull(err, "event.UserUpdated register")
}

func init() {
	registerDI()
	registerCommands()
	registerQueries()
	registerEvents()
}

func main() {
	ctx := context.TODO()

	// Use service from DI
	// it can be used anywhere in your project (after registered)
	logger, _ := DI.Get[*log.Logger]("logger")
	logger.Println("logger.Println called")

	// QueryBus Handle
	res, err := QueryBus.Handle[Info.Response](ctx, Info.Query{
		Id: uuid.New(),
	})
	Assertion.ErrorIsNull(err, "Info.Query handle")
	logger.Println("Name from query response: ", res.Name)

	// CommandBus Handle
	err = CommandBus.Execute(ctx, CreateNew.Command{
		Id:    uuid.New(),
		Name:  "some commandDTO Name",
		Email: "some@commandDTO.name",
	})
	Assertion.ErrorIsNull(err, "CreateNew.Command handle")

	// EventBus Handle
	err = EventBus.Dispatch(ctx, event.UserCreated{
		Id:    uuid.New(),
		Name:  "some eventDTO name",
		Email: "some@event.email",
	})
	Assertion.ErrorIsNull(err, "event.UserCreated handle")
}
