package rpc

// A list of supported RPC methods.
const (
	getBestBlockHashMethod         = "getbestblockhash"
	getBlockMethod                 = "getblock"
	getBlockHeightMethod           = "getblockheight"
	getBlockHashMethod             = "getblockhash"
	getBlockTemplateMethod         = "getblocktemplate"
	getConnectionCountMethod       = "getconnectioncount"
	getRawTransactionMethod        = "getrawtransaction"
	getTransactionInfoMethod       = "gettransactioninfo"
	sendTransactionMethod          = "sendtransaction"
	validateRawTransactionMethod   = "validaterawtransaction"
	createAccountMethod            = "createaccount"
	latestLedgerRootMethod         = "latestledgerroot"
	getLedgerProofMethod           = "getledgerproof"
	createRawTransactionMethod     = "createrawtransaction"
	createTransactionMethod        = "createtransaction"
	createTransactionKernelMethod  = "createtransactionkernel"
	decodeRecordMethod             = "decoderecord"
	decryptRecordMethod            = "decryptrecord"
	getRawRecordMethod             = "getrawrecord"
	getRecordCommitmentCountMethod = "getrecordcommitmentcount"
	getRecordCommitmentsMethod     = "getrecordcommitments"
)
