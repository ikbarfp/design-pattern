package bridge

import (
	"fmt"
)

type Browser interface {
	Open()
	GetName() BrowserName
}

type Chrome struct {
	computerOS OS
	version    string
}

func (c *Chrome) Open() {
	switch c.computerOS.GetName() {
	case MacComputer:
		c.version = "v1.0.0"
		fmt.Println("Opening Chrome Browser from", c.computerOS.GetName()+" with version", c.version)
	case WindowsComputer:
		c.version = "v2.0.0"
		fmt.Println("Opening Chrome Browser from", c.computerOS.GetName()+" with version", c.version)
	default:
		fmt.Println("Sorry Chrome Browser doesn't support", c.computerOS.GetName(), "OS")
	}
}

func (c *Chrome) GetName() BrowserName {
	return ChromeBrowser
}

type Firefox struct {
	computerOS OS
	version    string
}

func (f *Firefox) Open() {
	switch f.computerOS.GetName() {
	case MacComputer:
		f.version = "v3.0.0"
		fmt.Println("Opening Firefox Browser from", f.computerOS.GetName()+"with version", f.version)
	case WindowsComputer:
		f.version = "v4.0.0"
		fmt.Println("Opening Firefox Browser from", f.computerOS.GetName()+"with version", f.version)
	default:
		fmt.Println("Sorry Firefox Browser doesn't support", f.computerOS.GetName(), "OS!")
	}
}

func (f *Firefox) GetName() BrowserName {
	return FirefoxBrowser
}
