package account

import "testing"

func TestParsePrivateKeyInvalid(t *testing.T) {
	if _, err := ParsePrivateKey(""); err == nil {
		t.Fatal(err)
	}

	if _, err := ParsePrivateKey("APrivateKey1abcdefghijklmnopqrstuvwxyz"); err == nil {
		t.Fatal(err)
	}

	if _, err := ParsePrivateKey("APrivateKey1"); err == nil {
		t.Fatal(err)
	}
}

func TestParsePrivateKey(t *testing.T) {
	key := "APrivateKey1zkp8cC4jgHEBnbtu3xxs1Ndja2EMizcvTRDq5Nikdkukg1p"
	res, err := ParsePrivateKey(key)
	if err != nil {
		t.Fatal(err)
	}

	if res.String() != key {
		t.Fatalf("got %s want %s", res.String(), key)
	}
}
