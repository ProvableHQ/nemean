package record

/*
#cgo LDFLAGS: -L/usr/lib -laleo
#include <aleo.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
*/
import "C"
import (
	"fmt"
	"github.com/pinestreetlabs/aleo-wallet-sdk/account"
	"unsafe"
)

func handleCError() error {
	errLen := C.last_error_length()
	errMsg := C.CString("")
	C.last_error_message(errMsg, errLen)
	return fmt.Errorf("aleo : %v", C.GoString(errMsg))
}

// create a record
func newInputRecord(address *account.Address, value int64, payload [128]byte, randomness []byte) (*Record, error) {
	res := C.new_input_record(C.CString(address.String()), C.int64_t(value), (*C.uint8_t)(unsafe.Pointer(&payload[0])), (*C.uint8_t)(unsafe.Pointer(&randomness[0])), C.size_t(len(randomness)))
	if res == nil {
		return nil, handleCError()
	}
	programID := C.record_program_id((*C.record_t)(unsafe.Pointer(res)))

	commitmentRandomness := C.record_commitment_randomness((*C.record_t)(unsafe.Pointer(res)))

	return &Record{
		owner:                address.Copy(),
		value:                value,
		payload:              payload[:],
		programID:            C.GoString(programID),
		commitmentRandomness: C.GoString(commitmentRandomness),
	}, nil
}

func encryptRecord(record *Record) (string, error) {
	res := C.from_record(C.CString(record.owner.String()), C.int64_t(record.value), (*C.uint8_t)(unsafe.Pointer(&record.payload[0])))
	if res == nil {
		return "", handleCError()
	}

	defer C.record_free((*C.account_t)(unsafe.Pointer(res)))

	cipher := C.encrypt_record((*C.record_t)(unsafe.Pointer(res)))
	defer C.free(unsafe.Pointer(cipher))

	return C.GoString(cipher), nil
}

func decryptRecord(ciphertext string, viewKey *account.ViewKey) (*Record, error) {
	res := C.decrypt_record(C.CString(ciphertext), C.CString(viewKey.String()))
	if res == nil {
		return nil, handleCError()
	}
	defer C.record_free((*C.account_t)(unsafe.Pointer(res)))

	owner := C.record_owner((*C.record_t)(unsafe.Pointer(res)))
	defer C.free(unsafe.Pointer(owner))

	addr, err := account.ParseAddress(C.GoString(owner))
	if err != nil {
		return nil, err
	}
	value := C.record_value((*C.record_t)(unsafe.Pointer(res)))
	payload := C.record_payload((*C.record_t)(unsafe.Pointer(res)))
	buf := *(*[]byte)(unsafe.Pointer(&payload.data))

	programID := C.record_program_id((*C.record_t)(unsafe.Pointer(res)))
	commitmentRandomness := C.record_commitment_randomness((*C.record_t)(unsafe.Pointer(res)))
	defer C.free(unsafe.Pointer(programID))
	defer C.free(unsafe.Pointer(commitmentRandomness))

	return &Record{
		owner:                addr,
		value:                int64(value),
		payload:              buf,
		programID:            C.GoString(programID),
		commitmentRandomness: C.GoString(commitmentRandomness),
	}, nil
}
