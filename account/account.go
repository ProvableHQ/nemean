package account

import (
	"encoding/json"
	"github.com/pinestreetlabs/aleo-wallet-sdk/network"
)

// Account encompasses an Aleo Account.
type Account struct {
	privateKey *PrivateKey
	viewKey    *ViewKey
	address    *Address
}

type AccountJSON struct {
	PrivateKey string
	ViewKey    string
	Address    string
}

func (a *Account) MarshalJSON() ([]byte, error) {
	return json.Marshal(AccountJSON{
		PrivateKey: a.privateKey.String(),
		ViewKey:    a.viewKey.String(),
		Address:    a.address.String(),
	})
}

func (a *Account) UnmarshalJSON(b []byte) error {
	temp := &AccountJSON{}

	if err := json.Unmarshal(b, &temp); err != nil {
		return err
	}

	address, err := ParseAddress(temp.Address)
	if err != nil {
		return err
	}

	viewKey, err := ParseViewKey(temp.ViewKey)
	if err != nil {
		return err
	}

	privateKey, err := ParsePrivateKey(temp.PrivateKey)
	if err != nil {
		return err
	}

	a.privateKey = privateKey
	a.viewKey = viewKey
	a.address = address

	return nil
}

// ViewKey returns a copy of the Account's ViewKey.
func (a *Account) ViewKey() *ViewKey {
	return a.viewKey.Copy()
}

// Address returns a copy of the Account's Address.
func (a *Account) Address() *Address {
	return a.address.Copy()
}

// PrivateKey returns a copy of the Account's PrivateKey
func (a *Account) PrivateKey() *PrivateKey {
	return a.privateKey.Copy()
}

// FromSeed creates a new Account with the given 32 byte seed.
func FromSeed(seed [32]byte, params *network.Params) (*Account, error) {
	return fromSeed(seed, params)
}

// FromPrivateKey creates a new Account with the given privateKey.
func FromPrivateKey(privateKey string, params *network.Params) (*Account, error) {
	return fromPrivateKey(privateKey)
}
