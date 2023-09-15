package bridge

type BrowserName string

const (
	ChromeBrowser  BrowserName = "CHROME"
	FirefoxBrowser BrowserName = "FIREFOX"
)

type OSName string

const (
	MacComputer     OSName = "MAC OS"
	WindowsComputer OSName = "WINDOWS"
	LinuxComputer   OSName = "LINUX"
)
