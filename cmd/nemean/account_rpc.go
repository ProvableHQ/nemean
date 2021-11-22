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

func getBlockHash(ctx *cli.Context) error {
	profile, err := getProfile(ctx)
	if err != nil {
		return err
	}

	client, err := getClient(profile.host, profile.port)
	if err != nil {
		return err
	}

	resp, err := client.GetBlockHash(ctx.Int64("height"))
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", resp)

	return nil
}

func getBlockHashes(ctx *cli.Context) error {
	profile, err := getProfile(ctx)
	if err != nil {
		return err
	}

	client, err := getClient(profile.host, profile.port)
	if err != nil {
		return err
	}

	resp, err := client.GetBlockHashes(ctx.Int64("start"), ctx.Int64("end"))
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", resp)

	return nil
}

func getBlockHeader(ctx *cli.Context) error {
	profile, err := getProfile(ctx)
	if err != nil {
		return err
	}

	client, err := getClient(profile.host, profile.port)
	if err != nil {
		return err
	}

	resp, err := client.GetBlockHeader(ctx.Int64("height"))
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", resp)

	return nil
}

func getTransaction(ctx *cli.Context) error {
	profile, err := getProfile(ctx)
	if err != nil {
		return err
	}

	client, err := getClient(profile.host, profile.port)
	if err != nil {
		return err
	}

	resp, err := client.GetTransaction(ctx.String("id"))
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

func getBlockHeight(ctx *cli.Context) error {
	profile, err := getProfile(ctx)
	if err != nil {
		return err
	}

	client, err := getClient(profile.host, profile.port)
	if err != nil {
		return err
	}

	resp, err := client.GetBlockHeight(ctx.String("hash"))
	if err != nil {
		return err
	}

	fmt.Printf("%d\n", resp)

	return nil
}

func latestBlockHeight(ctx *cli.Context) error {
	profile, err := getProfile(ctx)
	if err != nil {
		return err
	}

	client, err := getClient(profile.host, profile.port)
	if err != nil {
		return err
	}

	resp, err := client.LatestBlockHeight()
	if err != nil {
		return err
	}

	fmt.Printf("%d\n", resp)

	return nil
}

func sendTransaction(ctx *cli.Context) error {
	profile, err := getProfile(ctx)
	if err != nil {
		return err
	}

	client, err := getClient(profile.host, profile.port)
	if err != nil {
		return err
	}

	resp, err := client.SendTransaction(ctx.String("txn"))
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

func latestLedgerRoot(ctx *cli.Context) error {
	profile, err := getProfile(ctx)
	if err != nil {
		return err
	}

	client, err := getClient(profile.host, profile.port)
	if err != nil {
		return err
	}

	resp, err := client.LatestLedgerRoot()
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", resp)

	return nil
}

func getLedgerProof(ctx *cli.Context) error {
	profile, err := getProfile(ctx)
	if err != nil {
		return err
	}

	client, err := getClient(profile.host, profile.port)
	if err != nil {
		return err
	}

	resp, err := client.GetLedgerProof(ctx.String("commitment"))
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", resp)

	return nil
}
