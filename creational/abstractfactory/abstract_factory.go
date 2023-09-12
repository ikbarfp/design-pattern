package abstractfactory

type Price interface {
	GetPrice() float64
}

type Specification interface {
	GetRAM() int
	GetStorage() int
	GetProcessor() string
	GetCamera() []string
}

type Phone interface {
	Price
	Specification
	GetOS() PhoneOS
	IsHaveBluetooth() bool
	IsDualSIM() bool
}

type Laptop interface {
	Price
	Specification
	GetOS() LaptopOS
	IsIntelEvo() bool
}

type Tablet interface {
	Price
	Specification
	GetOS() TabletOS
}

type GadgetFactory interface {
	CreatePhone(phoneName PhoneName) Phone
	CreateLaptop() Laptop
	CreateTablet() Tablet
}
