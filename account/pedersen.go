package account

import (
	"errors"
	"fmt"
	"github.com/pinestreetlabs/aleo-wallet-sdk/internal/crypto/twistededwards"
	"github.com/pinestreetlabs/aleo-wallet-sdk/internal/utils"
	"go.dedis.ch/kyber/v3/group/mod"
	"math/big"
)

// Pedersen.go contains the utilities for creating pedersen commitments.

const (
	PedersenWindowSize int = 192
	PedersenNumWindows int = 8
)

var errPedersenInputLen = errors.New("input for pedersen commitment invalid")

func PedersenCommitment(input []byte, bases [][]*twistededwards.ExtPoint, randomBases []*twistededwards.ExtPoint, random []byte) ([]byte, error) {
	hash, err := PedersenHash(input, bases)
	if err != nil {
		return nil, err
	}

	// h^r
	random = new(big.Int).SetBytes(random).Bytes()

	for i, base := range randomBases {
		if new(big.Int).SetBytes(random).Bit(i) > 0 {
			hash.Add(hash, base)
		}
	}

	// x-coordinate
	res := new(mod.Int)
	res.Mul(&hash.X, res.Inv(&hash.Z))
	buf, _ := res.MarshalBinary()

	return utils.ConvertByteOrder(buf), nil
}

func PedersenHash(input []byte, bases [][]*twistededwards.ExtPoint) (*twistededwards.ExtPoint, error) {
	// Check the input is within the available length.
	if maxLen := PedersenWindowSize * PedersenNumWindows; len(input) > maxLen {
		return nil, fmt.Errorf("%w : got %d want less than %d", errPedersenInputLen, len(input), maxLen)
	}

	// Pad the input to windowSize * numWindows.
	// Add trailing zero bytes.
	in := make([]byte, PedersenWindowSize)
	copy(in, input)

	bits := toBits(in)

	curve := twistededwards.ExtendedCurve{}

	res := make([]*twistededwards.ExtPoint, 0)

	// Turn the slice of bits into chunks of 8.
	chunks := intoChunks(bits, len(bits)/8)

	// For each chunk of 8 chunks.
	for i, chunk := range chunks {
		a := twistededwards.NewPoint(twistededwards.Zero, twistededwards.One, twistededwards.One, twistededwards.Zero,
			curve.Init(twistededwards.ParamBLS12377(), true))
		for j, bit := range chunk {
			if bit {
				b := bases[i][j]
				a.Add(a, b)
			}
		}
		res = append(res, a)
	}

	a := twistededwards.NewPoint(twistededwards.Zero, twistededwards.One, twistededwards.One, twistededwards.Zero,
		curve.Init(twistededwards.ParamBLS12377(), true))
	for _, b := range res {
		a.Add(a, b)
	}

	return a, nil
}

func toBits(bytes []byte) []bool {
	res := make([]bool, 0)
	for _, item := range bytes {
		for i := 0; i < 8; i++ {
			if (item>>i)&1 > 0 {
				res = append(res, true)
			} else {
				res = append(res, false)
			}
		}
	}
	return res
}

// intoChunks accepts a slice and returns a slice of the original into chunkSize slices.
func intoChunks(slice []bool, chunkSize int) (chunks [][]bool) {
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}
