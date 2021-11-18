package main

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
)

func getBlock(ctx *cli.Context) error {
	profile, err := getProfile(ctx)
	if err != nil {
		return err
	}

	client, err := getClient(profile.host, profile.port)
	if err != nil {
		return err
	}

	resp, err := client.GetBlock(ctx.Int64("height"))
	if err != nil {
		return err
	}

	body, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", body)

	return nil
}
