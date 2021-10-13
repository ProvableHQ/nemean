// Package rpc provides an RPC client to the Aleo Daemon.
package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Config struct {
	User string
	Password string
	Host string
	Port string
}

type Client struct {
	cfg *Config
	httpClient *http.Client
}

// NewClient returns a new RPC client.
func NewClient(cfg *Config) (*Client, error) {
	httpClient := &http.Client{}

	return &Client{
		httpClient: httpClient,
		cfg: cfg,
	}, nil
}

func newRequest(client *Client, body []byte) ([]byte, error) {

	url := fmt.Sprintf("http://%s:%s",client.cfg.Host, client.cfg.Port)

	fmt.Printf("%v\n",body)
	req, err := http.NewRequest("POST", url,bytes.NewBuffer(body))
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
	return ioutil.ReadAll(resp.Body)
}

type Request struct {
	Method string `json:"method"`
	Params []json.RawMessage `json:"params"`
	Id string `json:"id"`
	RpcVersion string `json:"jsonrpc"`
}

func newRequestBody(rpcVersion int, id string, method string, params []interface{}) ([]byte, error) {
	req := &Request{
		Method:     method,
		Params:     []json.RawMessage{},
		Id:         id,
		RpcVersion: "2.0",
	}
	fmt.Printf("%v\n",req)
	return json.Marshal(req)
}


type GetBlockCountResponse struct {
	Result int64 `json:"result"`
}

// TODO decoderawtransaction

// GetBlockCount returns the number of blocks in the best valid chain.
// https://github.com/AleoHQ/snarkOS/blob/master/rpc/README.md#getblockcount
func (c *Client) GetBlockCount() (int64, error) {
	req, err := newRequestBody(2, "documentation","getblockcount" ,nil)
	if err != nil {
		return 0, err
	}

	resp, err := newRequest(c, req)
	if err != nil {
		return 0, err
	}

	var count GetBlockCountResponse
	if err := json.Unmarshal(resp, &count); err != nil {
		return 0, err
	}

	return count.Result, nil
}

// GetBestBlockhash returns the block hash of the head of the best valid chain.
func (c *Client) GetBestBlockhash() (string, error) {
	req, err := newRequestBody(2, "documentation",getBestBlockHashMethod ,nil)
	if err != nil {
		return "", err
	}

	resp, err := newRequest(c, req)
	if err != nil {
		return "", err
	}

	var res Result
	if err := json.Unmarshal(resp, &res); err != nil {
		return "", err
	}

	return res.Result, nil
}

func (c *Client) GetBlock() (string, error) {
	req, err := newRequestBody(2, "documentation",getBlockMethod ,nil)
	if err != nil {
		return "", err
	}

	resp, err := newRequest(c, req)
	if err != nil {
		return "", err
	}

	var res Result
	if err := json.Unmarshal(resp, &res); err != nil {
		return "", err
	}

	return res.Result, nil
}

func (c * Client) GetBlockHash(height int64) (string, error) {
	req, err := newRequestBody(2, "documentation",getBlockHashMethod ,nil)
	if err != nil {
		return "", err
	}

	resp, err := newRequest(c, req)
	if err != nil {
		return "", err
	}

	var res Result
	if err := json.Unmarshal(resp, &res); err != nil {
		return "", err
	}

	return res.Result, nil
}

func (c *Client) GetBlockTemplate() (*GetBlockTemplateResponse, error) {
	req, err := newRequestBody(2, "documentation",getBlockTemplateMethod ,nil)
	if err != nil {
		return nil, err
	}

	resp, err := newRequest(c, req)
	if err != nil {
		return nil, err
	}

	var res GetBlockTemplateResponse
	if err := json.Unmarshal(resp, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetConnectionCount() (int, error) {
	req, err := newRequestBody(2, "documentation",getConnectionCountMethod ,nil)
	if err != nil {
		return 0, err
	}

	resp, err := newRequest(c, req)
	if err != nil {
		return 0, err
	}

	var res Result
	if err := json.Unmarshal(resp, &res); err != nil {
		return 0, err
	}

	count, err := strconv.ParseInt(res.Result,10,64)
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (c *Client) GetRawTransaction(txID string) (string, error) {
	req, err := newRequestBody(2, "documentation",getRawTransactionMethod ,nil)
	if err != nil {
		return "", err
	}

	resp, err := newRequest(c, req)
	if err != nil {
		return "", err
	}

	var res Result
	if err := json.Unmarshal(resp, &res); err != nil {
		return "", err
	}

	return res.Result, nil
}

func (c *Client) GetTransactionInfo(txID string) (*GetTransactionInfoResponse, error) {
	req, err := newRequestBody(2, "documentation",getTransactionInfoMethod ,nil)
	if err != nil {
		return nil, err
	}

	resp, err := newRequest(c, req)
	if err != nil {
		return nil, err
	}

	var res GetTransactionInfoResponse
	if err := json.Unmarshal(resp, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) SendTransaction(txHex string) (string, error) {
	req, err := newRequestBody(2, "documentation",sendTransactionMethod ,nil)
	if err != nil {
		return "", err
	}

	resp, err := newRequest(c, req)
	if err != nil {
		return "", err
	}

	var res Result
	if err := json.Unmarshal(resp, &res); err != nil {
		return "", err
	}

	return res.Result, nil
}