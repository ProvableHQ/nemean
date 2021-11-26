package rpc

// A list of supported RPC methods.
const (
	getBestBlockHashMethod        = "getbestblockhash"
	getBlockMethod                = "getblock"
	getBlockHeightMethod          = "getblockheight"
	getBlockHashMethod            = "getblockhash"
	getTransactionMethod          = "gettransaction"
	getTransitionMethod           = "gettransition"
	sendTransactionMethod         = "sendtransaction"
	latestLedgerRootMethod        = "latestledgerroot"
	getLedgerProofMethod          = "getledgerproof"
	getCiphertextMethod           = "getciphertext"
	getBlockHashesMethod          = "getblockhashes"
	getBlockHeaderMethod          = "getblockheader"
	getBlocksMethod               = "getblocks"
	getBlockTransactionsMethod    = "getblocktransactions"
	latestBlockMethod             = "latestblock"
	latestBlockHashMethod         = "latestblockhash"
	latestBlockHeaderMethod       = "latestblockheader"
	latestBlockHeightMethod       = "latestblockheight"
	latestBlockTransactionsMethod = "latestblocktransactions"
	getConnectedPeersMethod       = "getconnectedpeers"
)
