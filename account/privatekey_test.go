package account

import (
	"bytes"
	"testing"
)

func TestParsePrivateKey(t *testing.T) {
	k := "APrivateKey1uaf51GJ6LuMzLi2jy9zJJC3doAtngx52WGFZrcvK6aBsEgo"
	res, err := ParsePrivateKey(k)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(res.Seed, []byte{67, 224, 203, 43, 72, 114, 48, 110, 151, 178, 153, 154, 154, 200, 230, 117, 54, 196, 93, 219, 200, 5, 118, 56, 106, 80, 251, 4, 250, 233, 95, 40}) {
		t.Fatal("invalid seed")
	}

	if res.RPkCounter != 4398 {
		t.Fatal("invalid counter")
	}
}