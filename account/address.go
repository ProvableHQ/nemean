package account

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/btcsuite/btcutil/bech32"
	"strings"
)

// addrPrefix is the human-readable prefix for each aleo address.
var addrPrefix = "aleo"

// Address contains the encryption key which is used to generate a human readable address.
type Address struct {
	EncryptionKey []byte
}

var errInvalidAddrLen = errors.New("invalid address length")
var errInvalidPrefix = errors.New("invalid hrp")
var errAddrBech32m = errors.New("address is not encoded in bech32m")
var errInvalidAddrData = errors.New("missing data")

// ParseAddress converts a string into an Address.
func ParseAddress(addr string) (*Address, error) {
	if addrLen := len(addr); addrLen != 63 {
		return nil, fmt.Errorf("ParseAddress : %w : got %d", errInvalidAddrLen, addrLen)
	}

	if prefix := strings.ToLower(addr)[:4]; prefix != addrPrefix {
		return nil, fmt.Errorf("ParseAddress : %w", errInvalidPrefix)
	}

	_, data, version, err := bech32.DecodeGeneric(addr)
	if err != nil {
		return nil, fmt.Errorf("ParseAddress : %s", err)
	}

	if version == bech32.Version0 {
		return nil, fmt.Errorf("ParseAddress : %w", errAddrBech32m)
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("ParseAddress : %w", errInvalidAddrData)
	}

	return &Address{EncryptionKey: data}, nil
}

// String implements the stringer interface for Address.
// Returns a bech32 encoded string.
// If unable to encode to bech32, an empty string is returned.
func (a Address) String() string {
	var buf bytes.Buffer
	_ = binary.Write(&buf, binary.LittleEndian, a.EncryptionKey)

	addr, _ := bech32.Encode(addrPrefix, buf.Bytes())

	return addr
}

// Copy does a deep copy on Address.
func (a Address) Copy() *Address {
	newAddr := &Address{EncryptionKey: make([]byte, len(a.EncryptionKey))}
	copy(newAddr.EncryptionKey[:], a.EncryptionKey[:])
	return newAddr
}
