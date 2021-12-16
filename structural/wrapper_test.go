package structural

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapper(t *testing.T) {
	type test struct {
		pizza       Pizza
		toppings    []Topping
		expectPrice int
	}

	tests := []test{
		{
			pizza:       &Margaritta{},
			toppings:    []Topping{},
			expectPrice: 300,
		},
		{
			pizza:       &Margaritta{},
			toppings:    []Topping{Cheese},
			expectPrice: 330,
		},
		{
			pizza:       &Margaritta{},
			toppings:    []Topping{Cheese, Chicken},
			expectPrice: 370,
		},
	}

	for _, tc := range tests {
		tc = tc
		pizza := tc.pizza
		for _, topping := range tc.toppings {
			switch topping {
			case Cheese:
				pizza = &ToppingCheese{pizza: pizza} // wraps the earlier object with a topping
			case Chicken:
				pizza = &ToppingChicken{pizza: pizza} // wraps the earlier object with a topping
			}
		}
		assert.Equal(t, tc.expectPrice, pizza.Price())
	}
}
