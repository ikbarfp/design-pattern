package product

import "github.com/ikbarfp/design-pattern/creational/abstractfactory"

// GalaxyTabS9Ultra This object act as an
// end-product that implement abstraction behaviour
// from SmartPhone
type GalaxyTabS9Ultra struct{}

func (gts9u *GalaxyTabS9Ultra) GetPrice() float64 {
	return 24999000
}
func (gts9u *GalaxyTabS9Ultra) GetRAM() int {
	return 8
}
func (gts9u *GalaxyTabS9Ultra) GetStorage() int {
	return 512
}
func (gts9u *GalaxyTabS9Ultra) GetProcessor() string {
	return "Snapdragon 8 Gen 2"
}
func (gts9u *GalaxyTabS9Ultra) GetCamera() []string {
	return []string{""}
}
func (gts9u *GalaxyTabS9Ultra) GetOS() abstractfactory.TabletOS {
	return abstractfactory.AndroidOS
}
