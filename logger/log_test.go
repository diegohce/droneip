package logger

import (
	"bytes"
	"testing"
)

func TestLogger(t *testing.T) {

	l := Log(Info, "This is the message", "extra1", "value1", "extra2", 2)

	l.SetField("extra3", 3.14).Write()

	buf := bytes.Buffer{}

	l.WriteTo(&buf)

}
