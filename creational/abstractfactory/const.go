package abstractfactory

type PhoneOS string

const (
	Android PhoneOS = "ANDROID"
	IOS     PhoneOS = "IOS"
)

type PhoneName string

const (
	Iphone14Pro    PhoneName = "IPHONE_14_PRO"
	Iphone14ProMax PhoneName = "IPHONE_14_PRO_MAX"
	GalaxyZFold5   PhoneName = "GALAXY_Z_FOLD_5"
)

type LaptopOS string

const (
	WindowsOS LaptopOS = "WINDOWS"
	MacOS     LaptopOS = "MAC_OS"
	Linux     LaptopOS = "LINUX"
)

type LaptopName string

const (
	MacbookPro13Inch LaptopName = "MACBOOK_PRO_13_INCH"
)

type TabletOS string

const (
	AndroidOS TabletOS = "ANDROID_OS"
	IpadOS    TabletOS = "IPAD_OS"
)

type TabletName string

const (
	GalaxyTabS9Ultra TabletName = "GALAXY_TAB_S9_ULTRA"
)
