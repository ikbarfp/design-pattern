package logger

// StdLogPayload Standard object for logging
type StdLogPayload struct {
	Request    string
	Response   string
	HttpMethod string
	Url        string
	Header     string
}

// LogOutput The product of a factory to initiate
// a signature that log output must have / can do
type LogOutput interface {
	Write(payload StdLogPayload)
}
