package nominatim

import "fmt"

var (
	longAttractionFormat   = "%s at %s %s, %s, %s %s"
	longResidentialFormat  = "%s %s, %s, %s %s"
	shortAttractionFormat  = "%s at %s %s, %s"
	shortResidentialFormat = "%s %s, %s"
)

func (a Address) Long() (short string) {
	switch {
	case a.Attraction != "":
		short = fmt.Sprintf(
			longAttractionFormat,
			a.Attraction,
			a.HouseNumber,
			a.Road,
			a.City,
			a.State,
			a.PostCode,
		)
	case a.Fuel != "":
		short = fmt.Sprintf(
			longAttractionFormat,
			a.Fuel,
			a.HouseNumber,
			a.Road,
			a.City,
			a.State,
			a.PostCode,
		)
	case a.BusStop != "":
		short = fmt.Sprintf(
			longAttractionFormat,
			a.BusStop,
			a.HouseNumber,
			a.Road,
			a.City,
			a.State,
			a.PostCode,
		)
	case a.Supermarket != "":
		short = fmt.Sprintf(
			longAttractionFormat,
			a.Supermarket,
			a.HouseNumber,
			a.Road,
			a.City,
			a.State,
			a.PostCode,
		)
	case a.Building != "":
		short = fmt.Sprintf(
			longAttractionFormat,
			a.Building,
			a.HouseNumber,
			a.Road,
			a.City,
			a.State,
			a.PostCode,
		)
	default:
		short = fmt.Sprintf(
			longResidentialFormat,
			a.HouseNumber,
			a.Road,
			a.City,
			a.State,
			a.PostCode,
		)
	}

	return
}

func (a Address) Short() (short string) {
	switch {
	case a.Attraction != "":
		short = fmt.Sprintf(
			shortAttractionFormat,
			a.Attraction,
			a.HouseNumber,
			a.Road,
			a.City,
		)
	case a.Fuel != "":
		short = fmt.Sprintf(
			shortAttractionFormat,
			a.Fuel,
			a.HouseNumber,
			a.Road,
			a.City,
		)
	case a.BusStop != "":
		short = fmt.Sprintf(
			shortAttractionFormat,
			a.BusStop,
			a.HouseNumber,
			a.Road,
			a.City,
		)
	case a.Supermarket != "":
		short = fmt.Sprintf(
			shortAttractionFormat,
			a.Supermarket,
			a.HouseNumber,
			a.Road,
			a.City,
		)
	case a.Building != "":
		short = fmt.Sprintf(
			shortAttractionFormat,
			a.Building,
			a.HouseNumber,
			a.Road,
			a.City,
		)
	default:
		short = fmt.Sprintf(
			shortResidentialFormat,
			a.HouseNumber,
			a.Road,
			a.City,
		)
	}

	return
}
