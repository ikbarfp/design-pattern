package strategy

import "fmt"

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

	// Set first strategy
	bus := &Bus{}
	trip := NewTrip(bus, origin, destination)

	fmt.Println()
	fmt.Printf("\n>>> [By Bus] distance : %+v\n", bus.CalculateDistance(origin, destination))
	fmt.Printf("\n>>> [By Bus] duration : %+v\n", bus.CalculateETA(bus.CalculateDistance(origin, destination)).Minutes())
	fmt.Printf("\n>>> [By Bus] cost : %+v\n", bus.CalculateCost())

	// Set new strategy
	train := &Train{}
	trip.SetStrategy(train)

	fmt.Println()
	fmt.Printf("\n>>> [By Train] distance : %+v\n", train.CalculateDistance(origin, destination))
	fmt.Printf("\n>>> [By Train] duration : %+v\n", train.CalculateETA(train.CalculateDistance(origin, destination)).Minutes())
	fmt.Printf("\n>>> [By Train] cost : %+v\n", train.CalculateCost())

	// Set new strategy
	motorcycle := &Motorcycle{}
	trip.SetStrategy(motorcycle)

	fmt.Println()
	fmt.Printf("\n>>> [By Motorcycle] distance : %+v\n", motorcycle.CalculateDistance(origin, destination))
	fmt.Printf("\n>>> [By Motorcycle] duration : %+v\n", motorcycle.CalculateETA(motorcycle.CalculateDistance(origin, destination)).Minutes())
	fmt.Printf("\n>>> [By Motorcycle] cost : %+v\n", motorcycle.CalculateCost())

}
