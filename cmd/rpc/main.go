package main

import (
	"fmt"
	"github.com/pinestreetlabs/aleo-wallet-sdk/rpc"
)

func main() {
	cfg := &rpc.Config{
		User:     "aleo",
		Password: "password",
		Host:     "127.0.01",
		Port:     "3030",
	}

	client, err := rpc.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	result, err := client.GetBestBlockHash()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", result)
}
