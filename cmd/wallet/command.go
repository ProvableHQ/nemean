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
