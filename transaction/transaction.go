package transaction

import "github.com/pinestreetlabs/aleo-wallet-sdk/account"

type Transaction struct{}

func NewCoinbaseTransaction(address *account.Address, value int64, random []byte) string {
	return newCoinbaseTransaction(address, value, random)
}
