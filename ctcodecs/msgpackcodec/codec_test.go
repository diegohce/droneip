package msgpackcodec

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

	codec := &msgpackCodec{}

	var buf bytes.Buffer

	if err := codec.NewEncoder(&buf).Encode(&someGuy); err != nil {
		t.Fatal(err)
	}

	var anotherGuy Guy

	if err := codec.NewDecoder(&buf).Decode(&anotherGuy); err != nil {
		t.Fatal(err)
	}

	b, _ := codec.Marshal(&someGuy)

	codec.Unmarshal(b, &anotherGuy)
}
