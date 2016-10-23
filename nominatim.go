package nominatim

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Result struct {
	PlaceID     string      `json:"place_id"`
	License     string      `json:"license"`
	OsmType     string      `json:"osm_type"`
	OsmID       string      `json:"osm_id"`
	Lat         string      `json:"lat"`
	Lon         string      `json:"lon"`
	DisplayName string      `json:"display_name"`
	Address     Address     `json:"address"`
	NameDetails NameDetails `json:"namedetails"`
	ExtraTags   ExtraTags   `json:"extratags"`
	BoundingBox []string    `json:"boundingbox"`
}

type Address struct {
	Village     string `json:"village"`
	Town        string `json:"town"`
	Supermarket string `json:"supermarket"`
	Building    string `json:"building"`
	HouseNumber string `json:"house_number"`
	Road        string `json:"road"`
	Residential string `json:"residential"`
	City        string `json:"city"`
	County      string `json:"county"`
	State       string `json:"state"`
	Postcode    string `json:"postcode"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
}

type NameDetails struct {
	Name      string `json:"name"`
	HouseName string `json:"addr:housename"`
}

type ExtraTags struct {
	Phone        string `json:"phone"`
	Website      string `json:"website"`
	OpeningHours string `json:"opening_hours"`
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
