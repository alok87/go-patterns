package creational

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFluentBuilder(t *testing.T) {
	type test struct {
		name     string
		company  string
		address  string
		expected Employee
	}

	tests := []test{
		{
			name:     "emp1",
			company:  "comp1",
			address:  "earth",
			expected: Employee{"emp1", "comp1", "earth"},
		},
	}

	for _, tc := range tests {
		tc = tc
		builder := NewEmployeeBuilder()
		got := builder.
			Called(tc.name).
			WorksAt(tc.company).
			At(tc.address).
			Build()
		assert.Equal(t, tc.expected, got)
	}
}
