package bridge

func Execute() {
	// From this line we made the browser as parent object of the OS

	// Initialize new mac concrete object
	macOS := &MacOS{}
	macOS.BootingUp()

	// Initialize new windows concrete object
	windows := &WindowsOS{}
	windows.BootingUp()

	// Initialize new linux concrete object
	linux := &LinuxOS{}
	linux.BootingUp()

	// Inject the concrete chrome object by its OS
	// In this example, we use mac as the os
	macChrome := &Chrome{computerOS: macOS}
	macChrome.Open()

	// Inject the concrete chrome object by its OS
	// In this example, we use windows as the os
	windowsChrome := &Chrome{computerOS: windows}
	windowsChrome.Open()

	// Inject the concrete chrome object by its OS
	// In this example, we use linux as the os
	linuxChrome := &Chrome{computerOS: linux}
	linuxChrome.Open()

	// Inject the concrete firefox object by its OS
	// In this example, we use mac as the os
	macFirefox := &Firefox{computerOS: macOS}
	macFirefox.Open()

	// Inject the concrete firefox object by its OS
	// In this example, we use windows as the os
	windowsFirefox := &Firefox{computerOS: windows}
	windowsFirefox.Open()

	// Inject the concrete firefox object by its OS
	// In this example, we use linux as the os
	linuxFirefox := &Firefox{computerOS: linux}
	linuxFirefox.Open()

	// How about we change the paradigm?
	// Make the OS act as the parent object from
	// the browser and see what's happen

	// Initialize new mac object
	macbook := &MacOS{}
	macbook.BootingUp()

	// Try to open Chrome browser with linux version app
	// on macOS and see what happen
	macbook.OpenBrowser(&Chrome{computerOS: &LinuxOS{}})

}
