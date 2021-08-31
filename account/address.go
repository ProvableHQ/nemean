package account

import (
	"errors"
	"fmt"
	"github.com/btcsuite/btcutil/bech32"
	"strings"
)

var addrPrefix = "aleo"

type Address struct {
	EncryptionKey []byte
}

func ParseAddress(addr string)  (*Address, error){
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
