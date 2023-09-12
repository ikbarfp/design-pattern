package product

import "github.com/ikbarfp/design-pattern/creational/abstractfactory"

// IpadPro13Inch This object act as an
// end-product that implement abstraction behaviour
// from Tablet
type IpadPro13Inch struct{}

func (ip13 *IpadPro13Inch) GetPrice() float64 {
	return 18000000
}
func (ip13 *IpadPro13Inch) GetRAM() int {
	return 6
}
func (ip13 *IpadPro13Inch) GetStorage() int {
	return 512
}
func (ip13 *IpadPro13Inch) GetProcessor() string {
	return "M1 Pro ARM Chip"
}
func (ip13 *IpadPro13Inch) GetCamera() []string {
	return []string{"720p"}
}
func (ip13 *IpadPro13Inch) GetOS() abstractfactory.TabletOS {
	return abstractfactory.IpadOS
}
