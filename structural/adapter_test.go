package structural

import (
	"testing"

	capture "github.com/alok87/go-capture"
	"github.com/stretchr/testify/assert"
)

func TestAdapterExternalSource(t *testing.T) {
	t.Parallel()

	indigo := NewIndigo()
	assert.NotNil(t, indigo)
	out := capture.Output(func() {
		indigo.Get("6E 447")
	})
	assert.Equal(t, "indigo: flight: 6E 447, status: boarding\n", out)
}

func TestAdapterCache(t *testing.T) {
	t.Parallel()

	indigo := NewIndigo()
	assert.NotNil(t, indigo)

	cache := NewCacheAdapter(indigo)
	assert.NotNil(t, cache)
	out := capture.Output(func() {
		cache.Get("6E 447")
	})
	assert.Equal(t, "indigo: flight: 6E 447, status: boarding\n", out)

	out = capture.Output(func() {
		cache.Get("6E 447")
	})
	assert.Equal(t, "cache: flight: 6E 447, status: boarding\n", out)

	out = capture.Output(func() {
		cache.Get("6E 227")
	})
	assert.Equal(t, "indigo: flight: 6E 227, status: boarding\n", out)
}
