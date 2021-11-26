package rpc

import (
	"encoding/json"
	"math/big"
)

type Request struct {
	Method     string            `json:"method"`
	Params     []json.RawMessage `json:"params"`
	Id         string            `json:"id"`
	RpcVersion string            `json:"jsonrpc"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type Result struct {
	Result json.RawMessage `json:"result"`
	Error  *Error          `json:"error,omitempty"`
	ID     string          `json:"id"`
}

type Transactions struct {
	Transactions []Transaction `json:"transactions"`
}

type Block struct {
	BlockHash         string       `json:"block_hash"`
	PreviousBlockHash string       `json:"previous_block_hash"`
	Transactions      Transactions `json:"transactions"`
	BlockHeader       BlockHeader  `json:"header"`
}

type BlockHeader struct {
	PrevLedgerRoot   string              `json:"previous_ledger_root"`
	TransactionsRoot string              `json:"transactions_root"`
	Proof            string              `json:"string"`
	Metadata         BlockHeaderMetadata `json:"metadata"`
}

type BlockHeaderMetadata struct {
	DifficultyTarget big.Int `json:"difficulty_target"`
	Height           int64   `json:"height"`
	Nonce            string  `json:"nonce"`
	Timestamp        int64   `json:"timestamp"`
}

type GetBlockTemplateResponse struct {
	PreviousBlockHash string   `json:"previous_block_hash"`
	BlockHeight       int64    `json:"block_height"`
	Time              int64    `json:"time"`
	DifficultyTarget  int64    `json:"difficulty_target"`
	Transactions      []string `json:"transactions"` //todo
	CoinbaseValue     int64    `json:"coinbase_value"`
}

type GetConnectionCountResponse struct {
	Result int64 `json:"result"`
}

type GetRawTransactionResult struct {
	Result string `json:"result"`
}

type GetTransactionResponse struct {
	Transaction Transaction         `json:"transaction"`
	Metadata    TransactionMetadata `json:"metadata"`
}

type TransactionMetadata struct {
	BlockHash        string `json:"block_hash"`
	BlockHeight      int64  `json:"block_height"`
	BlockTimestamp   int64  `json:"block_timestamp"`
	TransactionIndex int64  `json:"transaction_index"`
}

type Transaction struct {
	TxId           string       `json:"transaction_id"`
	LedgerRoot     string       `json:"ledger_root"`
	InnerCircuitID string       `json:"inner_circuit_id"`
	Transitions    []Transition `json:"transitions"`
}

type Transition struct {
	CiphertextIDs []string `json:"ciphertext_ids"`
	Ciphertexts   []string `json:"ciphertexts"`
	Commitments   []string `json:"commitments"`
	Proof         string   `json:"proof"`
	SerialNumbers []string `json:"serial_numbers"`
	ID            string   `json:"transition_id"`
	ValueBalance  int64    `json:"value_balance"`
}

type SendTransactionResponse struct {
	Result string `json:"result"`
}

type CreateAccountResponse struct {
	PrivateKey string `json:"private_key"`
	ViewKey    string `json:"view_key"`
	Address    string `json:"address"`
}

type CreateRawTransactionResponse struct {
	EncodedTransaction string   `json:"encoded_transaction"`
	EncodedRecords     []string `json:"encoded_records"`
}

type CreateTransactionResponse struct {
	EncodedTransaction string   `json:"encoded_transaction"`
	EncodedRecords     []string `json:"encoded_records"`
}

type CreateTransactionKernelResponse struct {
	TransactionKernel string `json:"transaction_kernel"`
}
