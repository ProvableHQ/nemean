package account

// Account encompasses an Aleo Account.
type Account struct {
	privateKey *PrivateKey
	viewKey    *ViewKey
	address    *Address
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
func FromSeed(seed [32]byte) (*Account, error) {
	return fromSeed(seed)
}

// TODO FromPrivateKey(..) Account
