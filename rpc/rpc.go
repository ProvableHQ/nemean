// Package rpc provides an RPC client to the snarkOS Aleo client.
package rpc

import (
	"bytes"
	"encoding/json"
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

	// Handle response!
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result Result
	if err := json.Unmarshal(buf, &result); err != nil {
		return nil, err
	}

	if len(result.Error) > 0 {
		return nil, fmt.Errorf("%v", result.Error)
	}

	return &result, nil
}

type Request struct {
	Method     string            `json:"method"`
	Params     []json.RawMessage `json:"params"`
	Id         string            `json:"id"`
	RpcVersion string            `json:"jsonrpc"`
}

func newRequestBody(rpcVersion int, id string, method string, params []json.RawMessage) ([]byte, error) {
	req := &Request{
		Method:     method,
		Params:     params,
		Id:         id,
		RpcVersion: "2.0",
	}
	return json.Marshal(req)
}

// TODO decoderawtransaction

// GetBlockCount returns the number of blocks in the best valid chain.
func (c *Client) GetBlockCount() (int64, error) {
	req, err := newRequestBody(2, "", getBlockCountMethod, nil)
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

// GetBlock returns information about a block from a block hash.
func (c *Client) GetBlock(block string) (*GetBlockResponse, error) {
	param, err := json.Marshal(block)
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

	var res GetBlockResponse
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

// GetBlockTemplate returns the current mempool and consensus information known by this node.
func (c *Client) GetBlockTemplate() (res *GetBlockTemplateResponse, err error) {
	req, err := newRequestBody(2, "", getBlockTemplateMethod, nil)
	if err != nil {
		return nil, err
	}

	resp, err := newRequest(c, req)
	if err != nil {
		return nil, err
	}

	var result Result
	if err := json.Unmarshal(resp.Result, &result); err != nil {
		return nil, err
	}

	return res, nil
}

// GetConnectionCount returns the number of connected peers this node has.
func (c *Client) GetConnectionCount() (int, error) {
	req, err := newRequestBody(2, "", getConnectionCountMethod, nil)
	if err != nil {
		return 0, err
	}

	resp, err := newRequest(c, req)
	if err != nil {
		return 0, err
	}

	var res int
	if err := json.Unmarshal(resp.Result, &res); err != nil {
		return 0, err
	}

	return res, nil
}

// GetRawTransaction returns hex encoded bytes of a transaction from its transaction id.
func (c *Client) GetRawTransaction(txID string) (string, error) {
	param, err := json.Marshal(txID)
	if err != nil {
		return "", err
	}

	req, err := newRequestBody(2, "", getRawTransactionMethod, []json.RawMessage{param})
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

// GetTransactionInfo returns information about a transaction from a transaction id.
func (c *Client) GetTransactionInfo(txID string) (*GetTransactionInfoResponse, error) {
	param, err := json.Marshal(txID)
	if err != nil {
		return nil, err
	}

	req, err := newRequestBody(2, "", getTransactionInfoMethod, []json.RawMessage{param})
	if err != nil {
		return nil, err
	}

	resp, err := newRequest(c, req)
	if err != nil {
		return nil, err
	}

	var res GetTransactionInfoResponse
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

	req, err := newRequestBody(2, "documentation", sendTransactionMethod, []json.RawMessage{param})
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
