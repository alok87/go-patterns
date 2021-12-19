package structural

import (
	"testing"

	"github.com/alok87/go-capture"
	"github.com/stretchr/testify/assert"
)

func TestDecorateV1(t *testing.T) {
	hello := func(s string) string {
		return s
	}
	d := DecorateV1(hello)
	out := capture.Output(func() {
		d("raj")
	})
	assert.Equal(t, out, "pre post:raj")
}

func TestDecorateV2(t *testing.T) {
	hello := func(s string) string {
		return s
	}
	d := DecorateV2(hello)
	out := capture.Output(func() {
		d("vic")
	})
	assert.Equal(t, out, "pre post:vic")
}

func TestDecorateV3(t *testing.T) {
	hello := func(s string) string {
		return s
	}
	out := capture.Output(func() {
		DecorateV3(hello, "vic_raj8")
	})
	assert.Equal(t, out, "pre post:vic_raj8")
}
