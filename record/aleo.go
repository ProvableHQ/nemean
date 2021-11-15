package record

/*
#cgo LDFLAGS: -L../aleo -laleo
#include "../aleo/aleo.h"
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
	res := C.new_input_record(C.CString(address.String()), C.uint64_t(value), (*C.uint8_t)(unsafe.Pointer(&payload[0])), (*C.uint8_t)(unsafe.Pointer(&randomness[0])), C.size_t(len(randomness)))
	if res == nil {
		return nil, handleCError()
	}

	serialNumberNonce := C.record_serial_number_nonce((*C.record_t)(unsafe.Pointer(res)))

	programID := C.record_program_id((*C.record_t)(unsafe.Pointer(res)))

	commitmentRandomness := C.record_commitment_randomness((*C.record_t)(unsafe.Pointer(res)))

	return &Record{
		owner:                address.Copy(),
		value:                value,
		payload:              payload[:],
		programID:            C.GoString(programID),
		serialNumberNonce:    C.GoString(serialNumberNonce),
		commitmentRandomness: C.GoString(commitmentRandomness),
	}, nil
}

//func fromRecord(owner *account.Address, value int64, payload [128]byte, )

func encryptRecord(record *Record, randomness []byte) (string, error) {
	// fromrecord
	res := C.from_record(C.CString(record.owner.String()), C.uint64_t(record.value), (*C.uint8_t)(unsafe.Pointer(&record.payload[0])), C.CString(record.serialNumberNonce), C.CString(record.commitmentRandomness))
	if res == nil {
		return "", handleCError()
	}

	cipher := C.encrypt_record((*C.record_t)(unsafe.Pointer(res)), (*C.uint8_t)(unsafe.Pointer(&randomness[0])), C.size_t(len(randomness)))
	fmt.Println(C.GoString(cipher))

	return C.GoString(cipher), nil
}

func decryptRecord(ciphertext string, viewKey *account.ViewKey) (*Record, error) {
	res := C.decrypt_record(C.CString(ciphertext), C.CString(viewKey.String()))
	if res == nil {
		return nil, handleCError()
	}
	owner := C.record_owner((*C.record_t)(unsafe.Pointer(res)))

	addr, err := account.ParseAddress(C.GoString(owner))
	if err != nil {
		return nil, err
	}
	value := C.record_value((*C.record_t)(unsafe.Pointer(res)))
	payload := C.record_payload((*C.record_t)(unsafe.Pointer(res)))
	buf := *(*[]byte)(unsafe.Pointer(&payload.data))

	programID := C.record_program_id((*C.record_t)(unsafe.Pointer(res)))
	serialNumberNonce := C.record_serial_number_nonce((*C.record_t)(unsafe.Pointer(res)))
	commitmentRandomness := C.record_commitment_randomness((*C.record_t)(unsafe.Pointer(res)))

	return &Record{
		owner:                addr,
		value:                int64(value),
		payload:              buf,
		programID:            C.GoString(programID),
		serialNumberNonce:    C.GoString(serialNumberNonce),
		commitmentRandomness: C.GoString(commitmentRandomness),
	}, nil
}
