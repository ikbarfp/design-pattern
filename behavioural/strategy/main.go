package strategy

import (
	"log"
	"time"
)

// Execute - Define implementation
func Execute() {

	// Pre-define origin location
	origin := Location{
		ProvinceName: "JAWA BARAT",
		CityName:     "DEPOK",
		DistrictName: "PANCORAN MAS",
		VillageName:  "MAMPANG",
		PostalCode:   "16433",
	}

	// Pre-define destination location
	destination := Location{
		ProvinceName: "DKI JAKARTA",
		CityName:     "JAKARTA SELATAN",
		DistrictName: "CILANDAK",
		VillageName:  "GANDARIA SELATAN",
		PostalCode:   "12420",
	}

	var (
		distance = float64(0)
		duration = 0 * time.Minute
		cost     = float64(0)
	)

	// Set first strategy
	bus := &Bus{}
	trip := NewTrip(bus, origin, destination)

	distance = bus.CalculateDistance(origin, destination)
	duration = bus.CalculateETA(distance)
	cost = bus.CalculateCost(distance)
	log.Println("[By Bus] distance : ", distance)
	log.Println("[By Bus] duration : ", duration)
	log.Println("[By Bus] cost : ", cost)

	// Set new strategy
	train := &Train{}
	trip.SetStrategy(train)

	distance = train.CalculateDistance(origin, destination)
	duration = train.CalculateETA(distance)
	cost = train.CalculateCost(distance)
	log.Println("[By Train] distance : ", distance)
	log.Println("[By Train] duration : ", duration)
	log.Println("[By Train] cost : ", cost)

	// Set new strategy
	motorcycle := &Motorcycle{}
	trip.SetStrategy(motorcycle)

	distance = motorcycle.CalculateDistance(origin, destination)
	duration = motorcycle.CalculateETA(distance)
	cost = motorcycle.CalculateCost(distance)
	log.Println("[By Motorcycle] distance : ", distance)
	log.Println("[By Motorcycle] duration : ", duration)
	log.Println("[By Motorcycle] cost : ", cost)

}
