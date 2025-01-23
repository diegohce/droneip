package logger

// Logger is a logging facility with fixed fields written in every log line
type Logger struct {
	fields []interface{}
}

// NewLogger creates a new Logger with fixed fields
func NewLogger(fields ...interface{}) *Logger {

	l := Logger{
		fields: fields,
	}

	return &l
}

// AddFields appends a new fixed field to Logger
func (l *Logger) AddFields(fields ...interface{}) {
	l.fields = append(l.fields, fields...)
}

// Log creates a new log entry with `level` (info, error, debug, warn).
func (l *Logger) Log(level LogLevel, message string, fields ...interface{}) LogEntry {

	allFields := l.fields
	allFields = append(allFields, fields...)

	return Log(level, message, allFields...)
}

// LogInfo creates a new log entry with `level` info
func (l *Logger) LogInfo(message string, fields ...interface{}) LogEntry {
	return l.Log(Info, message, fields...)
}

// LogError creates a new log entry with `level` Error
func (l *Logger) LogError(message string, fields ...interface{}) LogEntry {
	return l.Log(Error, message, fields...)
}
