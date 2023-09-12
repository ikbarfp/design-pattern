package abstractfactory

import (
	"fmt"
	"github.com/ikbarfp/design-pattern/creational/abstractfactory/constant"
	"github.com/ikbarfp/design-pattern/creational/abstractfactory/samsung"
)

// Samsung This object act as a factory
// that produce end-product that implement
// abstraction behaviour from its super class
type Samsung struct{}

func (s Samsung) CreateSmartPhone(phoneName constant.PhoneName) SmartPhone {

	switch phoneName {
	case constant.GalaxyZFold5:
		return &samsung.GalaxyZFold5{}

	default:
		return nil
	}
}

func (s Samsung) CreateLaptop() Laptop {
	fmt.Println("Samsung doesn't make a laptop")
	return nil
}

func (s Samsung) CreateTablet() Tablet {
	return &samsung.GalaxyTabS9Ultra{}
}
