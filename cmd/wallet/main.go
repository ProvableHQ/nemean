package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
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
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
