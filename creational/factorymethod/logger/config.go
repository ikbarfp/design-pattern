package logger

// PlatformConfig Object that contains
// specific config for Platform. Define
// your properties as you needed here.
type PlatformConfig struct {
	Env string
}

// NewRelicConfig Object that contains
// specific config for NewRelic. Define
// your properties as you needed here.
type NewRelicConfig struct {
	AppName, License string
}

// FileConfig Object that contains
// specific config for File. Define
// your properties as you needed here.
type FileConfig struct {
}
