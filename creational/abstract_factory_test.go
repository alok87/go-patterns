package creational

import (
	"testing"

	capture "github.com/alok87/go-capture"
	"github.com/stretchr/testify/assert"
)

func TestWindowKitFactory(t *testing.T) {
	t.Parallel()
	var factory WindowKitFactory
	var err error

	factory, err = NewWindowKitFactory(MotifType)
	assert.Nil(t, err)
	assert.NotNil(t, factory)

	factory, err = NewWindowKitFactory(OpenLookType)
	assert.Nil(t, err)
	assert.NotNil(t, factory)

	factory, err = NewWindowKitFactory(WindowKitType(2))
	assert.NotNil(t, err)
}

func TestMotifScrollBar(t *testing.T) {
	t.Parallel()
	factory, err := NewWindowKitFactory(MotifType)
	assert.Nil(t, err)
	assert.NotNil(t, factory)

	scrollbar := factory.CreateScrollBar()
	_, ok := scrollbar.(motifScrollBar)
	assert.True(t, ok)

	out := capture.Stdout(func() {
		scrollbar.Scroll()
	})
	assert.Equal(t, "enjoy the motif scrolling\n", out)
}

func TestOpenLookScrollBar(t *testing.T) {
	t.Parallel()
	factory, err := NewWindowKitFactory(OpenLookType)
	assert.Nil(t, err)
	assert.NotNil(t, factory)

	scrollbar := factory.CreateScrollBar()
	_, ok := scrollbar.(openLookScrollBar)
	assert.True(t, ok)

	out := capture.Stdout(func() {
		scrollbar.Scroll()
	})
	assert.Equal(t, "openlook scrolling\n", out)
}

func TestMotifWindow(t *testing.T) {
	t.Parallel()
	factory, err := NewWindowKitFactory(MotifType)
	assert.Nil(t, err)
	assert.NotNil(t, factory)

	window := factory.CreateWindow()
	_, ok := window.(motifWindow)
	assert.True(t, ok)

	out := capture.Output(func() {
		window.Open()
	})
	assert.Equal(t, "opening motif window since 1990s\n", out)
}

func TestOpenLookWindow(t *testing.T) {
	t.Parallel()
	factory, err := NewWindowKitFactory(OpenLookType)
	assert.Nil(t, err)
	assert.NotNil(t, factory)

	window := factory.CreateWindow()
	_, ok := window.(openLookWindow)
	assert.True(t, ok)

	out := capture.Output(func() {
		window.Open()
	})
	assert.Equal(t, "opening window since I do not remember\n", out)
}
