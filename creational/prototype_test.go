package creational

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrototype(t *testing.T) {
	toInt := func(i int) *int {
		return &i
	}

	tofloat32 := func(i float32) *float32 {
		return &i
	}

	type test struct {
		min      int
		max      int
		target   float32
		expected AutoScaler
	}

	tests := []test{
		{
			min:    1,
			max:    32,
			target: 0.8,
			expected: AutoScaler{
				Spec: &AutoScalerSpec{
					Min:    toInt(1),
					Max:    toInt(32),
					Target: tofloat32(0.8),
				},
				Status: &AutoScalerStatus{
					Current:   27,
					Desired:   27,
					Available: 27,
				},
			},
		},
	}

	for _, tc := range tests {
		tc = tc
		autoscaler := NewAutoScaler(tc.min, tc.max, tc.target)
		copy := autoscaler.UpdateStatus()
		assert.Equal(t, tc.expected, copy)
	}
}
