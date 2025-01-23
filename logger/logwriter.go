package logger

// LogWriter interface to implement to make a log writer.
type LogWriter interface {
	// WriteLogLine receives a JSON string without the ending \n
	WriteLogLine(string)
}

var loggers []LogWriter

// RegisterLogWriter register the log writers where the log line
// will be propagated.
func RegisterLogWriter(lw LogWriter) {
	loggers = append(loggers, lw)
}
