package apple

import (
	"github.com/ikbarfp/design-pattern/creational/abstractfactory"
	"github.com/ikbarfp/design-pattern/creational/abstractfactory/apple/product"
)

// Apple This object act as a factory
// that produce end-product that implement
// abstraction behaviour from its super class
type Apple struct{}

func (a Apple) CreateSmartPhone(phoneName abstractfactory.PhoneName) abstractfactory.SmartPhone {

	switch phoneName {

	case abstractfactory.Iphone14Pro:
		return &product.Iphone14Pro{}

	case abstractfactory.Iphone14ProMax:
		return &product.Iphone14ProMax{}

	default:
		return nil
	}
}
func (a Apple) CreateLaptop() abstractfactory.Laptop {
	return &product.MacbookPro13Inch{}
}

func (a Apple) CreateTablet() abstractfactory.Tablet {
	return &product.IpadPro13Inch{}
}
