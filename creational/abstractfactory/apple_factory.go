package abstractfactory

import (
	"github.com/ikbarfp/design-pattern/creational/abstractfactory/apple"
	"github.com/ikbarfp/design-pattern/creational/abstractfactory/constant"
)

// Apple This object act as a factory
// that produce end-product that implement
// abstraction behaviour from its super class
type Apple struct{}

func (a Apple) CreateSmartPhone(phoneName constant.PhoneName) SmartPhone {

	switch phoneName {

	case constant.Iphone14Pro:
		return &apple.Iphone14Pro{}

	case constant.Iphone14ProMax:
		return &apple.Iphone14ProMax{}

	default:
		return nil
	}
}
func (a Apple) CreateLaptop() Laptop {
	return &apple.MacbookPro13Inch{}
}

func (a Apple) CreateTablet() Tablet {
	return &apple.IpadPro13Inch{}
}
