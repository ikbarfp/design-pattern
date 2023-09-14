package builder

import (
	"fmt"
	"log"
)

type SportsCarBuilder struct {
	exhaust, steeringWheelMaterial, tyreSpec, windowFilm, engine string
	totalSeat, totalTyre                                         int
	isExistSteeringWheel                                         bool
	transmission                                                 Transmission
}

func (s *SportsCarBuilder) AddExhaust() *SportsCarBuilder {

	if len(s.exhaust) == 0 {
		s.exhaust = "Racing Exhaust"
		return s
	}

	return s
}

func (s *SportsCarBuilder) AddSeat(numOfSeat int) *SportsCarBuilder {
	maxCap := 2
	diff := maxCap - s.totalSeat + numOfSeat

	if diff < 0 {
		log.Fatalln("can't add more new seat in sports car")
		return s
	}

	s.totalSeat += numOfSeat
	return s
}

func (s *SportsCarBuilder) AddSteeringWheel() *SportsCarBuilder {

	if s.isExistSteeringWheel {
		log.Fatalln("can't add more new steering wheel in sports car")
		return s
	}

	s.isExistSteeringWheel = true
	s.steeringWheelMaterial = "Premium Nappa Leather Steering Wheels"
	return s
}

func (s *SportsCarBuilder) AddTyre(numOfTyre int) *SportsCarBuilder {
	maxCap := 4
	diff := maxCap - s.totalTyre + numOfTyre

	if diff < 0 {
		log.Fatalln("can't add more new tyre in sports car")
		return s
	}

	s.totalTyre += numOfTyre
	s.tyreSpec = "Michelin Primacy 4 185/45 R20"
	return s
}

func (s *SportsCarBuilder) AddWindowFilm() *SportsCarBuilder {
	s.windowFilm = "V-Kool Premium Window Film"
	return s
}

func (s *SportsCarBuilder) BuildEngine() *SportsCarBuilder {
	s.engine = "3500cc V8 Biturbo"
	return s
}

func (s *SportsCarBuilder) BuildTransmission(transmission Transmission) *SportsCarBuilder {

	switch transmission {
	case DualClutch:
		s.transmission = DualClutch
	default:
		log.Fatalln("sports car doesn't support this transmission")
	}

	return s
}

func (s *SportsCarBuilder) BuildCar() (*Car, error) {

	if len(s.engine) == 0 {
		return nil, fmt.Errorf("there is no engine installed")
	}

	if len(s.transmission) == 0 {
		return nil, fmt.Errorf("there is no transmission installed")
	}

	if len(s.exhaust) == 0 {
		return nil, fmt.Errorf("there is no exhaust installed")
	}

	if s.totalTyre != 4 {
		return nil, fmt.Errorf("insufficient/excess number of tyre")
	}

	if !s.isExistSteeringWheel {
		return nil, fmt.Errorf("there is no steering wheel installed")
	}
	return &Car{
		exhaust:       s.exhaust,
		steeringWheel: s.steeringWheelMaterial,
		tyreSpec:      s.tyreSpec,
		tyre:          s.totalTyre,
		windowFilm:    s.windowFilm,
		engine:        s.engine,
		transmission:  s.transmission,
		seat:          s.totalSeat,
	}, nil
}
