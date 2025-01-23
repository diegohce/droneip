package textcodec

import (
	"bytes"
	"testing"
)

var raw = `hello, world!`

func TestCodec(t *testing.T) {

	codec := &textCodec{}

	var buf bytes.Buffer

	if err := codec.NewEncoder(&buf).Encode(&raw); err != nil {
		t.Fatal(err)
	}

	var result string

	if err := codec.NewDecoder(&buf).Decode(&result); err != nil {
		t.Fatal(err)
	}

	b, _ := codec.Marshal(&raw)

	var result2 string
	codec.Unmarshal(b, &result2)
}
