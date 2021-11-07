package account

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
)

var viewKeyPrefix = []byte{14, 138, 223, 204, 247, 224, 122}

// ViewKey is an Aleo view key.
type ViewKey struct {
	DecryptionKey []byte
}

// ParseViewKey parses a string encoded ViewKey.
func ParseViewKey(key string) (*ViewKey, error) {
	buf := base58.Decode(key)

	if keyLen := len(buf); keyLen != 39 {
		return nil, fmt.Errorf("invalid key length : got %d", keyLen)
	}

	if !bytes.Equal(buf[0:7], viewKeyPrefix) {
		return nil, errors.New("invalid prefix")
	}

	decryptionKey := buf[7:]

	return &ViewKey{DecryptionKey: decryptionKey}, nil
}

// String implements the stringer interface for ViewKey.
// Returns the base58 encoded string.
func (vk ViewKey) String() string {
	var buf bytes.Buffer
	buf.Write(viewKeyPrefix)
	binary.Write(&buf, binary.LittleEndian, vk.DecryptionKey)
	return base58.Encode(buf.Bytes())
}

// Copy does a deep copy on ViewKey.
func (vk ViewKey) Copy() *ViewKey {
	newViewKey := &ViewKey{DecryptionKey: make([]byte, len(vk.DecryptionKey))}
	copy(newViewKey.DecryptionKey[:], vk.DecryptionKey[:])
	return newViewKey
}
