package bridge

import (
	"fmt"
)

type Browser interface {
	Open()
	GetName() BrowserName
}

type Chrome struct {
	// computerOS This property need to be injected on concrete object,
	// because it will act as a bridge from current abstraction to another
	// abstraction
	computerOS OS
	version    string
}

func (c *Chrome) Open() {
	// We can put our logic here because this concrete object
	// can have different business case to another concrete object
	// from same abstraction

	// In this example,
	// Chrome can only be opened from Mac & Windows
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
	// computerOS This property need to be injected on concrete object,
	// because it will act as a bridge from current abstraction to another
	// abstraction
	computerOS OS
	version    string
}

func (f *Firefox) Open() {
	// We can put our logic here because this concrete object
	// can have different business case to another concrete object
	// from same abstraction

	// In this example,
	// Firefox can be opened from all OS
	switch f.computerOS.GetName() {
	case MacComputer:
		f.version = "v3.0.0"
		fmt.Println("Opening Firefox Browser from", f.computerOS.GetName()+" with version", f.version)
	case WindowsComputer:
		f.version = "v4.0.0"
		fmt.Println("Opening Firefox Browser from", f.computerOS.GetName()+" with version", f.version)
	default:
		fmt.Println("Opening Firefox Browser from", f.computerOS.GetName()+" with version", f.version)
	}
}

func (f *Firefox) GetName() BrowserName {
	return FirefoxBrowser
}
