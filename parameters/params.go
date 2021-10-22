package parameters

import (
	blstwisted "github.com/consensys/gnark-crypto/ecc/bls12-377/twistededwards"
	"github.com/pinestreetlabs/aleo-wallet-sdk/internal/crypto/twistededwards"
)

// AccountCommitment is the global parameters for the Account Commitment.
// This is used for commitment to generate an account view key.
type AccountCommitment struct {
	Bases      [][]*twistededwards.ExtPoint
	RandomBase []*twistededwards.ExtPoint
	Crh        string
}

// AccountSignature is the global parameters for the Account Signature.
// This is used during public key generation.
type AccountSignature struct {
	// GeneratorPowers is a list of Points on the Twisted Edwards BLS12 curve.
	GeneratorPowers []blstwisted.PointAffine
	// Salt
	Salt []byte
}

type AccountEncryption struct {
	GeneratorPowers []*twistededwards.ExtPoint
}

type Parameters struct {
	AccountCommitment *AccountCommitment
	AccountSignature  *AccountSignature
	AccountEncryption *AccountEncryption
}
