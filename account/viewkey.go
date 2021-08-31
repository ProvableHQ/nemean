package account

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
)

var viewKeyPrefix = []byte{14, 138, 223, 204, 247, 224, 122}

type ViewKey struct {
	DecryptionKey []byte
}

func ParseViewKey(key string) (*ViewKey, error){
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