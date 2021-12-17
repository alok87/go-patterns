package creational

import (
	"testing"

	capturer "github.com/alok87/go-capturer"
	"github.com/stretchr/testify/assert"
)

func TestFactoryMethod(t *testing.T) {
	type test struct {
		parkType ParkingType
		store    Repository
		expected string
	}

	tests := []test{
		{
			parkType: Street,
			store:    newmockRepository(func() {}, func() {}),
			expected: "street parked\nstreet unparked\nstreet free slots\n",
		},
		{
			parkType: Mall,
			store:    newmockRepository(func() {}, func() {}),
			expected: "mall parked\nmall unparked\nmall free slots\n",
		},
	}

	for _, tc := range tests {
		tc = tc
		parking := NewParking(tc.parkType, tc.store)
		assert.NotNil(t, parking)
		out := capturer.CaptureOutput(func() {
			parking.ParkVehicle()
			parking.UnParkVehicle()
			parking.DisplayFreeSlots()
		})
		assert.Equal(t, tc.expected, out)
	}
}

// mockRepository mocks and implements Repository
type mockRepository struct {
	CreateFunc func()
	UpdateFunc func()
}

func (m *mockRepository) Create() {
	m.CreateFunc()
}

func (m *mockRepository) Update() {
	m.UpdateFunc()
}

func newmockRepository(create func(), update func()) *mockRepository {
	return &mockRepository{create, update}
}
