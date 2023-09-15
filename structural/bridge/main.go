package bridge

func Execute() {
	macOS := &MacOS{}
	macOS.BootingUp()

	windows := &WindowsOS{}
	windows.BootingUp()

	linux := &LinuxOS{}
	linux.BootingUp()

	macChrome := &Chrome{computerOS: macOS}
	macChrome.Open()

	windowsChrome := &Chrome{computerOS: windows}
	windowsChrome.Open()

	linuxChrome := &Chrome{computerOS: linux}
	linuxChrome.Open()

}
