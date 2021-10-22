package account

import (
	"github.com/consensys/gnark-crypto/ecc/bls12-377/fr"
	"github.com/consensys/gnark-crypto/ecc/bls12-377/twistededwards"
	"github.com/pinestreetlabs/aleo-wallet-sdk/internal/utils"
	"github.com/pinestreetlabs/aleo-wallet-sdk/parameters"
	"math/big"
)

// generateSchnorrPublicKey creates a signature public key used for deriving the account view key.
func generateSchnorrPublicKey(key []byte, sigParams *parameters.AccountSignature) twistededwards.PointAffine {
	z := fr.Element{}
	z.SetOne()
	x := fr.Element{}
	x.SetZero()

	pubKeySig := twistededwards.NewPointAffine(x, z)

	key = utils.ConvertByteOrder(key)

	privateKey := new(big.Int).SetBytes(key)

	// Iterate over each bit in the private key.
	for i := 0; i < 256; i++ {
		if privateKey.Bit(i) > 0 {
			pubKeySig.Add(&pubKeySig, &sigParams.GeneratorPowers[i])
		}
	}

	return pubKeySig
}
