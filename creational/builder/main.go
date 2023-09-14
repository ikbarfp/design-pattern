package builder

import (
	"fmt"
	"log"
)

func Execute() {

	lowCostCarBuilder := &LowCostCarBuilder{}

	lowCostCar, err := lowCostCarBuilder.
		BuildEngine().
		BuildTransmission(CVT).
		AddExhaust().
		AddTyre(4).
		AddSteeringWheel().
		BuildCar()
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println("Low Cost Car exhaust is", lowCostCar.exhaust)
	fmt.Println("Low Cost Car steering wheel material is", lowCostCar.steeringWheel)
	fmt.Println("Low Cost Car tyre spec is", lowCostCar.tyreSpec)
	fmt.Println("Low Cost Car window film is", lowCostCar.windowFilm)
	fmt.Println("Low Cost Car engine is", lowCostCar.engine)
	fmt.Println("Low Cost Car total seats are", lowCostCar.seat)
	fmt.Println("Low Cost Car total tyre are", lowCostCar.tyre)
	fmt.Println("Low Cost Car transmission is", lowCostCar.transmission)
	fmt.Println()

	sportsCarBuilder := &SportsCarBuilder{}

	sportsCar, err := sportsCarBuilder.
		BuildEngine().
		BuildTransmission(DualClutch).
		AddExhaust().
		AddSeat(2).
		AddTyre(4).
		AddSteeringWheel().
		AddWindowFilm().
		BuildCar()
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println("Sports Car exhaust is", sportsCar.exhaust)
	fmt.Println("Sports Car steering wheel material is", sportsCar.steeringWheel)
	fmt.Println("Sports Car tyre spec is", sportsCar.tyreSpec)
	fmt.Println("Sports Car window film is", sportsCar.windowFilm)
	fmt.Println("Sports Car engine is", sportsCar.engine)
	fmt.Println("Sports Car total seats are", sportsCar.seat)
	fmt.Println("Sports Car total tyre are", sportsCar.tyre)
	fmt.Println("Sports Car transmission is", sportsCar.transmission)

}
