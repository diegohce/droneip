package msgpackcodec

import (
	"bytes"
	"io"

	"github.com/vmihailenco/msgpack"

	"github.com/diegohce/droneip/ctcodecs"
)

const mimeType = "application/msgpack"

type msgpackCodec struct{}

func (j *msgpackCodec) MimeType() string {
	return mimeType
}

func (j *msgpackCodec) NewEncoder(w io.Writer) ctcodecs.Encoder {
	return msgpack.NewEncoder(w).UseJSONTag(true)
}

func (j *msgpackCodec) NewDecoder(r io.Reader) ctcodecs.Decoder {
	return msgpack.NewDecoder(r).UseJSONTag(true)
}

func (j *msgpackCodec) Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	err := msgpack.NewEncoder(&buf).UseJSONTag(true).Encode(v)
	return buf.Bytes(), err
}

func (j *msgpackCodec) Unmarshal(data []byte, v interface{}) error {
	return msgpack.NewDecoder(bytes.NewReader(data)).UseJSONTag(true).Decode(v)
}

func init() {
	ctcodecs.Register(mimeType, &msgpackCodec{})
}
