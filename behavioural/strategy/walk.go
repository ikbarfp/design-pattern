package strategy

import "time"

const (
	defaultWalkSpeed = 10
)

type Walk struct{}

func (w Walk) CalculateCost() float64 {
	// NOTE: with assumption there will
	// be no cost at all with walk trip
	return 0
}

func (w Walk) CalculateETA(distance float64) time.Duration {
	// NOTE: with assumption there will
	// be constant speed at 10 km/h
	duration := distance / defaultWalkSpeed
	return time.Duration(duration)
}

func (w Walk) CalculateDistance(origin, destination Location) float64 {

	return 0
}
