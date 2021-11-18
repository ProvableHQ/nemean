package main

import (
	"fmt"
	"github.com/pinestreetlabs/aleo-wallet-sdk/account"
	"github.com/pinestreetlabs/aleo-wallet-sdk/network"
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

	fmt.Println("==============")
	fmt.Println("created account")
	fmt.Println(acc.PrivateKey().String())
	fmt.Println(acc.ViewKey().String())
	fmt.Println(acc.Address().String())
	fmt.Println("==============")
	return nil
}
