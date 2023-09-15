package bridge

import (
	"fmt"
	"time"
)

type OS interface {
	// BootingUp Just a gimmick method, please don't mind :D
	BootingUp()

	// OpenBrowser Use this method to see different paradigm as
	// this abstraction become the parent objects and then just
	// call this method then inject the browser
	OpenBrowser(browser Browser)

	// GetName For get OSName
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
