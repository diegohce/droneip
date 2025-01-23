package logger_test

import (
	"testing"

	"github.com/diegohce/droneip/logger"
)

func TestLoggers(t *testing.T) {

	logger.LogInfo("pepe").Write()

}
