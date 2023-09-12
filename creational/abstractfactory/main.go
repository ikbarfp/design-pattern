package abstractfactory

import (
	"fmt"
	"github.com/ikbarfp/design-pattern/creational/abstractfactory/constant"
)

func Execute() {
	var gadgetFactory GadgetFactory

	gadgetFactory = &Apple{}

	iphone14Pro := gadgetFactory.CreateSmartPhone(constant.Iphone14Pro)
	fmt.Println(iphone14Pro.GetCamera())
	fmt.Println(iphone14Pro.GetOS())
	fmt.Println(iphone14Pro.GetPrice())
	fmt.Println(iphone14Pro.GetProcessor())
	fmt.Println(iphone14Pro.GetRAM())
	fmt.Println(iphone14Pro.GetStorage())
	fmt.Println(iphone14Pro.IsDualSIM())
	fmt.Println(iphone14Pro.IsHaveBluetooth())

}
