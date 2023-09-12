package apple

import (
	"github.com/ikbarfp/design-pattern/creational/abstractfactory/constant"
)

// Iphone14Pro This object act as an
// end-product that implement abstraction behaviour
// from SmartPhone
type Iphone14Pro struct{}

func (ip14 *Iphone14Pro) GetPrice() float64 {
	return 20000000
}

func (ip14 *Iphone14Pro) GetRAM() int {
	return 6
}

func (ip14 *Iphone14Pro) GetStorage() int {
	return 256
}

func (ip14 *Iphone14Pro) GetProcessor() string {
	return "A16 Pro Bionic"
}

func (ip14 *Iphone14Pro) GetCamera() []string {
	return []string{}
}

func (ip14 *Iphone14Pro) GetOS() constant.PhoneOS {
	return constant.IOS
}

func (ip14 *Iphone14Pro) IsHaveBluetooth() bool {
	return true
}

func (ip14 *Iphone14Pro) IsDualSIM() bool {
	return true
}

type Iphone14ProMax struct{}

func (ip14pm *Iphone14ProMax) GetPrice() float64 {
	return 20000000
}

func (ip14pm *Iphone14ProMax) GetRAM() int {
	return 6
}

func (ip14pm *Iphone14ProMax) GetStorage() int {
	return 256
}

func (ip14pm *Iphone14ProMax) GetProcessor() string {
	return "A16 Pro Bionic"
}

func (ip14pm *Iphone14ProMax) GetCamera() []string {
	return []string{}
}

func (ip14pm *Iphone14ProMax) GetOS() constant.PhoneOS {
	return constant.IOS
}

func (ip14pm *Iphone14ProMax) IsHaveBluetooth() bool {
	return true
}

func (ip14pm *Iphone14ProMax) IsDualSIM() bool {
	return true
}
