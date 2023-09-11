package factorymethod

import "fmt"

// Vehicle this abstraction act as an end-product behaviour
// so any object who implemented this behaviour can act like Vehicle
type Vehicle interface {

	// Build in this method, we can have
	// specific condition or rules
	// depends on object whose act like Vehicle
	Build(carType CarType, classType VehicleClassType)

	// GetCarType in this method, we can have
	// specific condition or rules
	// depends on object whose act like Vehicle
	GetCarType() CarType

	// GetCarClassType in this method, we can have
	// specific condition or rules
	// depends on object whose act like Vehicle
	GetCarClassType() VehicleClassType

	// GetSpecification in this method, we can have
	// specific condition or rules
	// depends on object whose act like Vehicle
	GetSpecification() Specification
}

type Specification struct {
	power, torque float64
	engineType    EngineType
}

type MercedesBenz struct {
	carType       CarType
	specification Specification
}

func (mb *MercedesBenz) Build(carType CarType, classType VehicleClassType) {

	// In this example,
	// MercedesBenz want to produce luxury cars only
	if classType != Luxury {
		fmt.Println("Sorry, we only produce luxury car")
		return
	}

	// In this example,
	// MercedesBenz can produce various car
	// based on car types they capable of
	switch carType {
	case StationWagon:
		fmt.Println("Start building Mercedes Benz E-Class S213 E63 AMG Line . . .")
		mb.carType = StationWagon
		mb.specification = Specification{engineType: Turbo, power: 380, torque: 475}
	case Sedan:
		fmt.Println("Start building Mercedes Benz S-Class W223 S500 AMG Line . . .")
		mb.carType = Sedan
		mb.specification = Specification{engineType: Turbo, power: 410, torque: 500}
	case SUV:
		fmt.Println("Start building Mercedes Benz G-Class W463 G63 AMG Line . . .")
		mb.carType = Sedan
		mb.specification = Specification{engineType: Turbo, power: 410, torque: 550}
	default:
		fmt.Println("Sorry, we can't produce the car")
	}
}
func (mb *MercedesBenz) GetCarType() CarType {
	return mb.carType
}

func (mb *MercedesBenz) GetCarClassType() VehicleClassType {
	return Luxury
}

func (mb *MercedesBenz) GetSpecification() Specification {
	return mb.specification
}

type Hino struct {
	specification Specification
}

func (h *Hino) Build(carType CarType, classType VehicleClassType) {

	// In this example,
	// Hino want to produce fleet cars only
	if classType != Fleet {
		fmt.Println("Sorry, we only produce fleet car")
		return
	}

	// In this example,
	// Hino can produce various car
	// based on car types they capable of
	switch carType {
	case Truck:
		fmt.Println("Start building Hino Dutro 130HD 4x4 . . .")
		h.specification = Specification{engineType: Diesel, power: 140, torque: 175}
	default:
		fmt.Println("Sorry, we only produce truck")
	}
}

func (h *Hino) GetCarType() CarType {
	return Truck
}

func (h *Hino) GetCarClassType() VehicleClassType {
	return Fleet
}

func (h *Hino) GetSpecification() Specification {
	return h.specification
}

type Toyota struct {
	carType       CarType
	classType     VehicleClassType
	specification Specification
}

func (t *Toyota) Build(carType CarType, classType VehicleClassType) {

	// In this example,
	// Toyota can produce all car types except luxury cars
	if classType == Luxury {
		fmt.Println("Sorry, we can't produce the car")
		return
	}

	// In this example,
	// Toyota can produce various car
	// based on car types they capable of
	switch carType {
	case CityCar:
		fmt.Println("Start building Toyota Agya GR Sport . . .")
		t.carType = CityCar
		t.classType = LowCostGreenCar
		t.specification = Specification{engineType: Gasoline, power: 90, torque: 110}
	case SUV:
		fmt.Println("Start building Toyota Land Cruiser GR Sport . . .")
		t.carType = SUV
		t.classType = General
		t.specification = Specification{engineType: TurboDiesel, power: 230, torque: 295}
	case MPV:
		fmt.Println("Start building Toyota Avanza Q TSS . . .")
		t.carType = MPV
		t.classType = General
		t.specification = Specification{engineType: Gasoline, power: 230, torque: 295}
	case Sedan:
		fmt.Println("Start building Toyota Camry Hybrid . . .")
		t.carType = Sedan
		t.classType = General
		t.specification = Specification{engineType: Hybrid, power: 265, torque: 320}
	case Truck:
		fmt.Println("Start building Toyota Dyna 4x4 . . .")
		t.carType = Truck
		t.classType = Fleet
		t.specification = Specification{engineType: Diesel, power: 155, torque: 245}
	default:
		fmt.Println("Sorry, we can't produce the car")
	}
}

func (t *Toyota) GetCarType() CarType {
	return t.carType
}

func (t *Toyota) GetCarClassType() VehicleClassType {
	return t.classType
}

func (t *Toyota) GetSpecification() Specification {
	return t.specification
}

type Mazda struct {
	carType       CarType
	classType     VehicleClassType
	specification Specification
}

func (m *Mazda) Build(carType CarType, classType VehicleClassType) {

	// In this example,
	// Mazda can only produce luxury and general class of car
	switch classType {
	case LowCostGreenCar, Fleet:
		fmt.Println("Sorry, we can't produce lcgc or fleets")
		return
	}

	// In this example,
	// Toyota can produce various car
	// based on car types they capable of
	switch carType {
	case CityCar:
		fmt.Println("Start building Mazda 2 Elite Hatchback . . .")
		m.carType = CityCar
		m.classType = General
		m.specification = Specification{engineType: Gasoline, power: 115, torque: 145}
	case SUV:
		fmt.Println("Start building Mazda MX-30 EV Elite . . .")
		m.carType = SUV
		m.classType = General
		m.specification = Specification{engineType: Electric, power: 185, torque: 250}
	case Sedan:
		fmt.Println("Start building Mazda 6 Elite Sedan . .")
		m.carType = Sedan
		m.classType = General
		m.specification = Specification{engineType: Gasoline, power: 170, torque: 240}
	case StationWagon:
		fmt.Println("Start building Mazda 6 Elite Station Wagon . .")
		m.carType = StationWagon
		m.classType = General
		m.specification = Specification{engineType: Gasoline, power: 170, torque: 240}
	default:
		fmt.Println("Sorry, we can't produce the car")
	}
}

func (m *Mazda) GetCarType() CarType {
	return m.carType
}

func (m *Mazda) GetCarClassType() VehicleClassType {
	return m.classType
}
func (m *Mazda) GetSpecification() Specification {
	return m.specification
}
