// Package rpc provides an RPC client to the snarkOS Aleo client.
package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Config holds the configuration for the RPC client.
type Config struct {
	User     string
	Password string
	Host     string
	Port     string
}

// Client maintains a connection to the Aleo client.
type Client struct {
	cfg        *Config
	httpClient *http.Client
}

// NewClient returns a new RPC client.
func NewClient(cfg *Config) (*Client, error) {
	httpClient := &http.Client{}

	return &Client{
		httpClient: httpClient,
		cfg:        cfg,
	}, nil
}

// newRequest creates a serialized request.
func newRequest(client *Client, body []byte) (*Result, error) {
	url := fmt.Sprintf("http://%s:%s", client.cfg.Host, client.cfg.Port)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("content-type", "application/json")

	req.SetBasicAuth(client.cfg.User, client.cfg.Password)

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result Result
	if err := json.Unmarshal(buf, &result); err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, fmt.Errorf("%v", result.Error)
	}

	return &result, nil
}

func newRequestBody(rpcVersion int, id string, method string, params []json.RawMessage) ([]byte, error) {
	req := &Request{
		Method:     method,
		Params:     params,
		ID:         id,
		RPCVersion: "2.0",
	}
	return json.Marshal(req)
}

// GetBlockHeight the block height for the given the block hash.
func (c *Client) GetBlockHeight(blockHash string) (int64, error) {
	param, err := json.Marshal(blockHash)
	if err != nil {
		return 0, err
	}

	req, err := newRequestBody(2, "", getBlockHeightMethod, []json.RawMessage{param})
	if err != nil {
		return 0, err
	}

	resp, err := newRequest(c, req)
	if err != nil {
		return 0, err
	}

	var count int64
	if err := json.Unmarshal(resp.Result, &count); err != nil {
		return 0, err
	}

	return count, nil
}

// GetBestBlockHash returns the block hash of the head of the best valid chain.
func (c *Client) GetBestBlockHash() (string, error) {
	req, err := newRequestBody(2, "", getBestBlockHashMethod, nil)
	if err != nil {
		return "", err
	}

	resp, err := newRequest(c, req)
	if err != nil {
		return "", err
	}

	var res string
	if err := json.Unmarshal(resp.Result, &res); err != nil {
		return "", err
	}

	return res, nil
}

// GetBlock returns information about a block from a block height.
func (c *Client) GetBlock(blockNumber int64) (*Block, error) {
	param, err := json.Marshal(blockNumber)
	if err != nil {
		return nil, err
	}

	req, err := newRequestBody(2, "", getBlockMethod, []json.RawMessage{param})
	if err != nil {
		return nil, err
	}

	resp, err := newRequest(c, req)
	if err != nil {
		return nil, err
	}

	var res Block
	if err := json.Unmarshal(resp.Result, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// GetBlockHash returns the block hash of a block at the given block height in the best valid chain.
func (c *Client) GetBlockHash(height int64) (string, error) {
	param, err := json.Marshal(height)
	if err != nil {
		return "", err
	}

	req, err := newRequestBody(2, "", getBlockHashMethod, []json.RawMessage{param})
	if err != nil {
		return "", err
	}

	resp, err := newRequest(c, req)
	if err != nil {
		return "", err
	}

	var res string
	if err := json.Unmarshal(resp.Result, &res); err != nil {
		return "", err
	}

	return res, nil
}

// GetTransaction returns a transaction with metadata given the transaction ID.
func (c *Client) GetTransaction(txID string) (*GetTransactionResponse, error) {
	param, err := json.Marshal(txID)
	if err != nil {
		return nil, err
	}

	req, err := newRequestBody(2, "", getTransactionMethod, []json.RawMessage{param})
	if err != nil {
		return nil, err
	}

	resp, err := newRequest(c, req)
	if err != nil {
		return nil, err
	}

	var res GetTransactionResponse
	if err := json.Unmarshal(resp.Result, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// GetTransition returns a transition given an id.
func (c *Client) GetTransition(transitionID string) (*Transition, error) {
	param, err := json.Marshal(transitionID)
	if err != nil {
		return nil, err
	}

	req, err := newRequestBody(2, "", getTransitionMethod, []json.RawMessage{param})
	if err != nil {
		return nil, err
	}

	resp, err := newRequest(c, req)
	if err != nil {
		return nil, err
	}

	var res Transition
	if err := json.Unmarshal(resp.Result, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// SendTransaction sends raw transaction bytes to this node to be added into the mempool.
// If valid, the transaction will be stored and propagated to all peers.
func (c *Client) SendTransaction(txHex string) (string, error) {
	param, err := json.Marshal(txHex)
	if err != nil {
		return "", err
	}

	req, err := newRequestBody(2, "", sendTransactionMethod, []json.RawMessage{param})
	if err != nil {
		return "", err
	}

	resp, err := newRequest(c, req)
	if err != nil {
		return "", err
	}

	var res string
	if err := json.Unmarshal(resp.Result, &res); err != nil {
		return "", err
	}

	return res, nil
}

// LatestLedgerRoot returns the latest ledger root.
func (c *Client) LatestLedgerRoot() (string, error) {
	req, err := newRequestBody(2, "", latestLedgerRootMethod, nil)
	if err != nil {
		return "", err
	}

	resp, err := newRequest(c, req)
	if err != nil {
		return "", err
	}

	var res string
	if err := json.Unmarshal(resp.Result, &res); err != nil {
		return "", err
	}

	return res, nil
}

// GetLedgerProof returns the ledger proof of a given record commitment.
func (c *Client) GetLedgerProof(recordCommitment string) (string, error) {
	param, err := json.Marshal(recordCommitment)
	if err != nil {
		return "", err
	}

	req, err := newRequestBody(2, "", getLedgerProofMethod, []json.RawMessage{param})
	if err != nil {
		return "", err
	}

	resp, err := newRequest(c, req)
	if err != nil {
		return "", err
	}

	var res string
	if err := json.Unmarshal(resp.Result, &res); err != nil {
		return "", err
	}

	return res, nil
}

// GetCiphertext returns the ciphertext using a given id.
func (c *Client) GetCiphertext(id string) (string, error) {
	param, err := json.Marshal(id)
	if err != nil {
		return "", err
	}

	req, err := newRequestBody(2, "", getCiphertextMethod, []json.RawMessage{param})
	if err != nil {
		return "", err
	}

	resp, err := newRequest(c, req)
	if err != nil {
		return "", err
	}

	var res string
	if err := json.Unmarshal(resp.Result, &res); err != nil {
		return "", err
	}

	return res, nil
}

// GetBlockHashes returns blockhashes for a given range of block heights.
func (c *Client) GetBlockHashes(start, end int64) ([]string, error) {
	if start > end {
		return nil, errors.New("start > end")
	}

	startParam, err := json.Marshal(start)
	if err != nil {
		return nil, err
	}

	endParam, err := json.Marshal(end)
	if err != nil {
		return nil, err
	}

	req, err := newRequestBody(2, "", getBlockHashesMethod, []json.RawMessage{startParam, endParam})
	if err != nil {
		return nil, err
	}

	resp, err := newRequest(c, req)
	if err != nil {
		return nil, err
	}

	var res []string
	if err := json.Unmarshal(resp.Result, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetBlockHeader returns the blockheader of a given block height.
func (c *Client) GetBlockHeader(height int64) (*BlockHeader, error) {
	param, err := json.Marshal(height)
	if err != nil {
		return nil, err
	}

	req, err := newRequestBody(2, "", getBlockHeaderMethod, []json.RawMessage{param})
	if err != nil {
		return nil, err
	}

	resp, err := newRequest(c, req)
	if err != nil {
		return nil, err
	}

	var res BlockHeader
	if err := json.Unmarshal(resp.Result, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// GetBlocks returns a range of blocks with the given block heights.
func (c *Client) GetBlocks(start, end int64) ([]Block, error) {
	if start > end {
		return nil, errors.New("start > end")
	}

	startParam, err := json.Marshal(start)
	if err != nil {
		return nil, err
	}

	endParam, err := json.Marshal(end)
	if err != nil {
		return nil, err
	}

	req, err := newRequestBody(2, "", getBlocksMethod, []json.RawMessage{startParam, endParam})
	if err != nil {
		return nil, err
	}

	resp, err := newRequest(c, req)
	if err != nil {
		return nil, err
	}

	var res []Block
	if err := json.Unmarshal(resp.Result, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetBlockTransactions returns transactions in a block at a given height.
func (c *Client) GetBlockTransactions(height int64) (*Transactions, error) {
	param, err := json.Marshal(height)
	if err != nil {
		return nil, err
	}

	req, err := newRequestBody(2, "", getBlockTransactionsMethod, []json.RawMessage{param})
	if err != nil {
		return nil, err
	}

	resp, err := newRequest(c, req)
	if err != nil {
		return nil, err
	}

	var res Transactions
	if err := json.Unmarshal(resp.Result, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// LatestBlock returns the latest block.
func (c *Client) LatestBlock() (*Block, error) {
	req, err := newRequestBody(2, "", latestBlockMethod, nil)
	if err != nil {
		return nil, err
	}

	resp, err := newRequest(c, req)
	if err != nil {
		return nil, err
	}

	var res Block
	if err := json.Unmarshal(resp.Result, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// LatestBlockHash returns the latest blockhash.
func (c *Client) LatestBlockHash() (string, error) {
	req, err := newRequestBody(2, "", latestBlockHashMethod, nil)
	if err != nil {
		return "", err
	}

	resp, err := newRequest(c, req)
	if err != nil {
		return "", err
	}

	var res string
	if err := json.Unmarshal(resp.Result, &res); err != nil {
		return "", err
	}

	return res, nil
}

// LatestBlockHeader returns the latest blcok header.
func (c *Client) LatestBlockHeader() (*BlockHeader, error) {
	req, err := newRequestBody(2, "", latestBlockHeaderMethod, nil)
	if err != nil {
		return nil, err
	}

	resp, err := newRequest(c, req)
	if err != nil {
		return nil, err
	}

	var res BlockHeader
	if err := json.Unmarshal(resp.Result, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// LatestBlockHeight returns the latest blockheight.
func (c *Client) LatestBlockHeight() (int64, error) {
	req, err := newRequestBody(2, "", latestBlockHeightMethod, nil)
	if err != nil {
		return 0, err
	}

	resp, err := newRequest(c, req)
	if err != nil {
		return 0, err
	}

	var res int64
	if err := json.Unmarshal(resp.Result, &res); err != nil {
		return 0, err
	}

	return res, nil
}

// LatestBlockTransactions returns a list of transactions for the latest block.
func (c *Client) LatestBlockTransactions() (*Transactions, error) {
	req, err := newRequestBody(2, "", latestBlockTransactionsMethod, nil)
	if err != nil {
		return nil, err
	}

	resp, err := newRequest(c, req)
	if err != nil {
		return nil, err
	}

	var res Transactions
	if err := json.Unmarshal(resp.Result, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// GetConnectedPeers returns a list of connected peers.
func (c *Client) GetConnectedPeers() ([]string, error) {
	req, err := newRequestBody(2, "", getConnectedPeersMethod, nil)
	if err != nil {
		return nil, err
	}

	resp, err := newRequest(c, req)
	if err != nil {
		return nil, err
	}

	var res []string
	if err := json.Unmarshal(resp.Result, &res); err != nil {
		return nil, err
	}

	return res, nil
}
