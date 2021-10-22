package account

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/pinestreetlabs/aleo-wallet-sdk/internal/utils"
	"github.com/pinestreetlabs/aleo-wallet-sdk/parameters"
	"golang.org/x/crypto/blake2s"
	"math"
	"math/big"
)

// initialSkPrf is the initial slice of bytes used to initialize the sk_prf in Blake2s evaluation function.
var initialSkPrf = []byte{
	0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
}

// privateKeyPrefix is the human-readable prefix for private keys.
var privateKeyPrefix = []byte{127, 134, 189, 116, 210, 221, 210, 137, 144}

var errInvalidSeed = errors.New("invalid seed")

// PrivateKey contains the seed and relevant keys to manipulate an Aleo account.
type PrivateKey struct {
	Seed       []byte
	RPkCounter uint16
	SkSig      []byte
	SkPrf      []byte
}

// String implements the stringer interface for PrivateKey.
// Returns the base58 encoded string.
func (pk PrivateKey) String() string {
	var buf bytes.Buffer
	buf.Write(privateKeyPrefix)
	_ = binary.Write(&buf, binary.LittleEndian, pk.RPkCounter)
	buf.Write(pk.Seed)
	return base58.Encode(buf.Bytes())
}

// ParsePrivateKey accepts a private key string and returns the PrivateKey.
func ParsePrivateKey(key string) (*PrivateKey, error) {
	// An account private key is formatted as a Base58 string, comprised of 58 characters.
	buf := base58.Decode(key)

	if keyLen := len(buf); keyLen != 43 {
		return nil, fmt.Errorf("invalid key length : got %d", keyLen)
	}

	if !bytes.Equal(buf[0:9], privateKeyPrefix) {
		return nil, errors.New("invalid prefix")
	}

	// First 2 bytes are the counter.
	// Last 32 bytes are the seed.
	counter, seed := buf[9:11], buf[11:]

	return &PrivateKey{
		Seed:       seed,
		RPkCounter: binary.LittleEndian.Uint16(counter),
	}, nil
}

// NewSeed creates a uniformly random 32-byte account seed.
func NewSeed() ([]byte, error) {
	d := make([]byte, 32)
	if _, err := rand.Read(d); err != nil {
		return nil, err
	}
	return d, nil
}

//FromSeed creates a valid account private key from a provided seed.
// Ths seed must be 32 bytes.
func FromSeed(seed [32]byte, sigParams *parameters.AccountSignature, commitParams *parameters.AccountCommitment) (*PrivateKey, error) {
	// Construct private key components.
	// 1. sk_sig = Blake2s(seed, 0)
	// 2. sk_prf = Blake2s(seed, 1)
	// 3. r_pk = Blake2s(seed, counter)

	// 1. Create the secret key signature.
	skSig := blake2s.Sum256(append(seed[:], bytes.Repeat([]byte{0x00}, 32)...))
	// 2. Create the secret key prf.
	skPrf := blake2s.Sum256(append(seed[:], initialSkPrf...))

	// A counter is an u16 value that is iterated on until a valid view key
	// can be derived from private_key.
	var counter uint16 = 2

	// Loop until we find a valid view key, or until we reach the maximum counter.
	for {
		if counter > math.MaxUint16 {
			return nil, errInvalidSeed
		}

		var buf bytes.Buffer

		// Write the seed to the buffer.
		if err := binary.Write(&buf, binary.LittleEndian, seed[:]); err != nil {
			return nil, err
		}

		// Write the counter to the buffer.
		counterBuf := make([]byte, 32)
		binary.LittleEndian.PutUint16(counterBuf, counter)

		if err := binary.Write(&buf, binary.LittleEndian, counterBuf); err != nil {
			return nil, err
		}

		// 3. Create a random private key.
		out := blake2s.Sum256(buf.Bytes())

		// Convert to little-endian.
		rPK := utils.ConvertByteOrder(out[:])

		// Check that the coordinate is on the curve (mod scalar field on Edwards BLS12).
		if isOnCurve(rPK) {
			// Check if the private key can create a valid view key.
			ok, err := ValidPrivateKey(skSig[:], skPrf[:], sigParams, commitParams, rPK[:])
			if ok {
				break
			}

			if err != nil {
				return nil, err
			}
		}

		counter += 1
	}

	return &PrivateKey{
		Seed:       seed[:],
		RPkCounter: counter,
		SkSig:      skSig[:],
		SkPrf:      skPrf[:],
	}, nil
}

// isOnCurve checks that the coordinate is on the curve (mod scalar field on Edwards BLS12).
// TODO, move this into internal/crypto and use params.
func isOnCurve(x []byte) bool {
	mod, _ := hex.DecodeString("04AAD957A68B2955982D1347970DEC005293A3AFC43C8AFEB95AEE9AC33FD9FF")
	modulo := new(big.Int).SetBytes(mod)
	if new(big.Int).SetBytes(x[:]).Cmp(modulo) == -1 {
		return true
	}
	return false
}

// ValidPrivateKey checks if the private key can be used to generate a valid view key.
func ValidPrivateKey(key []byte, skPrf []byte, sigParams *parameters.AccountSignature, commitParams *parameters.AccountCommitment, rPK []byte) (bool, error) {
	// Generate a Schnorr public key with the provided key and account signature params.
	pkSig := generateSchnorrPublicKey(key, sigParams)

	// Create an input to be used to generate a pedersen commitment.
	var commitmentInput bytes.Buffer
	{
		x := pkSig.X.Bytes()
		if err := binary.Write(&commitmentInput, binary.LittleEndian, utils.ConvertByteOrder(x[:])); err != nil {
			return false, err
		}

		y := pkSig.Y.Bytes()
		if err := binary.Write(&commitmentInput, binary.LittleEndian, utils.ConvertByteOrder(y[:])); err != nil {
			return false, err
		}

		commitmentInput.Write(skPrf)
	}

	commitment, err := PedersenCommitment(commitmentInput.Bytes(), commitParams.Bases, commitParams.RandomBase, rPK)
	if err != nil {
		return false, err
	}

	// Check if the affine x-coordinate is on the curve.
	if !isOnCurve(commitment) {
		return false, nil
	}

	// MSB check
	accDecKey := toBits(commitment)
	if keyLen := len(accDecKey); keyLen == 0 || keyLen < 251 {
		return false, nil
	}

	for _, bit := range accDecKey[249:] {
		if bit {
			return false, nil
		}
	}

	return true, nil
}
