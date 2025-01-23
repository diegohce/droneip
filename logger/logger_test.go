package logger_test

import (
	"strings"
	"testing"

	"github.com/diegohce/droneip/logger"
)

func TestLoggerWithFields(t *testing.T) {

	l := logger.NewLogger("name", "diego", "age", 47)

	l.LogInfo("message 1", "dog", "simona").Write()

}

func TestLoggerWithoutFields(t *testing.T) {

	l := logger.NewLogger()

	l.LogInfo("message 2", "dog", "simona").Write()

}

type testLoger string

func (tl *testLoger) WriteLogLine(l string) {
	*tl = testLoger(l)
}

func TestCreditCardMask(t *testing.T) {
	var tl testLoger

	logger.RegisterLogWriter(&tl)

	testCases := []struct {
		input string
	}{
		{"4111111111111111"}, // (VISA, JCB, MasterCard)(len = 16)
		{"346823285239073"},  // (American Express)(len = 15)
	}

	for _, c := range testCases {
		logger.LogInfo(c.input).Write()

		if strings.Contains(string(tl), c.input) {
			t.Errorf("logline contains CC number")
		}

	}
}
