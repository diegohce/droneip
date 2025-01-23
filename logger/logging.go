package logger

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"regexp"
	"time"
)

var (
	reCC = regexp.MustCompile(`(?:4[0-9]{12}(?:[0-9]{3})?|[25][1-7][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})`)
)

func parseLevel(s string) (slog.Level, error) {
	var level slog.Level
	var err = level.UnmarshalText([]byte(s))
	return level, err
}

func getExtraFields(fields map[string]interface{}) []interface{} {
	result := make([]interface{}, 0, len(fields)*2)
	for key, value := range fields {
		result = append(result, key, value)
	}
	return result
}

func writeLogLine(w io.Writer, le LogEntry) {

	logLevel := slog.LevelInfo
	if l := os.Getenv("KIUGO_LOGGER_LOGLEVEL"); l != "" {
		var err error
		logLevel, err = parseLevel(l)
		if err != nil {
			logLevel = slog.LevelInfo
			slog.With(err).
				Error("Error parsing log level in KIUGO_LOGGER_LOGLEVEL. Using default 'info' level")
		}
	}

	logBuffer := bytes.Buffer{}

	// Use io.MultiWriter to write to both os.Stdout and logBuffer
	//mainLogger := slog.New(slog.NewJSONHandler(io.MultiWriter(w, &logBuffer), &slog.HandlerOptions{
	mainLogger := slog.New(slog.NewJSONHandler(&logBuffer, &slog.HandlerOptions{
		Level: logLevel,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if groups == nil {
				if a.Key == "time" {
					a.Key = "@timestamp"
					a.Value = slog.StringValue(a.Value.Time().Format(time.RFC3339))
				}
				if a.Key == "msg" {
					a.Key = "message"
				}
			}
			return a
		},
	}))

	lvl, err := parseLevel(le.Level())
	if err != nil {
		slog.With(err).
			Error("Error parsing loglevel. Using 'info'")
		lvl = slog.LevelInfo
	}

	args := le.Attributes()
	args = append(args, slog.Group("extra_fields", getExtraFields(le.Fields())...))
	mainLogger.Log(context.Background(), lvl, le.Message(), args...)

	var maskedLogLine string

	ccNo := reCC.FindString(logBuffer.String())
	if len(ccNo) > 0 {
		maskedLogLine = reCC.ReplaceAllString(logBuffer.String(), fmt.Sprintf("************%s", ccNo[len(ccNo)-4:]))

	} else {
		maskedLogLine = logBuffer.String()
	}
	fmt.Fprint(w, maskedLogLine)

	// Send the logged line to all registered loggers
	for _, log_writer := range loggers {
		//log_writer.WriteLogLine(logBuffer.String())
		log_writer.WriteLogLine(maskedLogLine)
	}
}
