package account

import (
	"bytes"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
)

// privateKeyPrefix is the human-readable prefix for private keys.
var privateKeyPrefix = []byte{127, 134, 189, 116, 210, 221, 210, 137, 145, 18, 253}

// PrivateKey contains the seed and relevant keys to manipulate an Aleo account.
type PrivateKey struct {
	Seed [32]byte
}

var errInvalidPKLen = errors.New("invalid private key length")

// ParsePrivateKey accepts a private key string and returns the PrivateKey.
func ParsePrivateKey(key string) (*PrivateKey, error) {
	// An account private key is formatted as a Base58 string, comprised of 58 characters.
	buf := base58.Decode(key)

	if keyLen := len(buf); keyLen != 43 {
		return nil, fmt.Errorf("ParsePrivateKey : %w: got %d", errInvalidPKLen, keyLen)
	}

	if !bytes.Equal(buf[0:11], privateKeyPrefix) {
		return nil, fmt.Errorf("ParsePrivateKey : %w : got %v want %v", errInvalidPrefix, buf[0:9], privateKeyPrefix)
	}

	// Last 32 bytes are the seed.
	var seed [32]byte
	copy(seed[:], buf[11:43])

	return &PrivateKey{Seed: seed}, nil
}

// NewSeed creates a uniformly random 32-byte account seed.
func NewSeed() ([32]byte, error) {
	var d [32]byte
	if _, err := rand.Read(d[:]); err != nil {
		return [32]byte{}, err
	}
	return d, nil
}

// String implements the stringer interface for PrivateKey.
// Returns the base58 encoded string.
func (pk PrivateKey) String() string {
	var buf bytes.Buffer
	buf.Write(privateKeyPrefix)
	buf.Write(pk.Seed[:])
	return base58.Encode(buf.Bytes())
}

// Copy does a deep copy on PrivateKey.
func (pk PrivateKey) Copy() *PrivateKey {
	newPrivateKey := &PrivateKey{Seed: [32]byte{}}
	copy(newPrivateKey.Seed[:], pk.Seed[:])
	return newPrivateKey
}
