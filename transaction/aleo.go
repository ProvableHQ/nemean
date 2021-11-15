package transaction

/*
#cgo LDFLAGS: -L../aleo -laleo
#include "../aleo/aleo.h"
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
*/
import "C"
import (
	"github.com/pinestreetlabs/aleo-wallet-sdk/account"
	"unsafe"
)

func newCoinbaseTransaction(address *account.Address, value int64, random []byte) string {
	res := C.new_coinbase_transaction(C.CString(address.String()), C.uint64_t(value), (*C.uint8_t)(unsafe.Pointer(&random[0])), C.size_t(len(random)))
	return C.GoString(res)
}
