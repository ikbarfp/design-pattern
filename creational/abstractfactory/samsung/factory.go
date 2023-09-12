package samsung

import (
	"fmt"
	"github.com/ikbarfp/design-pattern/creational/abstractfactory"
	"github.com/ikbarfp/design-pattern/creational/abstractfactory/samsung/product"
)

// Samsung This object act as a factory
// that produce end-product that implement
// abstraction behaviour from its super class
type Samsung struct{}

func (s Samsung) CreateSmartPhone(phoneName abstractfactory.PhoneName) abstractfactory.SmartPhone {

	switch phoneName {
	case abstractfactory.GalaxyZFold5:
		return nil

	default:
		return nil
	}
}

func (s Samsung) CreateLaptop() abstractfactory.Laptop {
	fmt.Println("Samsung doesn't make a laptop")
	return nil
}

func (s Samsung) CreateTablet() abstractfactory.Tablet {
	return &product.GalaxyTabS9Ultra{}
}
