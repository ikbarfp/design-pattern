package logger

// File This object as an end-product of LogOutput Signature
// You can define your File properties as you needed here
type File struct{}

// Write with specific rules & behaviour for File output
func (f File) Write(payload StdLogPayload) {
	// Do write log with file behaviour here
}

func (f File) LoadConfig() *FileConfig {
	// Load your related config here

	return &FileConfig{}
}
