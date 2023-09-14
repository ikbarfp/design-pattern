package builder

type CarType string

const (
	CoupeCar   CarType = "COUPE_CAR"
	LowCostCar CarType = "LOW_COST_CAR"
	SUVCar     CarType = "SUV_CAR"
)

type Transmission string

const (
	Manual     Transmission = "MANUAL"
	DualClutch Transmission = "DUAL_CLUTCH"
	CVT        Transmission = "CVT"
)
