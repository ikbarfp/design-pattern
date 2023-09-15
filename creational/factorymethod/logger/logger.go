package logger

// Logger The factory to initiate a signature
// that logger must have / can do.
type Logger interface {
	NewProducer() (LogOutput, error)
}

// PlatformLogger This concrete object
// represent as a log producer.
// You can define your PlatformLogger
// properties as you needed here.
type PlatformLogger struct {
	config *PlatformConfig
}

// NewProducer Initialize a log producer
// based on the concrete objects.
func (p *PlatformLogger) NewProducer() (LogOutput, error) {
	platform := &Platform{}

	// Get config
	p.config = platform.LoadConfig()

	// Need to log in
	token, err := platform.Login()
	if err != nil {
		return nil, err
	}

	platform.Token = token

	// So on and so forth . . .

	return platform, nil
}

// NewRelicLogger This concrete object
// represent as a log producer.
// You can define your NewRelicLogger
// properties as you needed here.
type NewRelicLogger struct {
	// Define your properties here
	config *NewRelicConfig
}

// NewProducer Initialize a log producer
// based on the concrete objects.
func (n *NewRelicLogger) NewProducer() (LogOutput, error) {
	newRelic := &NewRelic{}

	// Get config
	n.config = newRelic.LoadConfig()

	// So on and so forth . . .

	return newRelic, nil
}

// FileLogger This concrete object
// represent as a log producer.
// You can define your FileLogger
// properties as you needed here.
type FileLogger struct {
	config *FileConfig
}

// NewProducer Initialize a log producer
// based on the concrete objects.
func (f *FileLogger) NewProducer() (LogOutput, error) {
	file := &File{}

	// Get config
	f.config = file.LoadConfig()

	// So on and so forth . . .

	return file, nil
}
