package main

import "github.com/urfave/cli"

var newAccountCommand = cli.Command{
	Name:     "create",
	Category: "wallet",
	Usage:    "Create a new Aleo account.",
	Description: `
	The create command is used to create a new Aleo account.
	`,
	Action: newAccount,
}

var getBlockCommand = cli.Command{
	Name:     "getBlock",
	Category: "rpc",
	Usage:    "Get a block.",
	Description: `
	Gets a block from SnarkOS.
	`,
	Flags: []cli.Flag{
		cli.Int64Flag{
			Name:  "height",
			Usage: "block height",
		},
	},
	Action: getBlock,
}
