package account

/*
crypto.go contains the bindings to the underlying snarkvm aleo package.
To avoid upstream changes and re-implementing the snarkvm-curves crate in Go, we use Rust FFI.
*/

/*
#cgo LDFLAGS: -L../aleo -laleo
#include "../aleo/aleo.h"
#include <stdlib.h>
#include <stdio.h>
*/
import "C"
import (
	"unsafe"
)

func fromSeed(seed [32]byte) (*Account, error) {
	res := C.from_seed((*C.uint8_t)(unsafe.Pointer(&seed[0])), C.size_t(32))
	privateKeyC := C.account_private_key((*C.account_t)(unsafe.Pointer(res)))
	// todo free
	privateKey, err := ParsePrivateKey(C.GoString(privateKeyC))
	if err != nil {
		return nil, err
	}

	viewKeyC := C.account_view_key((*C.account_t)(unsafe.Pointer(res)))
	viewKey, err := ParseViewKey(C.GoString(viewKeyC))
	if err != nil {
		return nil, err
	}

	addressC := C.account_address((*C.account_t)(unsafe.Pointer(res)))
	address, err := ParseAddress(C.GoString(addressC))
	if err != nil {
		return nil, err
	}

	return &Account{
		privateKey: privateKey,
		viewKey:    viewKey,
		address:    address,
	}, nil
}
