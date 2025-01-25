package xmlcodec

import (
	"bytes"
	"testing"
)

type Guy struct {
	Name string
	Age  int
	Pets []string
}

var someGuy Guy = Guy{
	Name: "Diego",
	Age:  45,
	Pets: []string{"Simona", "Ada"},
}

func TestCodec(t *testing.T) {

	codec := &xmlCodec{}

	var buf bytes.Buffer

	if err := codec.NewEncoder(&buf).Encode(&someGuy); err != nil {
		t.Error(err)
	}

	var anotherGuy Guy

	if err := codec.NewDecoder(&buf).Decode(&anotherGuy); err != nil {
		t.Error(err)
	}

	b, _ := codec.Marshal(&someGuy)

	codec.Unmarshal(b, &anotherGuy)

	if codec.MimeType() != mimeType {
		t.Errorf("mime type: got %s want %s", codec.MimeType(), mimeType)
	}
}

func TestCodecAlt(t *testing.T) {

	codec := &xmlCodecAlt{}

	var buf bytes.Buffer

	if err := codec.NewEncoder(&buf).Encode(&someGuy); err != nil {
		t.Error(err)
	}

	var anotherGuy Guy

	if err := codec.NewDecoder(&buf).Decode(&anotherGuy); err != nil {
		t.Error(err)
	}

	b, _ := codec.Marshal(&someGuy)

	codec.Unmarshal(b, &anotherGuy)

	if codec.MimeType() != mimeTypeAlt {
		t.Errorf("mime type: got %s want %s", codec.MimeType(), mimeTypeAlt)
	}
}
