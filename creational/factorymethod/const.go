package factorymethod

type CarType string

const (
	CityCar      CarType = "CITY_CAR"
	SUV          CarType = "SUV"
	MPV          CarType = "MPV"
	Sedan        CarType = "SEDAN"
	StationWagon CarType = "STATION_WAGON"
	Truck        CarType = "TRUCK"
)

type VehicleClassType string

const (
	LowCostGreenCar VehicleClassType = "LCGC"
	General         VehicleClassType = "GENERAL"
	Fleet           VehicleClassType = "FLEET"
	Luxury          VehicleClassType = "LUXURY"
)

type EngineType string

const (
	Diesel      EngineType = "DIESEL"
	TurboDiesel EngineType = "TURBO_DIESEL"
	Gasoline    EngineType = "GASOLINE"
	Turbo       EngineType = "TURBO"
	Hybrid      EngineType = "HYBRID"
	Electric    EngineType = "ELECTRIC"
)
