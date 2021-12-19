package behavioural

import (
	"testing"

	"github.com/alok87/go-capture"
	"github.com/stretchr/testify/assert"
)

func TestChainOfResponsibility(t *testing.T) {
	e := NewEscalation("vicky", []string{"down"}, nil)
	e = NewEscalation("ricky", []string{"spillover"}, e)
	e = NewEscalation("micky", []string{"dns"}, e)

	out := capture.Output(func() {
		e.Handle("dns")
	})
	assert.Equal(t, "resolved by micky", out)

	out = capture.Output(func() {
		e.Handle("spillover")
	})
	assert.Equal(t, "resolved by ricky", out)

	out = capture.Output(func() {
		e.Handle("down")
	})
	assert.Equal(t, "resolved by vicky", out)

	out = capture.Output(func() {
		e.Handle("unknown")
	})
	assert.Equal(t, "timedout", out)

}
