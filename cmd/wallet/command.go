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

var newTransactionCommand = cli.Command{
	Name:     "send",
	Category: "wallet",
	Usage:    "Create a basic transfer transaction.",
	Description: `
	The send command creates a single transfer transaction that consumes
	a single record and returns a serialized transaction in hex.
	`,
	Action: newTransaction,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "to",
			Usage:    "recipient address",
			Required: true,
		},
		cli.StringSliceFlag{
			Name:     "ledger_proof",
			Usage:    "list of ledger proofs for input record",
			Required: true,
		},
		cli.Int64Flag{
			Name:     "amount",
			Usage:    "amount to send",
			Required: true,
		},
		cli.Int64Flag{
			Name:     "fee",
			Usage:    "network fee",
			Required: true,
		},
		cli.StringFlag{
			Name:     "private_key",
			Usage:    "private key to sign transaction",
			Required: true,
		},
		cli.StringFlag{
			Name:     "record",
			Usage:    "JSON input record to consume",
			Required: true,
		},
	},
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
			Name:     "height",
			Usage:    "block height",
			Required: true,
		},
	},
	Action: getBlock,
}
