package apple

import (
	"github.com/ikbarfp/design-pattern/creational/abstractfactory"
	"github.com/ikbarfp/design-pattern/creational/abstractfactory/apple/product"
)

type Apple struct{}

func (a Apple) CreatePhone(phoneName abstractfactory.PhoneName) abstractfactory.Phone {

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
