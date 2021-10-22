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

// ParseAddress converts a string into an Address.
func ParseAddress(addr string) (*Address, error) {
	if addrLen := len(addr); addrLen != 63 {
		return nil, fmt.Errorf("invalid address length : got %d", addrLen)
	}

	if prefix := strings.ToLower(addr)[:4]; prefix != addrPrefix {
		return nil, errors.New("invalid prefix")
	}

	_, data, _ := bech32.Decode(addr)

	if len(data) == 0 {
		return nil, errors.New("empty data")
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
