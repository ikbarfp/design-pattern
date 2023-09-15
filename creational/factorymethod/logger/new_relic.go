package logger

// NewRelic This object as an end-product of LogOutput Signature
// You can define your NewRelic properties as you needed here
type NewRelic struct{}

// Write with specific rules & behaviour for NewRelic output
func (n NewRelic) Write(payload StdLogPayload) {
	// Do write log with new relic behaviour here
}

func (n NewRelic) LoadConfig() *NewRelicConfig {
	// Load your related config here
	appName := "MY_APP_NAME"
	license := "MY_LICENSE"

	return &NewRelicConfig{
		AppName: appName,
		License: license,
	}
}
