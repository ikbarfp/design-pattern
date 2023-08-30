package strategy

import "time"

type Trip struct {
	strategy    Strategy
	origin      Location
	destination Location
	cost        float64
	eta         time.Duration
	distance    float64
}

type Location struct {
	ProvinceName string
	CityName     string
	DistrictName string
	VillageName  string
	PostalCode   string
	Lat          float64
	Lng          float64
}

func NewTrip(s Strategy, origin, destination Location) *Trip {
	return &Trip{
		strategy:    s,
		origin:      origin,
		destination: destination,
		// cost:        s.CalculateCost(),
		// eta:         s.CalculateETA(),
		// distance:    s.CalculateDistance(),
	}
}
