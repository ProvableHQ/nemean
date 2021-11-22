package main

import (
	"encoding/json"
	"fmt"
	"github.com/pinestreetlabs/aleo-wallet-sdk/account"
	"github.com/pinestreetlabs/aleo-wallet-sdk/network"
	"github.com/pinestreetlabs/aleo-wallet-sdk/record"
	"github.com/pinestreetlabs/aleo-wallet-sdk/transaction"
	"github.com/urfave/cli"
)

func newAccount(ctx *cli.Context) error {
	seed, err := account.NewSeed()
	if err != nil {
		return err
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

	fmt.Println(rec)
	return nil
}
