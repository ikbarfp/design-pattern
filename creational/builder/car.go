package builder

type Car struct {
	exhaust, steeringWheel, tyreSpec, windowFilm, engine string
	seat, tyre                                           int
	transmission                                         Transmission
}

type CarBuilder interface {
	AddExhaust()
	AddSeat()
	AddSteeringWheel()
	AddTyre()
	AddWindow()
	BuildEngine()
	BuildTransmission()
	BuildCar() *Car
}
