package nominatim

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type Result struct {
	PlaceID     string      `json:"place_id,omitempty"`
	License     string      `json:"license,omitempty"`
	OsmType     string      `json:"osm_type,omitempty"`
	OsmID       string      `json:"osm_id,omitempty"`
	Class       string      `json:"string,omitempty"`
	Type        string      `json:"type,omitempty"`
	Importance  string      `json:"importance,omitempty"`
	Lat         string      `json:"lat,omitempty"`
	Lon         string      `json:"lon,omitempty"`
	DisplayName string      `json:"display_name,omitempty"`
	Address     Address     `json:"address,omitempty"`
	NameDetails NameDetails `json:"namedetails,omitempty"`
	ExtraTags   ExtraTags   `json:"extratags,omitempty"`
	BoundingBox []string    `json:"boundingbox,omitempty"`
}

type Address struct {
	Village       string `json:"village,omitempty"`
	HouseNumber   string `json:"house_number,omitempty"`
	Road          string `json:"road,omitempty"`
	Residential   string `json:"residential,omitempty"`
	Town          string `json:"town,omitempty"`
	City          string `json:"city,omitempty"`
	County        string `json:"county,omitempty"`
	State         string `json:"state,omitempty"`
	PostCode      string `json:"postcode,omitempty"`
	Country       string `json:"country,omitempty"`
	CountryCode   string `json:"country_code,omitempty"`
	StateDistrict string `json:"state_district,omitempty"`

	Pedestrian  string `json:"pedestrian,omitempty"`
	Attraction  string `json:"attraction,omitempty"`
	Building    string `json:"building,omitempty"`
	Supermarket string `json:"supermarket,omitempty"`
	Fuel        string `json:"fuel,omitempty"`
	BusStop     string `json:"bus_stop,omitempty"`
}

type NameDetails struct {
	Name      string `json:"name,omitempty"`
	HouseName string `json:"addr:housename,omitempty"`
}

type ExtraTags struct {
	// See http://wiki.openstreetmap.org/wiki/Annotations
	Attribution string `json:"attribution,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Description string `json:"description,omitempty"`
	Email       string `json:"email,omitempty"`
	Fax         string `json:"fax,omitempty"`
	Image       string `json:"image,omitempty"`
	Note        string `json:"note,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Source      string `json:"source,omitempty"`
	SourceName  string `json:"source:name,omitempty"`
	SourceRef   string `json:"source:ref,omitempty"`
	Todo        string `json:"todo,omitempty"`
	Website     string `json:"website,omitempty"`
	Wikipedia   string `json:"wikipedia,omitempty"`

	// See http://wiki.openstreetmap.org/wiki/Map_Features#Properties
	OpeningHours string `json:"opening_hours,omitempty"`
	Fee          string `json:"fee,omitempty"`
}

type Options struct {
	Lat            float64
	Lon            float64
	Zoom           int
	AddressDetails bool
	ExtraTags      bool
	OsmIds         []string
}

var (
	ErrMissingLookupOsmIds  = errors.New("Missing OSM ids")
	email                   = ""
	reverseGeocodeUrlFormat = "https://nominatim.openstreetmap.org/reverse?format=json&lat=%f&lon=%f&zoom=%d&addressdetails=%d&extratags=%d&email=%s"
	lookupUrlFormat         = "https://nominatim.openstreetmap.org/lookup?osm_ids=%s&format=json&addressdetails=%d&extratags=%d&email=%s"
)

func SetEmail(e string) {
	email = e
}

func ReverseGeocode(options Options) (result Result, err error) {
	_, details, extraTags := options.parse()

	url := fmt.Sprintf(
		reverseGeocodeUrlFormat,
		options.Lat,
		options.Lon,
		options.Zoom,
		details,
		extraTags,
		email,
	)

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

func (r Result) OsmParam() (nodeID string) {
	switch r.OsmType {
	case "node":
		nodeID = fmt.Sprintf("N%s", r.OsmID)
	case "relation":
		nodeID = fmt.Sprintf("R%s", r.OsmID)
	case "way":
		nodeID = fmt.Sprintf("W%s", r.OsmID)
	default:
		panic("Unknown OSM Type")
	}

	return
}

func Lookup(options Options) (results []Result, err error) {
	osmList, details, extraTags := options.parse()

	if osmList == "" {
		err = ErrMissingLookupOsmIds
		return
	}

	url := fmt.Sprintf(
		lookupUrlFormat,
		osmList,
		details,
		extraTags,
		email,
	)

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return
	}

	return
}

func (o Options) parse() (osmList string, details, extraTags int) {
	details, extraTags = 0, 0

	if o.AddressDetails {
		details = 1
	}

	if o.ExtraTags {
		extraTags = 1
	}

	if len(o.OsmIds) > 0 {
		osmList = strings.Join(o.OsmIds, ",")
	}

	return
}
