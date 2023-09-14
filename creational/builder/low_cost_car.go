package builder

import (
	"fmt"
	"log"
)

type LowCostCarBuilder struct {
	exhaust, steeringWheelMaterial, tyreSpec, windowFilm, engine string
	totalSeat, totalTyre                                         int
	isExistSteeringWheel                                         bool
	transmission                                                 Transmission
}

func (lc *LowCostCarBuilder) AddExhaust() *LowCostCarBuilder {

	if len(lc.exhaust) == 0 {
		lc.exhaust = "Standard Exhaust"
		return lc
	}

	return lc
}

func (lc *LowCostCarBuilder) AddSeat(numOfSeat int) *LowCostCarBuilder {
	maxCap := 7
	diff := maxCap - lc.totalSeat + numOfSeat

	if diff < 0 {
		log.Fatalln("can't add more new seat in low cost car")
		return lc
	}

	lc.totalSeat += numOfSeat
	return lc
}

func (lc *LowCostCarBuilder) AddSteeringWheel() *LowCostCarBuilder {

	if lc.isExistSteeringWheel {
		log.Fatalln("can't add more new steering wheel in low cost car")
		return lc
	}

	lc.isExistSteeringWheel = true
	lc.steeringWheelMaterial = "Polyurethane Steering Wheels"
	return lc
}

func (lc *LowCostCarBuilder) AddTyre(numOfTyre int) *LowCostCarBuilder {
	maxCap := 4
	diff := maxCap - lc.totalTyre + numOfTyre

	if diff < 0 {
		log.Fatalln("can't add more new tyre in low cost car")
		return lc
	}

	lc.totalTyre += numOfTyre
	lc.tyreSpec = "Dunlop Enasave 170/65 R14"
	return lc
}

func (lc *LowCostCarBuilder) AddWindowFilm() *LowCostCarBuilder {
	lc.windowFilm = "Standard Window Film"
	return lc
}

func (lc *LowCostCarBuilder) BuildEngine() *LowCostCarBuilder {
	lc.engine = "1500cc Naturally Aspirated"
	return lc
}

func (lc *LowCostCarBuilder) BuildTransmission(transmission Transmission) *LowCostCarBuilder {

	switch transmission {
	case CVT:
		lc.transmission = CVT
	case Manual:
		lc.transmission = Manual
	default:
		log.Fatalln("low cost car doesn't support this transmission")
	}

	return lc
}

func (lc *LowCostCarBuilder) BuildCar() (*Car, error) {

	if len(lc.engine) == 0 {
		return nil, fmt.Errorf("there is no engine installed")
	}

	if len(lc.transmission) == 0 {
		return nil, fmt.Errorf("there is no transmission installed")
	}

	if len(lc.exhaust) == 0 {
		return nil, fmt.Errorf("there is no exhaust installed")
	}

	if lc.totalTyre != 4 {
		return nil, fmt.Errorf("insufficient/excess number of tyre")
	}

	if !lc.isExistSteeringWheel {
		return nil, fmt.Errorf("there is no steering wheel installed")
	}
	return &Car{
		exhaust:       lc.exhaust,
		steeringWheel: lc.steeringWheelMaterial,
		tyreSpec:      lc.tyreSpec,
		tyre:          lc.totalTyre,
		windowFilm:    lc.windowFilm,
		engine:        lc.engine,
		transmission:  lc.transmission,
		seat:          lc.totalSeat,
	}, nil
}
