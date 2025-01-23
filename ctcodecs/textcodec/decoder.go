package textcodec

import (
	"bytes"
	"errors"
	"io"

	"github.com/diegohce/droneip/ctcodecs"
)

type decoder struct {
	r io.Reader
}

func newDecoder(r io.Reader) ctcodecs.Decoder {
	return &decoder{r}
}

func (d *decoder) Decode(i interface{}) error {
	var buf bytes.Buffer

	if _, ok := i.(*string); !ok {
		return errors.New("not a string pointer")
	}

	if _, err := buf.ReadFrom(d.r); err != nil {
		return err
	}

	*(i.(*string)) = buf.String()

	return nil
}
