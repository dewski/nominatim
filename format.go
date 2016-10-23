package nominatim

import "fmt"

var (
	attractionFormat  = "%s at %s %s, %s, %s %s"
	residentialFormat = "%s %s, %s, %s %s"
)

func (a Address) Short() (short string) {
	switch {
	case a.Attraction != "":
		short = fmt.Sprintf(
			attractionFormat,
			a.Attraction,
			a.HouseNumber,
			a.Road,
			a.City,
			a.State,
			a.PostCode,
		)
	case a.Fuel != "":
		short = fmt.Sprintf(
			attractionFormat,
			a.Fuel,
			a.HouseNumber,
			a.Road,
			a.City,
			a.State,
			a.PostCode,
		)
	case a.BusStop != "":
		short = fmt.Sprintf(
			attractionFormat,
			a.BusStop,
			a.HouseNumber,
			a.Road,
			a.City,
			a.State,
			a.PostCode,
		)
	case a.Supermarket != "":
		short = fmt.Sprintf(
			attractionFormat,
			a.Supermarket,
			a.HouseNumber,
			a.Road,
			a.City,
			a.State,
			a.PostCode,
		)
	case a.Building != "":
		short = fmt.Sprintf(
			attractionFormat,
			a.Building,
			a.HouseNumber,
			a.Road,
			a.City,
			a.State,
			a.PostCode,
		)
	default:
		short = fmt.Sprintf(
			residentialFormat,
			a.HouseNumber,
			a.Road,
			a.City,
			a.State,
			a.PostCode,
		)
	}

	return
}
