package gobcodec

import (
	"bytes"
	"encoding/gob"
	"io"

	"github.com/diegohce/droneip/ctcodecs"
)

const mimeType = "application/gob"

type gobCodec struct{}

func (j *gobCodec) MimeType() string {
	return mimeType
}

func (j *gobCodec) NewEncoder(w io.Writer) ctcodecs.Encoder {
	return gob.NewEncoder(w)
}

func (j *gobCodec) NewDecoder(r io.Reader) ctcodecs.Decoder {
	return gob.NewDecoder(r)
}

func (j *gobCodec) Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(v)
	return buf.Bytes(), err
}

func (j *gobCodec) Unmarshal(data []byte, v interface{}) error {
	return gob.NewDecoder(bytes.NewReader(data)).Decode(v)
}

func init() {
	ctcodecs.Register(mimeType, &gobCodec{})
}
