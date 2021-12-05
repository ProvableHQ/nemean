package record

import (
	"encoding/json"
	"fmt"
	"github.com/pinestreetlabs/aleo-wallet-sdk/account"
)

// Record is a fundamental data structure for encoding user assets and application state.
type Record struct {
	owner                *account.Address
	value                int64
	payload              []byte
	programID            string
	serialNumberNonce    string
	commitmentRandomness string
}

// JSON is a helper struct for JSON serialization.
type JSON struct {
	Owner                string `json:"owner"`
	Value                int64  `json:"value"`
	Payload              []byte `json:"payload"`
	ProgramID            string `json:"program_id"`
	SerialNumberNonce    string `json:"serial_number_nonce"`
	CommitmentRandomness string `json:"commitment_randomness"`
}

// MarshalJSON implements the marshaller interface.
func (r *Record) MarshalJSON() ([]byte, error) {
	return json.Marshal(JSON{
		Owner:                r.owner.String(),
		Value:                r.value,
		Payload:              r.payload,
		ProgramID:            r.programID,
		SerialNumberNonce:    r.serialNumberNonce,
		CommitmentRandomness: r.commitmentRandomness,
	})
}

// UnmarshalJSON implements the marshaller interface.
func (r *Record) UnmarshalJSON(b []byte) error {
	temp := &JSON{}

	if err := json.Unmarshal(b, &temp); err != nil {
		return err
	}

	addr, err := account.ParseAddress(temp.Owner)
	if err != nil {
		return err
	}

	r.owner = addr
	r.value = temp.Value
	r.payload = temp.Payload
	r.programID = temp.ProgramID
	r.serialNumberNonce = temp.SerialNumberNonce
	r.commitmentRandomness = temp.CommitmentRandomness

	return nil
}

// NewRecord returns a new Record from the given inputs.
func NewRecord(owner *account.Address, value int64, payload []byte, programID string, serialNumberNonce string, commitmentRandomness string) *Record {
	return &Record{
		owner:                owner,
		value:                value,
		payload:              payload,
		programID:            programID,
		serialNumberNonce:    serialNumberNonce,
		commitmentRandomness: commitmentRandomness,
	}
}

// Owner returns the Record's owner.
func (r Record) Owner() *account.Address {
	return r.owner.Copy()
}

// Value returns the Record's value.
func (r Record) Value() int64 {
	return r.value
}

// Payload returns the Record's payload.
func (r Record) Payload() []byte {
	return r.payload
}

// ProgramID returns the Record's programID.
func (r Record) ProgramID() string {
	return r.programID
}

// SerialNumberNonce returns the Record's serial number nonce.
func (r Record) SerialNumberNonce() string {
	return r.serialNumberNonce
}

// CommitmentRandomness returns the Record's commitment randomness.
func (r Record) CommitmentRandomness() string {
	return r.commitmentRandomness
}

// Owner returns the Record's owner.
func (r Record) String() string {
	return fmt.Sprintf("owner: %v\nvalue: %v\n,payload: %v\n, programID: %s\nserialNumberNonce: %s\ncommitmentRandomness: %s\n",
		r.owner, r.value, r.payload, r.programID, r.serialNumberNonce, r.commitmentRandomness)
}

// NewInputRecord creates a new record.
func NewInputRecord(address *account.Address, value int64, payload [128]byte, randomness []byte) (*Record, error) {
	return newInputRecord(address, value, payload, randomness)
}

// EncryptRecord encrypts a record.
func EncryptRecord(record *Record, encrypt []byte) (string, error) {
	return encryptRecord(record, encrypt)
}

// DecryptRecord decrypts a record.
func DecryptRecord(ciphertext string, viewKey *account.ViewKey) (*Record, error) {
	return decryptRecord(ciphertext, viewKey)
}
