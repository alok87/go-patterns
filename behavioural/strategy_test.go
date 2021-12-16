package behavioural

import (
	"fmt"
	"testing"

	"github.com/alok87/go-capturer"
	"github.com/stretchr/testify/assert"
)

func TestStrategy(t *testing.T) {
	type test struct {
		strategy    Strategy
		result      string
		errExpected bool
	}

	tests := []test{
		{
			strategy: RoundRobin,
			result:   "roundRobin\nloadbalanced\n",
		},
		{
			strategy: WeightedRoundRobin,
			result:   "weightedRoundRobin\nloadbalanced\n",
		},
		{
			strategy: LeastConnection,
			result:   "leastConn\nloadbalanced\n",
		},
		{
			strategy:    Strategy(0),
			result:      "roundRobin\nloadbalanced\n",
			errExpected: false,
		},
		{
			strategy:    Strategy(30),
			result:      "",
			errExpected: true,
		},
	}

	// default scenario
	lb := NewLoadbalancer()
	assert.NotNil(t, lb)
	defaultOut := capturer.CaptureOutput(func() {
		lb.Balance()
	})
	assert.Equal(t, "roundRobin\nloadbalanced\n", defaultOut)

	for _, tc := range tests {
		tc = tc

		fmt.Printf("%+v\n", tc)
		err := lb.SetStrategy(tc.strategy)
		if tc.errExpected && err == nil {
			t.Errorf("Error:%v expected, got nil", err)
			continue
		}
		if !tc.errExpected && err != nil {
			t.Errorf("Error not expected, got err:%v", err)
			continue
		}

		if tc.errExpected {
			continue
		}

		assert.Nil(t, err)
		assert.NotNil(t, lb)

		out := capturer.CaptureOutput(func() {
			lb.Balance()
		})
		assert.Equal(t, tc.result, out)
	}
}
