package nominatim

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResultAddressShortFormat(t *testing.T) {
	testTable := []struct {
		houseNumber string
		attraction  string
		building    string
		supermarket string
		fuel        string
		busStop     string
		expected    string
	}{
		{"1600", "White House", "", "", "", "", "White House at 1600 Pennsylvania Avenue Northwest, Washington, District of Columbia 20500"},
		{"1600", "", "Building", "", "", "", "Building at 1600 Pennsylvania Avenue Northwest, Washington, District of Columbia 20500"},
		{"1600", "", "", "Raley's", "", "", "Raley's at 1600 Pennsylvania Avenue Northwest, Washington, District of Columbia 20500"},
		{"1600", "", "", "", "Chevron", "", "Chevron at 1600 Pennsylvania Avenue Northwest, Washington, District of Columbia 20500"},
		{"1600", "", "", "", "", "Bus Stop", "Bus Stop at 1600 Pennsylvania Avenue Northwest, Washington, District of Columbia 20500"},
	}

	for _, tt := range testTable {
		result := Address{
			Attraction:  tt.attraction,
			Building:    tt.building,
			Supermarket: tt.supermarket,
			Fuel:        tt.fuel,
			BusStop:     tt.busStop,
			HouseNumber: tt.houseNumber,
			Road:        "Pennsylvania Avenue Northwest",
			City:        "Washington",
			State:       "District of Columbia",
			PostCode:    "20500",
			Country:     "United States of America",
		}

		assert.Equal(t, tt.expected, result.Long())
	}
}

func TestResultAddressLongFormat(t *testing.T) {
	testTable := []struct {
		houseNumber string
		attraction  string
		building    string
		supermarket string
		fuel        string
		busStop     string
		expected    string
	}{
		{"1600", "White House", "", "", "", "", "White House at 1600 Pennsylvania Avenue Northwest, Washington"},
		{"1600", "", "Building", "", "", "", "Building at 1600 Pennsylvania Avenue Northwest, Washington"},
		{"1600", "", "", "Raley's", "", "", "Raley's at 1600 Pennsylvania Avenue Northwest, Washington"},
		{"1600", "", "", "", "Chevron", "", "Chevron at 1600 Pennsylvania Avenue Northwest, Washington"},
		{"1600", "", "", "", "", "Bus Stop", "Bus Stop at 1600 Pennsylvania Avenue Northwest, Washington"},
	}

	for _, tt := range testTable {
		result := Address{
			Attraction:  tt.attraction,
			Building:    tt.building,
			Supermarket: tt.supermarket,
			Fuel:        tt.fuel,
			BusStop:     tt.busStop,
			HouseNumber: tt.houseNumber,
			Road:        "Pennsylvania Avenue Northwest",
			City:        "Washington",
			State:       "District of Columbia",
			PostCode:    "20500",
			Country:     "United States of America",
		}

		assert.Equal(t, tt.expected, result.Short())
	}
}
