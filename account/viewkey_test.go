package account

import (
	"testing"
)

func TestParseViewKey(t *testing.T) {
	k := "AViewKey1m8gvywHKHKfUzZiLiLoHedcdHEjKwo5TWo6efz8gK7wF"
	res, err := ParseViewKey(k)
	if err != nil {
		t.Fatal(err)
	}

	if k != res.String() {
		t.Fatalf("invalid stringer : got %s want %s", res.String(), k)
	}
}

func TestParseViewKeyInvalid(t *testing.T) {
	if _, err := ParseViewKey(""); err == nil {
		t.Fatal(err)
	}

	if _, err := ParseViewKey("AViewKey1abcdefghijklmnopqrstuvwxyz"); err == nil {
		t.Fatal(err)
	}

	if _, err := ParseViewKey("AViewKey1"); err == nil {
		t.Fatal(err)
	}
}
