package strategy

import (
	"fmt"
	"log"
	"time"
)

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
	log.Println("Starting a new trip, let's go! ")
	fmt.Println()
	return &Trip{
		strategy:    s,
		origin:      origin,
		destination: destination,
	}
}
