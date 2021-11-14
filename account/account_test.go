package account

import (
	"github.com/pinestreetlabs/aleo-wallet-sdk/network"
	"testing"
)

func TestFromPrivateKey(t *testing.T) {
	expected := struct {
		sk   string
		vk   string
		addr string
	}{
		sk:   "APrivateKey1zkp8cC4jgHEBnbtu3xxs1Ndja2EMizcvTRDq5Nikdkukg1p",
		vk:   "AViewKey1iAf6a7fv6ELA4ECwAth1hDNUJJNNoWNThmREjpybqder",
		addr: "aleo1d5hg2z3ma00382pngntdp68e74zv54jdxy249qhaujhks9c72yrs33ddah",
	}
	acc, err := FromPrivateKey(expected.sk, network.Testnet2())
	if err != nil {
		t.Fatal(err)
	}

	if sk := acc.PrivateKey().String(); sk != expected.sk {
		t.Fatalf("got %s want %s", sk, expected.sk)
	}

	if sk := acc.ViewKey().String(); sk != expected.vk {
		t.Fatalf("got %s want %s", sk, expected.vk)
	}

	if sk := acc.Address().String(); sk != expected.addr {
		t.Fatalf("got %s want %s", sk, expected.addr)
	}
}

func TestFromPrivateKeyInvalid(t *testing.T) {
	if _, err := FromPrivateKey("APrivateKey1", network.Testnet2()); err == nil {
		t.Fatal("expected err")
	}

	if _, err := FromPrivateKey("APrivateKey1abcdefghijklmnopqrstuvwxyz", network.Testnet2()); err == nil {
		t.Fatal("expected err")
	}

	if _, err := FromPrivateKey("", network.Testnet2()); err == nil {
		t.Fatal("expected err")
	}
}
