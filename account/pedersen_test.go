package account

import (
	"github.com/consensys/gnark-crypto/ecc/bls12-377/fr"
	"github.com/pinestreetlabs/aleo-wallet-sdk/parameters"
	"testing"
)

func TestPedersenHash(t *testing.T) {
	params, _ := parameters.Load()

	input := []byte{66, 74, 11, 145, 100, 99, 130, 89, 163, 76, 43, 120, 39, 210, 45, 246, 90, 99, 196, 204, 248, 249, 81, 18, 98, 129, 227, 187, 223, 125, 77, 10, 51, 244, 232, 225, 249, 218, 217, 141, 51, 181, 187, 65, 96, 148, 73, 47, 129, 46, 71, 202, 250, 82, 135, 25, 147, 218, 97, 92, 115, 5, 43, 11, 144, 112, 228, 39, 76, 80, 243, 121, 42, 169, 114, 92, 122, 59, 231, 131, 252, 132, 245, 137, 35, 225, 97, 182, 221, 193, 210, 192, 49, 145, 61, 242}
	res, _ := PedersenHash(input, params.AccountCommitment.Bases)

	projX := fr.Element{}
	{
		buf, _ := res.X.MarshalBinary()
		projX.SetBytes(buf)
	}
	projY := fr.Element{}
	{
		buf, _ := res.Y.MarshalBinary()
		projY.SetBytes(buf)
	}
	projZ := fr.Element{}
	{
		buf, _ := res.Z.MarshalBinary()
		projZ.SetBytes(buf)
	}
	projT := fr.Element{}
	{
		buf, _ := res.T.MarshalBinary()
		projT.SetBytes(buf)
	}

	if expected := "5396547586132350007729949313861469404261305737086483362749900864719137061534"; projX.String() != expected {
		t.Fatalf("got %s want %s", projX.String(), expected)
	}

	if expected := "5347278242009712367012721659736647660641315059904831743920501275402848766032"; projY.String() != expected {
		t.Fatalf("got %s want %s", projY.String(), expected)
	}

	if expected := "1797456748704958900957750485032022246420453371611725673379215846573002912605"; projZ.String() != expected {
		t.Fatalf("got %s want %s", projZ.String(), expected)
	}

	if expected := "7137163236384982232555300942526908950902465003603796367373482163011780041459"; projT.String() != expected {
		t.Fatalf("got %s want %s", projT.String(), expected)
	}
}
