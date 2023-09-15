package logger

func Execute() {
	var logger Logger

	// Can be customizable as you needed
	input := "NEW_RELIC"

	// Can be customizable as you needed
	payload := StdLogPayload{
		Request:    "",
		Response:   "",
		HttpMethod: "",
		Url:        "",
		Header:     "",
	}

	// Business rules to decide
	// which logger will be used
	switch input {

	case "PLATFORM":
		logger = &PlatformLogger{}
	default:
		logger = &NewRelicLogger{}
	}

	// Call the logger producer
	producer, err := logger.NewProducer()
	if err != nil {
		// Do something
	}

	// Do write the log
	producer.Write(payload)
}
