package nominatim

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Result struct {
	PlaceID     string      `json:"place_id,omitempty"`
	License     string      `json:"license,omitempty"`
	OsmType     string      `json:"osm_type,omitempty"`
	OsmID       string      `json:"osm_id,omitempty"`
	Lat         string      `json:"lat,omitempty"`
	Lon         string      `json:"lon,omitempty"`
	DisplayName string      `json:"display_name,omitempty"`
	Address     Address     `json:"address,omitempty"`
	NameDetails NameDetails `json:"namedetails,omitempty"`
	ExtraTags   ExtraTags   `json:"extratags,omitempty"`
	BoundingBox []string    `json:"boundingbox,omitempty"`
}

type Address struct {
	Village     string `json:"village,omitempty,omitempty"`
	Town        string `json:"town,omitempty"`
	Supermarket string `json:"supermarket,omitempty"`
	Building    string `json:"building,omitempty"`
	HouseNumber string `json:"house_number,omitempty"`
	Road        string `json:"road,omitempty"`
	Residential string `json:"residential,omitempty"`
	City        string `json:"city,omitempty"`
	County      string `json:"county,omitempty"`
	State       string `json:"state,omitempty"`
	Postcode    string `json:"postcode,omitempty"`
	Country     string `json:"country,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
}

type NameDetails struct {
	Name      string `json:"name,omitempty"`
	HouseName string `json:"addr:housename,omitempty"`
}

type ExtraTags struct {
	Phone        string `json:"phone,omitempty"`
	Website      string `json:"website,omitempty"`
	OpeningHours string `json:"opening_hours,omitempty"`
}

type Options struct {
	Lat            float64
	Lon            float64
	Zoom           int
	AddressDetails bool
	ExtraTags      bool
}

var (
	email     = ""
	urlFormat = "https://nominatim.openstreetmap.org/reverse?format=json&lat=%f&lon=%f&zoom=%d&addressdetails=%d&extratags=%d&email=%s"
)

func SetEmail(e string) {
	email = e
}

func ReverseGeocode(options Options) (result Result, err error) {
	details := 0
	if options.AddressDetails {
		details = 1
	}

	extraTags := 0
	if options.ExtraTags {
		extraTags = 1
	}

	url := fmt.Sprintf(urlFormat, options.Lat, options.Lon, options.Zoom, details, extraTags, email)
	resp, err := http.Get(url)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return result, err
	}

	return result, nil
}
