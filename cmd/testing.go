package main

import (
	"fmt"

	"github.com/dewski/nominatim"
)

func main() {
	options := nominatim.Options{
		Lat:            38.806597,
		Lon:            -121.205412,
		Zoom:           18,
		AddressDetails: true,
		ExtraTags:      true,
	}

	resp, err := nominatim.ReverseGeocode(options)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", resp)
}
