package bridge

import (
	"fmt"
	"time"
)

type OS interface {
	BootingUp()
	OpenBrowser(browser Browser)
	GetName() OSName
}

type MacOS struct{}

func (m *MacOS) BootingUp() {
	fmt.Println("Loading Mac OS, please wait . . .")
	time.Sleep(1 * time.Second)
	fmt.Println("Welcome to Mac!")
}

func (m *MacOS) OpenBrowser(browser Browser) {
	browser.Open()
}

func (m *MacOS) GetName() OSName {
	return MacComputer
}

type WindowsOS struct{}

func (w *WindowsOS) BootingUp() {
	fmt.Println("Loading Windows OS, please wait . . .")
	time.Sleep(2 * time.Second)
	fmt.Println("Welcome to Windows!")
}
func (w *WindowsOS) OpenBrowser(browser Browser) {
	browser.Open()
}

func (w *WindowsOS) GetName() OSName {
	return WindowsComputer
}

type LinuxOS struct{}

func (w *LinuxOS) BootingUp() {
	fmt.Println("Loading Linux OS, please wait . . .")
	time.Sleep(1 * time.Second)
	fmt.Println("Welcome to Linux!")
}
func (w *LinuxOS) OpenBrowser(browser Browser) {
	browser.Open()
}

func (w *LinuxOS) GetName() OSName {
	return LinuxComputer
}
