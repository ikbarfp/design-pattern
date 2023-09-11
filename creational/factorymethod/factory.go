package factorymethod

import (
	"fmt"
	"time"
)

type VehicleFactory interface {
	// Produce in this method, we can have
	// specific condition or rules
	// based on the factory object
	Produce(timeNow time.Time) Vehicle
}

type GermanyFactory struct{}

func (gf GermanyFactory) Produce(timeNow time.Time) Vehicle {

	// In this example,
	// Germany Factory will produce Mercedes-Benz car
	// only on monday, wednesday and friday
	weekday := timeNow.Weekday().String()
	switch weekday {

	case time.Monday.String():
		fmt.Println("Produce new Mercedes Benz on monday")
		return &MercedesBenz{}

	case time.Wednesday.String():
		fmt.Println("Produce new Mercedes Benz on wednesday")
		return &MercedesBenz{}

	case time.Friday.String():
		fmt.Println("Produce new Mercedes Benz on wednesday")
		return &MercedesBenz{}

	default:
		fmt.Println("Factory is on a holiday, please come back later")
		return nil
	}
}

type JapanFactory struct{}

func (jf JapanFactory) Produce(timeNow time.Time) Vehicle {

	// In this example,
	// Japan Factory will produce
	// various brand car
	// from monday to saturday
	weekday := timeNow.Weekday().String()
	switch weekday {

	case time.Monday.String():
		fmt.Println("Produce new Toyota on monday")
		return &Toyota{}

	case time.Tuesday.String():
		fmt.Println("Produce new Hino on tuesday")
		return &Hino{}

	case time.Wednesday.String():
		fmt.Println("Produce new Toyota on wednesday")
		return &Toyota{}

	case time.Thursday.String():
		fmt.Println("Produce new Mazda on thursday")
		return &Mazda{}

	case time.Friday.String():
		fmt.Println("Produce new Toyota on friday")
		return &Toyota{}

	case time.Saturday.String():
		fmt.Println("Produce new Mazda on saturday")
		return &Mazda{}

	default:
		fmt.Println("Factory is on a holiday, please come back later")
		return nil
	}
}
