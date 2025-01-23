package textcodec

import (
	"errors"
	"io"

	"github.com/diegohce/droneip/ctcodecs"
)

const mimeType = "text/plain"

type textCodec struct{}

func (c *textCodec) MimeType() string {
	return mimeType
}

func (c *textCodec) NewEncoder(w io.Writer) ctcodecs.Encoder {
	return newEncoder(w)
}

func (c *textCodec) NewDecoder(r io.Reader) ctcodecs.Decoder {
	return newDecoder(r)
}

func (c *textCodec) Marshal(v interface{}) ([]byte, error) {
	s, ok := v.(*string)
	if !ok {
		return nil, errors.New("not a string pointer")
	}
	return []byte(*s), nil
}

func (c *textCodec) Unmarshal(data []byte, v interface{}) error {
	if _, ok := v.(*string); !ok {
		return errors.New("not a string pointer")
	}
	*(v.(*string)) = string(data)

	return nil
}

func init() {
	ctcodecs.Register(mimeType, &textCodec{})
}
