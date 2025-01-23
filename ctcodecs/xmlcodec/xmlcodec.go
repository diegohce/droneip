package xmlcodec

import (
	"encoding/xml"
	"io"

	"github.com/diegohce/droneip/ctcodecs"
)

type xmlCodec struct{}

const mimeType = "application/xml"
const mimeTypeAlt = "text/xml"

func (j *xmlCodec) MimeType() string {
	return mimeType
}

func (j *xmlCodec) NewEncoder(w io.Writer) ctcodecs.Encoder {
	return xml.NewEncoder(w)
}

func (j *xmlCodec) NewDecoder(r io.Reader) ctcodecs.Decoder {
	return xml.NewDecoder(r)
}

func (j *xmlCodec) Marshal(v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}

func (j *xmlCodec) Unmarshal(data []byte, v interface{}) error {
	return xml.Unmarshal(data, v)
}

type xmlCodecAlt struct {
	xmlCodec
}

func (j *xmlCodecAlt) MimeType() string {
	return mimeTypeAlt
}

func init() {
	ctcodecs.Register(mimeType, &xmlCodec{})
	ctcodecs.Register(mimeTypeAlt, &xmlCodecAlt{})
}
