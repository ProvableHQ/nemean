package parameters

import (
	"bytes"
	"encoding/binary"
	"github.com/consensys/gnark-crypto/ecc/bls12-377/fr"
	"github.com/consensys/gnark-crypto/ecc/bls12-377/twistededwards"
	extendedTec "github.com/pinestreetlabs/aleo-wallet-sdk/internal/crypto/twistededwards"
	"io/ioutil"
	"math/big"
)

// loadAccountSignature reads the account signature parameters from file.
func loadAccountSignature(filePath string) (*AccountSignature, error) {
	// Read from file.
	body, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(body)

	// The generator powers are a list of Affine Coordinates (x,y).
	var generatorPowerLen uint32
	if err := binary.Read(r, binary.LittleEndian, &generatorPowerLen); err != nil {
		return nil, err
	}

	generators := make([]twistededwards.PointAffine, 0)

	var i uint32
	for i = 0; i < generatorPowerLen; i++ {
		x := make([]byte, 32)
		if err := binary.Read(r, binary.LittleEndian, &x); err != nil {
			return nil, err
		}

		y := make([]byte, 32)
		if err := binary.Read(r, binary.LittleEndian, &y); err != nil {
			return nil, err
		}

		x = convertOrder(x)
		y = convertOrder(y)
		xCoor := fr.Element{}
		xCoor.SetBytes(x)
		yCoor := fr.Element{}
		yCoor.SetBytes(y)

		point := twistededwards.PointAffine{
			X: xCoor,
			Y: yCoor,
		}

		generators = append(generators, point)
	}

	// Read salt.
	salt := make([]byte, 32)
	if err := binary.Read(r, binary.LittleEndian, &salt); err != nil {
		return nil, err
	}

	return &AccountSignature{GeneratorPowers: generators, Salt: salt}, nil
}

// convertOrder is a temporary helper function to convert byte order.
func convertOrder(buf []byte) []byte {
	for i := 0; i < len(buf)/2; i++ {
		buf[i], buf[len(buf)-i-1] = buf[len(buf)-i-1], buf[i]
	}
	return buf
}

// loadAccountCommitment reads account commitment from file.
func loadAccountCommitment(filePath string) (*AccountCommitment, error) {
	body, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	r := bytes.NewReader(body)

	var numBases uint32
	if err := binary.Read(r, binary.LittleEndian, &numBases); err != nil {
		return nil, err
	}

	bases := make([][]*extendedTec.ExtPoint, numBases)
	var i uint32
	for i = 0; i < numBases; i++ {
		var baseLen uint32
		if err := binary.Read(r, binary.LittleEndian, &baseLen); err != nil {
			return nil, err
		}

		points := make([]*extendedTec.ExtPoint, baseLen)

		var j uint32
		for j = 0; j < baseLen; j++ {
			x := make([]byte, 32)
			if err := binary.Read(r, binary.LittleEndian, &x); err != nil {
				panic(err)
			}

			y := make([]byte, 32)
			if err := binary.Read(r, binary.LittleEndian, &y); err != nil {
				panic(err)
			}

			t := make([]byte, 32)
			if err := binary.Read(r, binary.LittleEndian, &t); err != nil {
				panic(err)
			}

			z := make([]byte, 32)
			if err := binary.Read(r, binary.LittleEndian, &z); err != nil {
				panic(err)
			}

			x = convertOrder(x)
			y = convertOrder(y)
			z = convertOrder(z)
			t = convertOrder(t)

			xCoor := new(big.Int).SetBytes(x)
			yCoor := new(big.Int).SetBytes(y)
			zCoor := new(big.Int).SetBytes(z)
			tCoor := new(big.Int).SetBytes(t)

			curve := extendedTec.ExtendedCurve{}
			points[j] = extendedTec.NewPoint(xCoor, yCoor, zCoor, tCoor, curve.Init(extendedTec.ParamBLS12377(), true))
		}

		bases[i] = points
	}

	var numRandomBase uint32
	if err := binary.Read(r, binary.LittleEndian, &numRandomBase); err != nil {
		return nil, err
	}

	randomBases := make([]*extendedTec.ExtPoint, numRandomBase)

	for i = 0; i < numRandomBase; i++ {
		x := make([]byte, 32)
		if err := binary.Read(r, binary.LittleEndian, &x); err != nil {
			panic(err)
		}

		y := make([]byte, 32)
		if err := binary.Read(r, binary.LittleEndian, &y); err != nil {
			panic(err)
		}

		t := make([]byte, 32)
		if err := binary.Read(r, binary.LittleEndian, &t); err != nil {
			panic(err)
		}

		z := make([]byte, 32)
		if err := binary.Read(r, binary.LittleEndian, &z); err != nil {
			panic(err)
		}

		x = convertOrder(x)
		y = convertOrder(y)
		z = convertOrder(z)
		t = convertOrder(t)

		xCoor := new(big.Int).SetBytes(x)
		yCoor := new(big.Int).SetBytes(y)
		zCoor := new(big.Int).SetBytes(z)
		tCoor := new(big.Int).SetBytes(t)

		curve := extendedTec.ExtendedCurve{}
		randomBases[i] = extendedTec.NewPoint(xCoor, yCoor, zCoor, tCoor, curve.Init(extendedTec.ParamBLS12377(), true))
	}

	return &AccountCommitment{
		Bases:      bases,
		RandomBase: randomBases,
	}, nil
}

func loadAccountEncryption(filePath string) (*AccountEncryption, error) {
	body, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	r := bytes.NewReader(body)

	var numBases uint32
	if err := binary.Read(r, binary.LittleEndian, &numBases); err != nil {
		return nil, err
	}

	powers := make([]*extendedTec.ExtPoint, numBases)
	var i uint32
	for i = 0; i < numBases; i++ {
		x := make([]byte, 32)
		if err := binary.Read(r, binary.LittleEndian, &x); err != nil {
			panic(err)
		}

		y := make([]byte, 32)
		if err := binary.Read(r, binary.LittleEndian, &y); err != nil {
			panic(err)
		}

		t := make([]byte, 32)
		if err := binary.Read(r, binary.LittleEndian, &t); err != nil {
			panic(err)
		}

		z := make([]byte, 32)
		if err := binary.Read(r, binary.LittleEndian, &z); err != nil {
			panic(err)
		}

		x = convertOrder(x)
		y = convertOrder(y)
		z = convertOrder(z)
		t = convertOrder(t)

		xCoor := new(big.Int).SetBytes(x)
		yCoor := new(big.Int).SetBytes(y)
		zCoor := new(big.Int).SetBytes(z)
		tCoor := new(big.Int).SetBytes(t)

		curve := extendedTec.ExtendedCurve{}
		powers[i] = extendedTec.NewPoint(xCoor, yCoor, zCoor, tCoor, curve.Init(extendedTec.ParamBLS12377(), true))
	}

	return &AccountEncryption{
		GeneratorPowers: powers,
	}, nil
}

func Load() (*Parameters, error) {
	accountSignatureFile := "/Users/philipglazman/personal/aleo/snarkVM/parameters/src/global/account_signature.params"
	accountCommitmentFile := "/Users/philipglazman/personal/aleo/snarkVM/parameters/src/global/account_commitment.params"
	accountEncryptionFile := "/Users/philipglazman/personal/aleo/snarkVM/parameters/src/global/account_encryption.params"

	accountSignature, err := loadAccountSignature(accountSignatureFile)
	if err != nil {
		return nil, err
	}
	accountCommitment, err := loadAccountCommitment(accountCommitmentFile)
	if err != nil {
		return nil, err
	}

	accountEncryption, err := loadAccountEncryption(accountEncryptionFile)
	if err != nil {
		return nil, err
	}

	return &Parameters{
		AccountCommitment: accountCommitment,
		AccountSignature:  accountSignature,
		AccountEncryption: accountEncryption,
	}, nil
}
