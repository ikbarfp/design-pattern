package strategy

import "time"

const (
	defaultBusSpeed = 40
	// price/km with assumption at Rp7.000/liters and fuel consumption at 1:10
	defaultBusPricePerKM = 700
)

type Bus struct{}

func (b Bus) CalculateCost() float64 {

	return 0
}

func (b Bus) CalculateETA(distance float64) time.Duration {
	// NOTE: with assumption there will
	// be constant speed at 40 km/h
	duration := distance / defaultBusSpeed
	return time.Duration(duration)
}

func (b Bus) CalculateDistance(origin, destination Location) float64 {

	return 0
}
