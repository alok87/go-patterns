package structural

/*

Summary:
Add new functionality to objects dynamically by wrapping
the core objects inside a wrapper

Example:
Interface: Pizza
And many implementations but of 2 categories of implementation
1st: we create pizza type like peperroni and margaritta which gives the base pizza price "return 299"
2nd: then we add toppings on its base using wrapped structs on it which also implements Pizza interface "t.pizza.Price() + 20"
so that they can again be wrapped.

Benefit:
Using wrapper we can wrap objects countless number of times as
wrapper objects follow the same interface. Resulting objects gets a stacking
behaviour of all objects

Ref: https://refactoring.guru/design-patterns/decorator/go/example
*/

type Topping int

const (
	Cheese Topping = iota
	Chicken
)

// Pizza defines what makes a Pizza, this is the core
type Pizza interface {
	Price() int
}

// Margaritta wraps the Pizza interface, it is an actual pizza implementation
type Margaritta struct {
	pizza Pizza
}

func (t *Margaritta) Price() int {
	return 300
}

// Pepperoni wraps the Pizza interface, it is an actual pizza implementation
type Pepperoni struct {
	pizza Pizza
}

func (t *Pepperoni) Price() int {
	return 400
}

// ToppingCheese wraps the Pizza interface, extends pizza functionality
type ToppingCheese struct {
	pizza Pizza
}

func (t *ToppingCheese) Price() int {
	return t.pizza.Price() + 30
}

// ToppingChicken wraps the Pizza interface, extends pizza functionality
type ToppingChicken struct {
	pizza Pizza
}

func (t *ToppingChicken) Price() int {
	return t.pizza.Price() + 40
}
