# Design Patterns for Go

Design patterns for Go. 

Patters have `Summary`, `Example` and `Benefit` section present as comments. The focus is to bring out the **benefit** of the pattern with a simple real world example.

## Categorization:
- Creational: Concerns with the process of object creation.
- Structural: Deals with composition of objects.
- Behavioural: Ways in which objects interact and distribute responsibility.

### Creational
- [Abstract Factory](./creational/abstract_factory.go)
- [Fluent Builder](./creational/fluent_builder.go)
- [Classic Builder](./creational/classic_builder.go)
- [Factory Method](./creational/factory_method.go)
- [Prototype](./creational/prototype.go)

### Structural
- [Wrapper](./structural/wrapper.go)
- [Adapter](./structural/adapter.go)
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

Feel free to create issues and pull requests to make it better and clearer for the reader.
```
make test
```

## Referrences
Previous work on this subject which helped in making this:
- [Design Patterns by Gangs of Four(GOF)](https://en.wikipedia.org/wiki/Design_Patterns) is a book on this.
- [Go Patterns by bvwells](https://github.com/bvwells/go-patterns) have a lot of good simple examples.
- [Refactoring Guru](https://refactoring.guru/design-patterns/go) and [GolanyByExample](https://golangbyexample.com/all-design-patterns-golang/) also have a lot of examples.
- [Builder Pattern Series by Surya Reddy](https://devcharmander.medium.com/design-patterns-in-golang-the-builder-dac468a71194) demostrate various styles on this pattern.
- [Fluent and Classic builder pattern](https://medium.com/@sawomirkowalski/design-patterns-builder-fluent-interface-and-classic-builder-d16ad3e98f6c) separates the builder pattern into two types, internet is filled with confusing articles on it. There are actually 2 different builder patterns.
- [Prototype Pattern by Surya Reddy](https://devcharmander.medium.com/design-pattern-in-golang-prototype-e864522e4eeb)
- [DudeWhoCodes uses Adapater Pattern](https://dudewho.codes/adapter-pattern-in-go/)
