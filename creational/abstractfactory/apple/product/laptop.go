package product

import "github.com/ikbarfp/design-pattern/creational/abstractfactory"

// MacbookPro13Inch This object act as an
// end-product that implement abstraction behaviour
// from Laptop
type MacbookPro13Inch struct{}

func (mp13 *MacbookPro13Inch) GetPrice() float64 {
	return 25000000
}
func (mp13 *MacbookPro13Inch) GetRAM() int {
	return 16
}
func (mp13 *MacbookPro13Inch) GetStorage() int {
	return 512
}
func (mp13 *MacbookPro13Inch) GetProcessor() string {
	return "M2 Pro ARM Chip"
}
func (mp13 *MacbookPro13Inch) GetCamera() []string {
	return []string{"1440p"}
}
func (mp13 *MacbookPro13Inch) GetOS() abstractfactory.LaptopOS {
	return abstractfactory.MacOS
}
func (mp13 *MacbookPro13Inch) IsIntelEvo() bool {
	return false
}
