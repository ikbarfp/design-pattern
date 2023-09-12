package abstractfactory

import "github.com/ikbarfp/design-pattern/creational/abstractfactory/constant"

type Price interface {

	// GetPrice This method act as
	// general behaviour of a gadget
	// on Price scope
	GetPrice() float64
}

type Specification interface {

	// GetRAM This method act as
	// general behaviour of a gadget
	// on Specification scope
	GetRAM() int

	// GetStorage This method act as
	// general behaviour of a gadget
	// on Specification scope
	GetStorage() int

	// GetProcessor This method act as
	// general behaviour of a gadget
	// on Specification scope
	GetProcessor() string

	// GetCamera This method act as
	// general behaviour of a gadget
	// on Specification scope
	GetCamera() []string
}

// SmartPhone This interface represent
// product behaviour for smartphones
type SmartPhone interface {
	Price
	Specification

	// GetOS This method can only be used
	// by factories that make smartphones
	GetOS() constant.PhoneOS

	// IsHaveBluetooth This method can only be used
	// by factories that make smartphones
	IsHaveBluetooth() bool

	// IsDualSIM This method can only be used
	// by factories that make smartphones
	IsDualSIM() bool
}

// Laptop This interface represent
// product behaviour for laptops
type Laptop interface {
	Price
	Specification

	// GetOS This method can only be used
	// by factories that make laptops
	GetOS() constant.LaptopOS

	// IsIntelEvo This method can only be used
	// by factories that make laptops
	IsIntelEvo() bool
}

// Tablet This interface represent
// product behaviour for tablets
type Tablet interface {
	Price
	Specification

	// GetOS This method can only be used
	// by factories that make tablets
	GetOS() constant.TabletOS
}

// GadgetFactory This interface act as an abstraction factory
// that produce a factory, not an end-product
type GadgetFactory interface {

	// CreateSmartPhone This method can have various
	// implementations depending on the factory
	// that abstracts this method
	CreateSmartPhone(phoneName constant.PhoneName) SmartPhone

	// CreateLaptop This method can have various
	// implementations depending on the factory
	// that abstracts this method
	CreateLaptop() Laptop

	// CreateTablet This method can have various
	// implementations depending on the factory
	// that abstracts this method
	CreateTablet() Tablet
}
