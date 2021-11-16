package record

import (
	"fmt"
	"github.com/pinestreetlabs/aleo-wallet-sdk/account"
)

type Record struct {
	owner                *account.Address
	value                int64
	payload              []byte
	programID            string
	serialNumberNonce    string
	commitmentRandomness string
}

func (r Record) Owner() *account.Address {
	return r.owner.Copy()
}

func (r Record) Value() int64 {
	return r.value
}

func (r Record) Payload() []byte {
	return r.payload
}

func (r Record) ProgramID() string {
	return r.programID
}

func (r Record) SerialNumberNonce() string {
	return r.serialNumberNonce
}

func (r Record) CommitmentRandomness() string {
	return r.commitmentRandomness
}

func (r Record) String() string {
	return fmt.Sprintf("owner: %v\nvalue: %v\n,payload: %v\n, programID: %s\nserialNumberNonce: %s\ncommitmentRandomness: %s\n",
		r.owner, r.value, r.payload, r.programID, r.serialNumberNonce, r.commitmentRandomness)
}

func NewInputRecord(address *account.Address, value int64, payload [128]byte, randomness []byte) (*Record, error) {
	return newInputRecord(address, value, payload, randomness)
}

func EncryptRecord(record *Record, encrypt []byte) (string, error) {
	return encryptRecord(record, encrypt)
}

func DecryptRecord(ciphertext string, viewKey *account.ViewKey) (*Record, error) {
	return decryptRecord(ciphertext, viewKey)
}
