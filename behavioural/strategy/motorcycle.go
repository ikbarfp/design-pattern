package strategy

import "time"

const (
	defaultMotorcycleSpeed = 50
	// price/km with assumption at Rp15.000/liters and fuel consumption at 1:25
	defaultMotorcyclePricePerKM = 600
)

type Motorcycle struct{}

func (m Motorcycle) CalculateCost(distance float64) float64 {
	return distance * defaultMotorcyclePricePerKM
}

func (m Motorcycle) CalculateETA(distance float64) time.Duration {
	// NOTE: with assumption there will
	// be constant speed at 50 km/h
	duration := distance / defaultMotorcycleSpeed
	return time.Duration(duration * float64(time.Hour))
}

func (m Motorcycle) CalculateDistance(origin, destination Location) float64 {

	return 15
}
