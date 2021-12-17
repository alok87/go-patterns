package creational

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClassicBuilder(t *testing.T) {
	type test struct {
		maker    RobotMaker
		expected Robot
	}

	tests := []test{
		{
			maker: Tesla,
			expected: Robot{
				head:  true,
				torso: true,
				legs:  false,
				arms:  true,
			},
		},
		{
			maker: Apple,
			expected: Robot{
				head:  true,
				torso: true,
				legs:  true,
				arms:  true,
			},
		},
		{
			maker: 0,
			expected: Robot{
				head:  true,
				torso: true,
				legs:  true,
				arms:  true,
			},
		},
	}

	for _, tc := range tests {
		tc = tc
		d := NewDirector(tc.maker)
		got := d.Build()
		assert.Equal(t, tc.expected, got)
	}

	// test set
	d := NewDirector(Apple)
	robot := d.Build()
	assert.Equal(t, robot, Robot{true, true, true, true})
	d.SetRobotMaker(Tesla)
	robot = d.Build()
	assert.Equal(t, robot, Robot{true, true, false, true})
}
