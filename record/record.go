package record

import (
	"encoding/json"
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

type RecordJSON struct {
	Owner                string `json:"owner"`
	Value                int64  `json:"value"`
	Payload              []byte `json:"payload"`
	ProgramID            string `json:"program_id"`
	SerialNumberNonce    string `json:"serial_number_nonce"`
	CommitmentRandomness string `json:"commitment_randomness"`
}

func (r *Record) MarshalJSON() ([]byte, error) {
	return json.Marshal(RecordJSON{
		Owner:                r.owner.String(),
		Value:                r.value,
		Payload:              r.payload,
		ProgramID:            r.programID,
		SerialNumberNonce:    r.serialNumberNonce,
		CommitmentRandomness: r.commitmentRandomness,
	})
}

func (r *Record) UnmarshalJSON(b []byte) error {
	temp := &RecordJSON{}

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
