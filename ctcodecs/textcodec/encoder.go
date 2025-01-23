package textcodec

import (
	"errors"
	"io"

	"github.com/diegohce/droneip/ctcodecs"
)

type encoder struct {
	w io.Writer
}

func newEncoder(w io.Writer) ctcodecs.Encoder {
	return &encoder{w}
}

func (e *encoder) Encode(i interface{}) error {
	s, ok := i.(*string)
	if !ok {
		return errors.New("not a string pointer")
	}
	_, err := io.WriteString(e.w, *s)
	return err
}
