package structural

import "fmt"

/*

Summary:
An adapter makes communication between incompatible objects possible.

macTypeC port -> typeC 2 HDMI adapter -> TV HDMI
hindi -> translator -> hindi

Example:
We would use a real world example of caching a external data
source response using an Adapter in between to limit API calls to ext service.

Client -> Adapter -> Adaptee
Client -> CacheAdapter -> Indigo (ext data source)
Client -> FlightData -> FlightData (as interfaces if u see)

Both CacheAdapter and Indigo have the same interface for the client

Benefit:
Client is able to communicate to cache in between as if it was
speaking to an external service. Made possible using adapter.

So the benefit is it does not require a change in the main data interface
to introduce some functionality in between. If both can cater to the same
interface.
*/

// FlightData is the interface that is used to get the flight data
type FlightData interface {
	Get(flightNumber string)
}

// Indigo is the external flight datasource
// it implements FlightData interface
type Indigo struct {
}

func NewIndigo() FlightData {
	return &Indigo{}
}

func (a *Indigo) Get(flightNumber string) {
	fmt.Printf(
		"indigo: flight: %v, status: boarding\n", flightNumber,
	)
}

// cacheAdapter also implements FlightData
// it is the wrapper in between that acts like an adapter to the client
type cacheAdapter struct {
	flightData FlightData

	// cache caches the Indigo data to limit the no of api calls
	// to this ext data service
	cache map[string]bool
}

func NewCacheAdapter(d FlightData) FlightData {
	return &cacheAdapter{
		flightData: d,
		cache:      make(map[string]bool),
	}
}

func (a *cacheAdapter) Get(flightNumber string) {
	got, ok := a.cache[flightNumber]
	if ok && got {
		fmt.Printf(
			"cache: flight: %v, status: boarding\n", flightNumber,
		)
		return
	}

	// call to the main data service
	a.flightData.Get(flightNumber)

	a.cache[flightNumber] = true
}
