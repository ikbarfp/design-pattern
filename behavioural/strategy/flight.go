package strategy

import "time"

const (
	defaultFlightSpeed = 500
	// price/km with assumption at RpXXX/liters and fuel consumption at 1:XX
	defaultFlightPricePerKM = 2000
)

type Flight struct{}

func (f Flight) CalculateCost(distance float64) float64 {
	return distance * defaultFlightPricePerKM
}

func (f Flight) CalculateETA(distance float64) time.Duration {
	// NOTE: with assumption there will
	// be constant speed at 500 km/h
	duration := distance / defaultFlightSpeed
	return time.Duration(duration * float64(time.Hour))
}

func (f Flight) CalculateDistance(origin, destination Location) float64 {

	return 0
}
