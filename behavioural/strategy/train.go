package strategy

import "time"

const (
	defaultTrainSpeed = 100
	// price/km with assumption at Rp2.000/kWH and watt consumption at 1:10
	defaultTrainPricePerKM = 250
)

type Train struct{}

func (t Train) CalculateCost() float64 {

	return 0
}

func (t Train) CalculateETA(distance float64) time.Duration {
	// NOTE: with assumption there will
	// be constant speed at 100 km/h
	duration := distance / defaultTrainSpeed
	return time.Duration(duration)
}

func (t Train) CalculateDistance(origin, destination Location) float64 {

	return 0
}
