package rpc

import (
	"encoding/json"
	"math/big"
)

// Request is the RPC request object.
type Request struct {
	Method     string            `json:"method"`
	Params     []json.RawMessage `json:"params"`
	ID         string            `json:"id"`
	RPCVersion string            `json:"jsonrpc"`
}

// Error is the RPC error response.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

// Result is the RPC result response.
type Result struct {
	Result json.RawMessage `json:"result"`
	Error  *Error          `json:"error,omitempty"`
	ID     string          `json:"id"`
}

// Transactions is a list of Transactions.
type Transactions struct {
	Transactions []Transaction `json:"transactions"`
}

// Block contains the block response object.
type Block struct {
	BlockHash         string       `json:"block_hash"`
	PreviousBlockHash string       `json:"previous_block_hash"`
	Transactions      Transactions `json:"transactions"`
	BlockHeader       BlockHeader  `json:"header"`
}

// BlockHeader contains the blockheader response object.
type BlockHeader struct {
	PrevLedgerRoot   string              `json:"previous_ledger_root"`
	TransactionsRoot string              `json:"transactions_root"`
	Proof            string              `json:"string"`
	Metadata         BlockHeaderMetadata `json:"metadata"`
}

// BlockHeaderMetadata contains the metadata in a blockheader.
type BlockHeaderMetadata struct {
	DifficultyTarget big.Int `json:"difficulty_target"`
	Height           int64   `json:"height"`
	Nonce            string  `json:"nonce"`
	Timestamp        int64   `json:"timestamp"`
}

// GetTransactionResponse contains the transaction response.
type GetTransactionResponse struct {
	Transaction Transaction         `json:"transaction"`
	Metadata    TransactionMetadata `json:"metadata"`
}

// TransactionMetadata contains the metadata for a transaction object.
type TransactionMetadata struct {
	BlockHash        string `json:"block_hash"`
	BlockHeight      int64  `json:"block_height"`
	BlockTimestamp   int64  `json:"block_timestamp"`
	TransactionIndex int64  `json:"transaction_index"`
}

// Transaction is the transaction object.
type Transaction struct {
	TxID           string       `json:"transaction_id"`
	LedgerRoot     string       `json:"ledger_root"`
	InnerCircuitID string       `json:"inner_circuit_id"`
	Transitions    []Transition `json:"transitions"`
}

// Transition is a transition object.
type Transition struct {
	CiphertextIDs []string `json:"ciphertext_ids"`
	Ciphertexts   []string `json:"ciphertexts"`
	Commitments   []string `json:"commitments"`
	Proof         string   `json:"proof"`
	SerialNumbers []string `json:"serial_numbers"`
	ID            string   `json:"transition_id"`
	ValueBalance  int64    `json:"value_balance"`
}
