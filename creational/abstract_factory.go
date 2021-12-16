package creational

import (
	"fmt"
)

/*

Summary:
Abstract Factory provides an interface for creating generic product objects.
Removes dependencies of specifying concrete implementations from clients
that create product objects.

Example:
Motif and OpenLook are the implementations.
Abstract Factory is used to create the objects of these implementations.

Benefit:
This design of creating objects helps the client to create
objects without specifying concrete implementation. Enables decoupling.
*/

const (
	MotifType WindowKitType = iota
	OpenLookType
)

type WindowKitType int

// WindowKitFactory is an abstract factory interface for creating user interface
// toolkit, it prevents the clients from commiting to a implementation
// of the interface.
type WindowKitFactory interface {
	// CreateScrollBar does the work and creates a ScrollBar
	CreateScrollBar() ScrollBar
	// CreateWindow does the work and creates a Window
	CreateWindow() Window
}

// NewWindowKitFactory creates the factory for the specified
// WindowKit, using that WindowKit factory we can create the window kit.
func NewWindowKitFactory(t WindowKitType) (WindowKitFactory, error) {
	switch t {
	case MotifType:
		return NewMotifFactory(), nil
	case OpenLookType:
		return NewOpenLookFactory(), nil
	default:
		return nil, fmt.Errorf("unsupported window kit: %v", t)
	}
}

// NewMotifFactory creates the factory for creating Motif type of WindowKit
func NewMotifFactory() WindowKitFactory {
	return &motifFactory{}
}

// NewOpenLookFactory creates the factory for creating OpenLook type of WindowKit
func NewOpenLookFactory() WindowKitFactory {
	return &openLookFactory{}
}

// motifFactory is the factory for creating motif types window kit
type motifFactory struct{}

func (m motifFactory) CreateScrollBar() ScrollBar {
	return motifScrollBar{}
}

func (m motifFactory) CreateWindow() Window {
	return motifWindow{}
}

// openLookFactory is the factory for creating openLook types window kit
type openLookFactory struct{}

func (m openLookFactory) CreateScrollBar() ScrollBar {
	return openLookScrollBar{}
}

func (m openLookFactory) CreateWindow() Window {
	return openLookWindow{}
}

// ScrollBar is an interface to specify what a scrollbar could be
type ScrollBar interface {
	Scroll()
}

// Window is an interface to specify what a scrollbar could be
type Window interface {
	Open()
}

// Below are 2 scroll bar implementations:

type motifScrollBar struct{}

func (m motifScrollBar) Scroll() {
	fmt.Println("enjoy the motif scrolling")
}

type openLookScrollBar struct{}

func (m openLookScrollBar) Scroll() {
	fmt.Println("openlook scrolling")
}

// Below are 2 window implementations:

type motifWindow struct{}

func (m motifWindow) Open() {
	fmt.Println("opening motif window since 1990s")
}

type openLookWindow struct{}

func (m openLookWindow) Open() {
	fmt.Println("opening window since I do not remember")
}
