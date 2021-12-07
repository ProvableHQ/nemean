package transaction

/*
#cgo LDFLAGS: -L/usr/lib -laleo
#include <aleo.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
*/
import "C"
import (
	"errors"
	"fmt"
	"github.com/pinestreetlabs/aleo-wallet-sdk/account"
	"github.com/pinestreetlabs/aleo-wallet-sdk/record"
	"unsafe"
)

func handleCError() error {
	errLen := C.last_error_length()
	errMsg := C.CString("")
	C.last_error_message(errMsg, errLen)
	return fmt.Errorf("aleo : %v", C.GoString(errMsg))
}

func newCoinbaseTransaction(address *account.Address, value int64, random []byte) string {
	res := C.new_coinbase_transaction(C.CString(address.String()), C.int64_t(value), (*C.uint8_t)(unsafe.Pointer(&random[0])), C.size_t(len(random)))
	defer C.free(unsafe.Pointer(res))
	return C.GoString(res)
}

func newTransferTransaction(privateKey *account.PrivateKey, to *account.Address, in *record.Record, ledgerProofs []string, amount, fee int64) (string, error) {
	inRecord := C.from_record(C.CString(in.Owner().String()), C.int64_t(in.Value()), (*C.uint8_t)(unsafe.Pointer(&in.Payload()[0])))
	if inRecord == nil {
		return "", handleCError()
	}

	defer C.free(unsafe.Pointer(inRecord))

	if len(ledgerProofs) != 2 {
		return "", errors.New("wrong number of ledger proofs")
	}

	txn := C.new_transfer_transaction(inRecord, C.CString(ledgerProofs[0]), C.CString(ledgerProofs[1]), C.CString(privateKey.String()), C.int64_t(amount), C.int64_t(fee), C.CString(to.String()))
	if txn == nil {
		return "", handleCError()
	}

	defer C.free(unsafe.Pointer(txn))

	return C.GoString(txn), nil
}
