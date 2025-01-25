package ctcodecs_test

import (
	"bytes"
	"testing"

	"github.com/diegohce/droneip/ctcodecs"
	_ "github.com/diegohce/droneip/ctcodecs/gobcodec"
	_ "github.com/diegohce/droneip/ctcodecs/jsoncodec"
	_ "github.com/diegohce/droneip/ctcodecs/msgpackcodec"
	_ "github.com/diegohce/droneip/ctcodecs/xmlcodec"
)

type Guy struct {
	Name string
	Age  int
	Pets []string
}

var someGuy Guy = Guy{
	Name: "Diego",
	Age:  45,
	Pets: []string{"Simona", "Macacha"},
}

func TestCodecs(t *testing.T) {

	contentTypes := []string{
		"application/json",
		"application/msgpack",
		"application/gob",
		"application/xml",
	}

	for _, ct := range contentTypes {
		codec, _ := ctcodecs.New(ct)

		var buf bytes.Buffer

		if err := codec.NewEncoder(&buf).Encode(&someGuy); err != nil {
			t.Error(err)
		}

		var anotherGuy Guy

		if err := codec.NewDecoder(&buf).Decode(&anotherGuy); err != nil {
			t.Error(err)
		}

		if codec.MimeType() != ct {
			t.Errorf("mime type: got %s want %s", codec.MimeType(), ct)
		}
	}
}

func TestUnregisteredContentType(t *testing.T) {

	_, err := ctcodecs.New("Baby-Shark")
	if err == nil {
		t.Error("expecting error for UnregisteredContentType")
	}
}

func TestCodecsList(t *testing.T) {
	l := ctcodecs.List()
	if len(l) <= 0 {
		t.Error("empty codecs list")
	}
}
