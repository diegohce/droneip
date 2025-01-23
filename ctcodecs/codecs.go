package ctcodecs

import (
	"fmt"
	"io"
)

type Encoder interface {
	Encode(i interface{}) error
}

type Decoder interface {
	Decode(i interface{}) error
}

type Codec interface {
	MimeType() string
	NewEncoder(w io.Writer) Encoder
	NewDecoder(r io.Reader) Decoder
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

var codecs map[string]Codec = map[string]Codec{}
var codecsList []string

// New returns a codec for the mime type in mimeType. Error if there's no codec
// for specified mime type.
func New(mimeType string) (Codec, error) {
	codec, ok := codecs[mimeType]
	if !ok {
		return nil, fmt.Errorf("no codec for %s", mimeType)
	}
	return codec, nil
}

// Register adds a new codec for mimeType.
func Register(mimeType string, c Codec) {
	codecs[mimeType] = c
	codecsList = append(codecsList, mimeType)
}

func List() []string {
	return codecsList
}
