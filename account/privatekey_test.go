package account

import (
	"bytes"
	"github.com/pinestreetlabs/aleo-wallet-sdk/parameters"
	"testing"
)

func TestParsePrivateKey(t *testing.T) {
	// TODO change
	k := "APrivateKey1uaf51GJ6LuMzLi2jy9zJJC3doAtngx52WGFZrcvK6aBsEgo"
	res, err := ParsePrivateKey(k)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(res.Seed, []byte{67, 224, 203, 43, 72, 114, 48, 110, 151, 178, 153, 154, 154, 200, 230, 117, 54, 196, 93, 219, 200, 5, 118, 56, 106, 80, 251, 4, 250, 233, 95, 40}) {
		t.Fatal("invalid seed")
	}

	if res.RPkCounter != 4398 {
		t.Fatal("invalid cAPrivateKey1xpBtAQmv5sHHWwZqya9UVbVBcGtnX95TAN7XSAQ6yLqE5bCounter")
	}

	if k != res.String() {
		t.Fatal("invalid stringer")
	}
}

func TestParsePrivateKey2(t *testing.T) {
	privKey := "APrivateKey1zkp8cC4jgHEBnbtu3xxs1Ndja2EMizcvTRDq5Nikdkukg1p"

	key, err := ParsePrivateKey(privKey)
	if err != nil {
		t.Fatal(err)
	}

	// Create view key.
	if key.ViewKey().String() != "AViewKey1iAf6a7fv6ELA4ECwAth1hDNUJJNNoWNThmREjpybqder" {
		t.Fatalf("want %s got %s", "AViewKey1iAf6a7fv6ELA4ECwAth1hDNUJJNNoWNThmREjpybqder", key.ViewKey().String())
	}
}

func TestFromSeed(t *testing.T) {
	// Test case from dpc/src/account/tests.rs
	seed := [32]byte{225, 188, 136, 113, 36, 134, 74, 147, 46, 205, 27, 245, 37, 173, 115, 101, 220, 243, 27, 56, 238, 226, 66, 152, 152, 245, 198, 104, 39, 128, 69, 183}
	params, err := parameters.Load("../params/")
	if err != nil {
		t.Fatal(err)
	}

	sk, err := FromSeed(seed, params.AccountSignature, params.AccountCommitment)
	if err != nil {
		t.Fatal(err)
	}

	if expected := "APrivateKey1xpBtAQmv5sHHWwZqya9UVbVBcGtnX95TAN7XSAQ6yLqE5bC"; sk.String() != expected {
		t.Fatalf("%s != %s", sk.String(), expected)
	}
}
