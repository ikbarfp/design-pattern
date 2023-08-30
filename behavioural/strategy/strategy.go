package strategy

import "time"

type Strategy interface {
	CalculateCost() float64
	CalculateETA(distance float64) time.Duration
	CalculateDistance(origin, destination Location) float64
}

func (t *Trip) SetStrategy(s Strategy) {
	t.strategy = s
}
