package logger

import (
	"strings"
	"testing"
)

func TestLoggerWithFields(t *testing.T) {

	l := NewLogger("name", "diego", "age", 47)

	l.LogInfo("message 1", "dog", "simona").Write()

}

func TestLoggerWithoutFields(t *testing.T) {

	l := NewLogger()

	l.LogInfo("message 2", "dog", "simona").Write()

}

func TestLoggerLogError(t *testing.T) {
	l := NewLogger()
	l.LogError("message 2", "dog", "simona").Write()
}

func TestAddFields(t *testing.T) {

	l := NewLogger()
	l.AddFields("fixed_field", "fixed_value")
}

func TestLogLevels(t *testing.T) {
	LogInfo("sample error log entry")
	LogWarning("sample error log entry")
	LogError("sample error log entry")
	LogDebug("sample error log entry")
}

type testLoger string

func (tl *testLoger) WriteLogLine(l string) {
	*tl = testLoger(l)
}

func TestCreditCardMask(t *testing.T) {
	var tl testLoger

	RegisterLogWriter(&tl)

	testCases := []struct {
		input string
	}{
		{"4111111111111111"}, // (VISA, JCB, MasterCard)(len = 16)
		{"346823285239073"},  // (American Express)(len = 15)
	}

	for _, c := range testCases {
		LogInfo(c.input).Write()

		if strings.Contains(string(tl), c.input) {
			t.Errorf("logline contains CC number")
		}

	}
}
