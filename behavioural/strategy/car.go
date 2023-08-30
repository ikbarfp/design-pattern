package strategy

import "time"

const (
	defaultCarSpeed = 60
	// price/km with assumption at Rp15.000/liters and fuel consumption at 1:15
	defaultCarPricePerKm = 1000
)

type Car struct{}

func (c Car) CalculateCost() float64 {

	return 0
}

func (c Car) CalculateETA(distance float64) time.Duration {
	// NOTE: with assumption there will
	// be constant speed at 60 km/h
	duration := distance / defaultCarSpeed
	return time.Duration(duration)
}

func (c Car) CalculateDistance(origin, destination Location) float64 {

	return 0
}
