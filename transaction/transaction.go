package transaction

import (
	"github.com/pinestreetlabs/aleo-wallet-sdk/account"
	"github.com/pinestreetlabs/aleo-wallet-sdk/record"
)

type Transaction struct{}

func NewCoinbaseTransaction(address *account.Address, value int64, random []byte) string {
	return newCoinbaseTransaction(address, value, random)
}

func NewTransferTransaction(account *account.Account, in *record.Record, ledgerProofs []string, amount, fee int64) (string, error) {
	return newTransferTransaction(account, in, ledgerProofs, amount, fee)
}
