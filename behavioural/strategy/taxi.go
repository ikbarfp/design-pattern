package strategy

import "time"

const (
	defaultTaxiSpeed = 60
	// price/km with assumption at Rp10.000/liters and fuel consumption at 1:16
	defaultTaxiPricePerKM = 625
)

type Taxi struct{}

func (t Taxi) CalculateCost() float64 {

	return 0
}

func (t Taxi) CalculateETA(distance float64) time.Duration {
	// NOTE: with assumption there will
	// be constant speed at 60 km/h
	duration := distance / defaultTaxiSpeed
	return time.Duration(duration)
}

func (t Taxi) CalculateDistance(origin, destination Location) float64 {

	return 0
}
