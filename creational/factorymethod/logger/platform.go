package logger

// Platform This object as an end-product of LogOutput Signature
// You can define your Platform properties as you needed here
type Platform struct {
	Token string
}

// Write with specific rules & behaviour for Platform output
func (p Platform) Write(payload StdLogPayload) {
	// Do write log with platform behaviour here
}

func (p Platform) LoadConfig() *PlatformConfig {
	// Load your related config here
	env := "MY_ENV"

	return &PlatformConfig{
		Env: env,
	}
}

func (p Platform) Login() (string, error) {
	// Login to platform auth here
	return "", nil
}
