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

func (t *Trip) GetDistance(s Strategy) {
	t.distance = s.CalculateDistance(t.origin, t.destination)
}

func (t *Trip) GetCost(s Strategy) {
	t.cost = s.CalculateCost()
}

func (t *Trip) GetETA(s Strategy) {
	t.eta = s.CalculateETA(t.distance)
}


