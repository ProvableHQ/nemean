package main

import (
	"errors"
	"github.com/pinestreetlabs/aleo-wallet-sdk/rpc"
	"github.com/urfave/cli"
	"log"
	"os"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "nemean"
	app.Usage = "An unfairly private wallet for Aleo."
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "rpchost",
			Value: "",
			Usage: "the host:port of SnarkOS",
		},
	}

	app.Commands = []cli.Command{
		newAccountCommand,
		getBlockCommand,
		getBlockHashCommand,
		getBlockHashesCommand,
		getBlockHeaderCommand,
		getBlockHeightCommand,
		latestBlockHeightCommand,
		getTransactionCommand,
		newTransactionCommand,
		latestLedgerRootCommand,
		sendTransactionCommand,
		getLedgerProofCommand,
		decryptRecordCommand,
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

type profile struct {
	host string
	port string
}

func getProfile(ctx *cli.Context) (*profile, error) {
	snarkos := strings.Split(ctx.GlobalString("rpchost"), ":")
	if len(snarkos) != 2 {
		return nil, errors.New("invalid rpchost")
	}
	return &profile{
		host: snarkos[0],
		port: snarkos[1],
	}, nil
}

func getClient(host, port string) (*rpc.Client, error) {
	return rpc.NewClient(&rpc.Config{
		User:     "",
		Password: "",
		Host:     host,
		Port:     port,
	})
}
