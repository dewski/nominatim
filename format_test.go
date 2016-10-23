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
		assert.Equal(t, tt.expected, result.Short())
	}

	// {PlaceID:15401429 License: OsmType:node OsmID:1387190560 Class: Type:fuel Importance:0 Lat:38.8062335 Lon:-121.2049252 DisplayName:7 Eleven, Granite Drive, Rocklin, Placer County, California, 95650, United States of America Address:{Village: HouseNumber: Road:Granite Drive Residential: Town: City:Rocklin County:Placer County State:California PostCode:95650 Country:United States of America CountryCode:us Building: Supermarket: Fuel:7 Eleven BusStop:} NameDetails:{Name: HouseName:} ExtraTags:{Phone: Website: OpeningHours:} BoundingBox:[]}
	// Test goes here
}
