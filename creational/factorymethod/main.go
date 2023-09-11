package factorymethod

import (
	"fmt"
	"time"
)

func Execute() {
	var factory VehicleFactory
	timeNow := time.Now()

	factory = &JapanFactory{}
	japaneseVehicle := factory.Produce(timeNow)

	// i.e -> we want to build a vehicle from
	// Japan factory with some criteria
	// (as arguments from method Build()) :
	// it should be SUV and have a general class
	japaneseVehicle.Build(SUV, General)
	fmt.Println(japaneseVehicle.GetCarType())
	fmt.Println(japaneseVehicle.GetCarClassType())
	fmt.Println(japaneseVehicle.GetSpecification())

	// i.e -> we want to build a vehicle from
	// Japan factory with some criteria
	// (as arguments from method Build()) :
	// it should be sedan and have a luxury class
	// but in this example, there is no luxury sedan
	// could be produced on Japan factory
	japaneseVehicle.Build(Sedan, Luxury)
	fmt.Println(japaneseVehicle.GetCarType())
	fmt.Println(japaneseVehicle.GetCarClassType())
	fmt.Println(japaneseVehicle.GetSpecification())

	factory = &GermanyFactory{}
	bavarianVehicle := factory.Produce(timeNow)

	// i.e -> we want to build a vehicle from
	// Germany factory with some criteria
	// (as arguments from method Build()) :
	// it should be station wagon and have a luxury class
	bavarianVehicle.Build(StationWagon, Luxury)
	fmt.Println(bavarianVehicle.GetCarType())
	fmt.Println(bavarianVehicle.GetCarClassType())
	fmt.Println(bavarianVehicle.GetSpecification())

	// i.e -> we want to build a vehicle from
	// Germany factory with some criteria
	// (as arguments from method Build()) :
	// it should be SUV and have a low-cost green car class
	// but in this example, there is no low-cost green car SUV
	// could be produced on Germany factory
	bavarianVehicle.Build(SUV, LowCostGreenCar)
	fmt.Println(bavarianVehicle.GetCarType())
	fmt.Println(bavarianVehicle.GetCarClassType())
	fmt.Println(bavarianVehicle.GetSpecification())
}
