package logger

import (
	"io"
	"os"
)

// LogLevel is one of Info, Error, Debug, Warn | Warning
type LogLevel int8

const (
	Info LogLevel = iota
	Error
	Debug
	Warn
)
const Warning = Warn

func (l LogLevel) String() string {
	switch l {
	case Error:
		return "error"
	case Debug:
		return "debug"
	case Warn:
		return "warn"
	}
	return "info"
}

// LogEntry is the log line that will be sent to the default log channel.
type LogEntry struct {
	message string
	fields  map[string]interface{}
	attrs   []interface{}
	level   string
	//caller  string
}

// Log creates a new log entry with `level` (info, error, debug, warn).
func Log(level LogLevel, message string, fields ...interface{}) LogEntry {

	l := LogEntry{
		level:   level.String(),
		message: message,
		fields:  make(map[string]interface{}),
	}

	if fields != nil {
		var key string

		for i, f := range fields {
			if ((i + 1) % 2) != 0 {
				key = f.(string)
			} else {
				l.fields[key] = f
			}
		}
	}

	return l
}

// LogInfo creates a new log entry with `level` info
func LogInfo(message string, fields ...interface{}) LogEntry {
	return Log(Info, message, fields...)
}

// LogError creates a new log entry with `level` error
func LogError(message string, fields ...interface{}) LogEntry {
	return Log(Error, message, fields...)
}

// LogWarning creates a new log entry with `level` warning
func LogWarning(message string, fields ...interface{}) LogEntry {
	return Log(Warn, message, fields...)
}

// LogDebug creates a new log entry with `level` debug
func LogDebug(message string, fields ...interface{}) LogEntry {
	return Log(Debug, message, fields...)
}

// Write sends the logEntry to de logging channel.
func (l LogEntry) Write() {
	writeLogLine(os.Stdout, l)
}

func (l LogEntry) WriteTo(w io.Writer) {
	writeLogLine(w, l)
}

// Level returns the level to be logged to.
func (l LogEntry) Level() string {
	return l.level
}

// Message returns the message to be logged.
func (l LogEntry) Message() string {
	return l.message
}

// Fields returns the extra info for the log entry.
func (l LogEntry) Fields() map[string]interface{} {
	return l.fields
}

// SetField adds an extra_field with name 'name' of value 'value'.
func (l LogEntry) SetField(name string, value interface{}) LogEntry {
	l.fields[name] = value
	return l
}

func (l LogEntry) Attributes() []interface{} {
	return l.attrs
}

func (l LogEntry) SetAttribute(name string, value interface{}) LogEntry {
	l.attrs = append(l.attrs, name, value)
	return l
}
