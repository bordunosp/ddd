

### DDD (Common part)  

---

```shell
go get github.com/bordunosp/ddd@v0.0.34
```

---

- [AggregateRoot](https://github.com/bordunosp/ddd/blob/main/docs/AggregateRoot.md) - is the parent "Entity" to all other Entities and Value Objects within the Aggregate. A Repository operates upon an Aggregate Root.
- [Entity](https://github.com/bordunosp/ddd/blob/main/docs/Entity.md) - is something that has an identifier and an owner
- [DI](https://github.com/bordunosp/ddd/blob/main/docs/DI.md) - (Dependency Injection) is a design pattern in which an object or function receives other objects or functions that it depends on
- [Command](https://github.com/bordunosp/ddd/blob/main/docs/CommandBus.md) - is a behavioral design pattern that turns a request into a stand-alone object that contains all information about the request.
- [Query](https://github.com/bordunosp/ddd/blob/main/docs/QueryBus.md) - is a behavioral design pattern like Command but can return any value (an object structure that can be interpreted into a SQL query)
- [Event](https://github.com/bordunosp/ddd/blob/main/docs/EventBus.md) - is fundamental design pattern used to create a communication channel and communicate through it via events
- [Assertion](https://github.com/bordunosp/ddd/blob/main/docs/Assertion.md) - is a list of functions for more readable error checking
- [Sanitizer](https://github.com/bordunosp/ddd/blob/main/docs/Sanitizer.md) - way of sanitize (command, query, event) DTO
- [Validator](https://github.com/bordunosp/ddd/blob/main/docs/Validator.md) - way of validate (command, query, event) DTO
- [Saga](https://github.com/bordunosp/ddd/blob/main/docs/Saga.md) - The Saga architecture pattern provides transaction management using a sequence of local transactions
