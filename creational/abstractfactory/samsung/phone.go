package samsung

import (
	"github.com/ikbarfp/design-pattern/creational/abstractfactory/constant"
)

// GalaxyZFold5 This object act as an
// end-product that implement abstraction behaviour
// from SmartPhone
type GalaxyZFold5 struct{}

func (gzf5 *GalaxyZFold5) GetPrice() float64 {
	return 29900000
}
func (gzf5 *GalaxyZFold5) GetRAM() int {
	return 12
}
func (gzf5 *GalaxyZFold5) GetStorage() int {
	return 512
}
func (gzf5 *GalaxyZFold5) GetProcessor() string {
	return "Snapdragon 8 Gen 2"
}
func (gzf5 *GalaxyZFold5) GetCamera() []string {
	return []string{""}
}
func (gzf5 *GalaxyZFold5) GetOS() constant.PhoneOS {
	return constant.Android
}
func (gzf5 *GalaxyZFold5) IsHaveBluetooth() bool {
	return true
}
func (gzf5 *GalaxyZFold5) IsDualSIM() bool {
	return true
}
