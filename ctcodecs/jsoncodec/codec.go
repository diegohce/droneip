package jsoncodec

import (
	"encoding/json"
	"io"

	"github.com/diegohce/droneip/ctcodecs"
)

type jsonCodec struct{}

const mimeType = "application/json"

func (j *jsonCodec) MimeType() string {
	return mimeType
}

func (j *jsonCodec) NewEncoder(w io.Writer) ctcodecs.Encoder {
	return json.NewEncoder(w)
}

func (j *jsonCodec) NewDecoder(r io.Reader) ctcodecs.Decoder {
	return json.NewDecoder(r)
}

func (j *jsonCodec) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (j *jsonCodec) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func init() {
	ctcodecs.Register(mimeType, &jsonCodec{})
}
