package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/pinestreetlabs/aleo-wallet-sdk/account"
	"github.com/pinestreetlabs/aleo-wallet-sdk/network"
	"github.com/pinestreetlabs/aleo-wallet-sdk/record"
	"github.com/pinestreetlabs/aleo-wallet-sdk/transaction"
	"github.com/urfave/cli"
)

var errInvalidSeed = errors.New("invalid seed")

func newAccount(ctx *cli.Context) (err error) {
	var seed [32]byte
	if ctx.NumFlags() == 1 {
		in := ctx.String("from")

		buf, err := base64.StdEncoding.DecodeString(in)
		if err != nil {
			return fmt.Errorf("%w : %v", errInvalidSeed, err)
		}

		if len(buf) != 32 {
			return fmt.Errorf("%w : got len %d", errInvalidSeed, len(buf))
		}

		copy(seed[:], buf)
	} else {
		seed, err = account.NewSeed()
		if err != nil {
			return err
		}
	}

	acc, err := account.FromSeed(seed, network.Testnet2())
	if err != nil {
		return err
	}

	resp, err := json.Marshal(acc)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", resp)
	return nil
}

func fromAccount(ctx *cli.Context) error {
	key := ctx.String("from")

	acc, err := account.FromPrivateKey(key, network.Testnet2())
	if err != nil {
		return err
	}

	resp, err := json.Marshal(acc)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", resp)
	return nil
}

func newTransaction(ctx *cli.Context) error {
	inputRec := ctx.String("record")

	var rec record.Record
	if err := json.Unmarshal([]byte(inputRec), &rec); err != nil {
		return err
	}

	to, err := account.ParseAddress(ctx.String("to"))
	if err != nil {
		return err
	}

	sk, err := account.ParsePrivateKey(ctx.String("private_key"))
	if err != nil {
		return err
	}

	proofs := ctx.StringSlice("ledger_proof")
	amount := ctx.Int64("amount")
	fee := ctx.Int64("fee")

	txn, err := transaction.NewTransferTransaction(sk, to, &rec, proofs, amount, fee)
	if err != nil {
		return err
	}

	fmt.Println(txn)
	return nil
}

func decryptRecord(ctx *cli.Context) error {
	vk, err := account.ParseViewKey(ctx.String("viewkey"))
	if err != nil {
		return err
	}

	rec, err := record.DecryptRecord(ctx.String("ciphertext"), vk)
	if err != nil {
		return err
	}

	resp, err := json.Marshal(rec)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", resp)
	return nil
}
