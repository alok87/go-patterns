# Design Patterns for Go

Design patterns in Go. Motive is to keep the examples as close to real world problems when possible. Suggestions and reviews are welcome.

## Categorization:
- Creational: Concerns with the process of object creation.
- Structural: Deals with composition of objects.
- Behavioural: Ways in which objects interact and distribute responsibility.

### Creational
- [Abstract Factory](./creational/abstract_factory.go)
- [Fluent Builder](./creational/fluent_builder.go)
- [Classic Builder](./creational/classic_builder.go)
- Factory Method
- Prototype
- Solitare


### Structural
- [Wrapper](./structural/wrapper.go)
- Adapter
- Bridge
- Decorator
- ObjectPool/Flyweight
- Proxy
- Glue
- Composite

### Behavioural
- [Strategy](./behavioural/strategy.go)
- Template Method
- Chain Of Responsibility
- Command
- Iterator
- Mediator
- Memento
- Observer
- State
- Interpretor
- Iterator
- Walker

## Contributing

Aim is to explain the benefit of the pattern with a simple real world example when possible. Feel free to create issues and pull requests to make it better and clearer for the reader.

## Referrences
Previous work on this subject which can also help in understanding:
- [Design Patterns by Gangs of Four(GOF)](https://en.wikipedia.org/wiki/Design_Patterns) is a book on this.
- [Go Patterns by bvwells](https://github.com/bvwells/go-patterns) have a lot of nice examples.
- [Refactoring Guru](https://refactoring.guru/design-patterns/go), also have a lot of examples.
- [Builder Pattern in Go](https://devcharmander.medium.com/design-patterns-in-golang-the-builder-dac468a71194) is a nice blog series on various styles on this pattern.
- [Fluent vs classic builder pattern](https://medium.com/@sawomirkowalski/design-patterns-builder-fluent-interface-and-classic-builder-d16ad3e98f6c)