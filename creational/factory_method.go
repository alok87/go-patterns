package creational

import "fmt"

/*

Summary:
Factory pattern is used to create product's objects without
specifying concrete types. NewProduct() method is used to create the
concrete product instead of &Product{}

Example:
Parking lot service. This service has a persistent storage
as a dependency and Factory pattern helps in doing dependency injection.

Benefit:
1. Dependency Injection/Unit testable: All dependencies are passed as interfaces
in the factory method as parameters, making the code unit-testable/mockable.
2. Reusability: All clients can just call the Factory method to make the object.
3. Extensibility: All client codes remain untouched, when we add some
new things in object making. Only one method changes.
*/

type ParkingType int

const (
	Street ParkingType = iota
	Mall
)

// Repository is the external dependency of Parking
type Repository interface {
	Create()
	Update()
}

type Parking interface {
	ParkVehicle()
	UnParkVehicle()
	DisplayFreeSlots()
}

// NewParking is the factory method that makes store be passed
// as an external dependency which can be mocked. It is the single
// function that would be used by all clients for creating all
// kinds of parking: onstreet, mall, etc..
func NewParking(parkType ParkingType, store Repository) Parking {
	switch parkType {
	case Street:
		return newOnStreetParking(store)
	case Mall:
		return newMallParking(store)
	}

	return newMallParking(store)
}

type streetParking struct {
	store Repository
}

func newOnStreetParking(store Repository) *streetParking {
	return &streetParking{store}
}
func (p *streetParking) ParkVehicle()      { fmt.Println("street parked") }
func (p *streetParking) UnParkVehicle()    { fmt.Println("street unparked") }
func (p *streetParking) DisplayFreeSlots() { fmt.Println("street free slots") }

type mallParking struct {
	store Repository
}

func newMallParking(store Repository) *mallParking {
	return &mallParking{store}
}
func (p *mallParking) ParkVehicle()      { fmt.Println("mall parked") }
func (p *mallParking) UnParkVehicle()    { fmt.Println("mall unparked") }
func (p *mallParking) DisplayFreeSlots() { fmt.Println("mall free slots") }
