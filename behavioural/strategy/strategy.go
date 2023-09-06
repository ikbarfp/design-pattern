package strategy

import (
	"fmt"
	"log"
	"time"
)

type Strategy interface {
	CalculateCost(distance float64) float64
	CalculateETA(distance float64) time.Duration
	CalculateDistance(origin, destination Location) float64
}

func (t *Trip) SetStrategy(s Strategy) {
	fmt.Println()
	log.Println("Set new strategy on runtime . . .")
	time.Sleep(1 * time.Second)
	fmt.Println()
	t.strategy = s
}
