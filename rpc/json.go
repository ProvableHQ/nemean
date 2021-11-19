package rpc

import "encoding/json"

type Request struct {
	Method     string            `json:"method"`
	Params     []json.RawMessage `json:"params"`
	Id         string            `json:"id"`
	RpcVersion string            `json:"jsonrpc"`
}

type Result struct {
	Result json.RawMessage `json:"result"`
	Error  string          `json:"error"`
	ID     string          `json:"id"`
}

type GetBlockResponse struct {
	BlockHash         string `json:"block_hash"`
	PreviousBlockHash string `json:"previous_block_hash"`
	//Confirmations          int64    `json:"confirmations"`
	//DifficultyTarget       int64    `json:"difficulty_target"`
	//Hash                   string   `json:"hash"`
	//Height                 int64    `json:"height"`
	//MerkleRoot             int64    `json:"merkle_root"`
	//Nonce                  int64    `json:"nonce"`
	//PedersenMerkleRootHash int64    `json:"pedersen_merkle_root_hash"`
	//PreviousBlockHash      string   `json:"previous_block_hash"`
	//Proof                  string   `json:"string"`
	//Size                   int64    `json:"size"`
	//Time                   int64    `json:"time"`
	//Transactions           []string `json:"transactions"`
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

type GetTransactionInfoResponse struct {
	TxId                string      `json:"txid"`
	Size                int         `json:"size"`
	OldSerialNumbers    []int       `json:"old_serial_numbers"`
	NewCommitments      []string    `json:"new_commitments"`
	Memo                string      `json:"memo"`
	NetworkID           int         `json:"network_id"`
	Digest              string      `json:"digest"`
	TransactionProof    string      `json:"transaction_proof"`
	ProgramCommitment   string      `json:"program_commitment"`
	LocalDataRoot       string      `json:"local_data_root"`
	ValueBalance        int         `json:"value_balance"`
	Signatures          []string    `json:"signatures"`
	EncryptedRecords    []string    `json:"encrypted_records"`
	TransactionMetadata interface{} `json:"transaction_metadata"`
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
