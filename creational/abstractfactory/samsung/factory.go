package samsung

import (
	"fmt"
	"github.com/ikbarfp/design-pattern/creational/abstractfactory"
	"github.com/ikbarfp/design-pattern/creational/abstractfactory/samsung/product"
)

type Samsung struct{}

func (s Samsung) CreatePhone(phoneName abstractfactory.PhoneName) abstractfactory.Phone {

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
