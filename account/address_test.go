package account

import "testing"

func TestParseAddress(t *testing.T) {
	addr := "aleo1d5hg2z3ma00382pngntdp68e74zv54jdxy249qhaujhks9c72yrs33ddah"
	res, err := ParseAddress(addr)
	if err != nil {
		t.Fatal(err)
	}

	if addr != res.String() {
		t.Fatalf("invalid stringer : got %s want %s", res.String(), addr)
	}
}

func TestParseAddressInvalid(t *testing.T) {
	if _, err := ParseAddress(""); err == nil {
		t.Fatal(err)
	}

	if _, err := ParseAddress("aleo1abcdefghijklmnopqrstuvwxyz"); err == nil {
		t.Fatal(err)
	}

	if _, err := ParseAddress("aleo1"); err == nil {
		t.Fatal(err)
	}
}
