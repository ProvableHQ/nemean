package main

import "github.com/urfave/cli"

var newAccountCommand = cli.Command{
	Name:     "create",
	Category: "wallet",
	Usage:    "Create a new Aleo account.",
	Description: `
	The create command is used to create a new Aleo account.
	`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "from",
			Usage:    "base64 encoded 32 byte seed",
			Required: false,
		},
	},
	Action: newAccount,
}

var fromAccountCommand = cli.Command{
	Name:     "account",
	Category: "wallet",
	Usage:    "View account address using a private key.",
	Description: `
	The account command is used to view the address of an account using a private key.
	`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "from",
			Usage:    "base32m encoded private key",
			Required: true,
		},
	},
	Action: fromAccount,
}

var newTransactionCommand = cli.Command{
	Name:     "send",
	Category: "wallet",
	Usage:    "Create a basic transfer transaction.",
	Description: `
	The send command creates a single transfer transaction that consumes
	a single record and returns a serialized transaction in hex.
	`,
	Action: newTransaction,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "to",
			Usage:    "recipient address",
			Required: true,
		},
		cli.StringSliceFlag{
			Name:     "ledger_proof",
			Usage:    "list of ledger proofs for input record",
			Required: true,
		},
		cli.Int64Flag{
			Name:     "amount",
			Usage:    "amount to send",
			Required: true,
		},
		cli.Int64Flag{
			Name:     "fee",
			Usage:    "network fee",
			Required: true,
		},
		cli.StringFlag{
			Name:     "private_key",
			Usage:    "private key to sign transaction",
			Required: true,
		},
		cli.StringFlag{
			Name:     "record",
			Usage:    "the ciphertext of the record to consume",
			Required: true,
		},
	},
}

var getBlockCommand = cli.Command{
	Name:     "getblock",
	Category: "rpc",
	Usage:    "Get a block.",
	Description: `
	Gets a block from SnarkOS.
	`,
	Flags: []cli.Flag{
		cli.Int64Flag{
			Name:     "height",
			Usage:    "block height",
			Required: true,
		},
	},
	Action: getBlock,
}

var getBlockHashCommand = cli.Command{
	Name:     "getblockhash",
	Category: "rpc",
	Usage:    "Get a block hash.",
	Description: `
	Gets a block hash from SnarkOS.
	`,
	Flags: []cli.Flag{
		cli.Int64Flag{
			Name:     "height",
			Usage:    "block height",
			Required: true,
		},
	},
	Action: getBlockHash,
}

var getBlockHashesCommand = cli.Command{
	Name:     "getblockhashes",
	Category: "rpc",
	Usage:    "Get a block hashes.",
	Description: `
	Gets block hashes from SnarkOS.
	`,
	Flags: []cli.Flag{
		cli.Int64Flag{
			Name:     "start",
			Usage:    "start block height",
			Required: true,
		},
		cli.Int64Flag{
			Name:     "end",
			Usage:    "end block height",
			Required: true,
		},
	},
	Action: getBlockHashes,
}

var getBlockHeaderCommand = cli.Command{
	Name:     "getblockheader",
	Category: "rpc",
	Usage:    "Get a block header.",
	Description: `
	Gets a block header from SnarkOS.
	`,
	Flags: []cli.Flag{
		cli.Int64Flag{
			Name:     "height",
			Usage:    "block height",
			Required: true,
		},
	},
	Action: getBlockHeader,
}

var getTransactionCommand = cli.Command{
	Name:     "gettransaction",
	Category: "rpc",
	Usage:    "Get a transaction.",
	Description: `
	Gets a transaction from SnarkOS.
	`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "id",
			Usage:    "transaction_id",
			Required: true,
		},
	},
	Action: getTransaction,
}

var getTransitionCommand = cli.Command{
	Name:     "gettransition",
	Category: "rpc",
	Usage:    "Get a transition.",
	Description: `
	Gets a transition from SnarkOS.
	`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "id",
			Usage:    "transition_id",
			Required: true,
		},
	},
	Action: getTransition,
}

var getBlockHeightCommand = cli.Command{
	Name:     "getblockheight",
	Category: "rpc",
	Usage:    "Get a block height.",
	Description: `
	Gets a block height from SnarkOS.
	`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "hash",
			Usage:    "block hash",
			Required: true,
		},
	},
	Action: getBlockHeight,
}

var getBlocksCommand = cli.Command{
	Name:     "getblocks",
	Category: "rpc",
	Usage:    "Gets blocks from start to end heights.",
	Description: `
	Gets blocks from SnarkOS.
	`,
	Flags: []cli.Flag{
		cli.Int64Flag{
			Name:     "start",
			Usage:    "start height",
			Required: true,
		},
		cli.Int64Flag{
			Name:     "end",
			Usage:    "end height",
			Required: true,
		},
	},
	Action: getBlocks,
}

var getBlocksTransactionsCommand = cli.Command{
	Name:     "getblocktransactions",
	Category: "rpc",
	Usage:    "Gets the transactions from the block of the given block height.",
	Description: `
	Gets the transactions from the block of the given block height from SnarkOS.
	`,
	Flags: []cli.Flag{
		cli.Int64Flag{
			Name:     "height",
			Usage:    "block height",
			Required: true,
		},
	},
	Action: getBlockTransactions,
}

var getCiphertextCommand = cli.Command{
	Name:     "getciphertext",
	Category: "rpc",
	Usage:    "Gets a ciphertext given the ciphertext ID.",
	Description: `
	Gets a ciphertext given the ciphertext ID from SnarkOS.
	`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "id",
			Usage:    "ciphertext id",
			Required: true,
		},
	},
	Action: getCiphertext,
}

var latestBlockCommand = cli.Command{
	Name:     "latestblock",
	Category: "rpc",
	Usage:    "Get the latest block.",
	Description: `
	Gets the latest block from SnarkOS.
	`,
	Action: latestBlock,
}

var latestBlockHeaderCommand = cli.Command{
	Name:     "latestblockheader",
	Category: "rpc",
	Usage:    "Get the latest block header.",
	Description: `
	Gets the latest block header from SnarkOS.
	`,
	Action: latestBlockHeader,
}

var latestBlockTransactionsCommand = cli.Command{
	Name:     "latestblocktransactions",
	Category: "rpc",
	Usage:    "Get the latest block transactions.",
	Description: `
	Gets the latest block transactions from SnarkOS.
	`,
	Action: latestBlockTransactions,
}

var latestBlockHeightCommand = cli.Command{
	Name:     "latestblockheight",
	Category: "rpc",
	Usage:    "Get the latest block height.",
	Description: `
	Gets the latest block height from SnarkOS.
	`,
	Action: latestBlockHeight,
}

var sendTransactionCommand = cli.Command{
	Name:     "sendtransaction",
	Category: "rpc",
	Usage:    "Broadcasts a transaction.",
	Description: `
	Broadcasts a transaction using SnarkOS.
	`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "txn",
			Usage:    "transaction hex",
			Required: true,
		},
	},
	Action: sendTransaction,
}

var latestLedgerRootCommand = cli.Command{
	Name:     "latestledgerroot",
	Category: "rpc",
	Usage:    "Gets the latest ledger root.",
	Description: `
	Returns the latest ledger root from the canonical chain.
	`,
	Action: latestLedgerRoot,
}

var getLedgerProofCommand = cli.Command{
	Name:     "getledgerproof",
	Category: "rpc",
	Usage:    "Gets the ledger proof.",
	Description: `
	Returns the ledger proof for the given commitment with the current ledger root.
	`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "commitment",
			Usage:    "The record commitment to generate a ledger proof of inclusion for.",
			Required: true,
		},
	},
	Action: getLedgerProof,
}

var decryptRecordCommand = cli.Command{
	Name:     "decrypt_record",
	Category: "wallet",
	Usage:    "Decrypts a record.",
	Description: `
	Decrypts a record using a view key.
	`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "ciphertext",
			Usage:    "The ciphertext of the record.",
			Required: true,
		},
		cli.StringFlag{
			Name:     "viewkey",
			Usage:    "The view key that can decrypt the record.",
			Required: true,
		},
	},
	Action: decryptRecord,
}

var encryptRecordCommand = cli.Command{
	Name:     "encrypt_record",
	Category: "wallet",
	Usage:    "Encrypts a record.",
	Description: `
	Encrypts a record. Returns the ciphertext of the record.
	`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "record",
			Usage:    "JSON record",
			Required: true,
		},
	},
	Action: encryptRecord,
}

var newRecordCommand = cli.Command{
	Name:     "new_record",
	Category: "wallet",
	Usage:    "Creates a new record.",
	Description: `
	Creates a new record.
	`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "owner",
			Usage:    "The owner of the record.",
			Required: true,
		},
		cli.StringFlag{
			Name:     "payload",
			Usage:    "The payload of the record.",
			Required: true,
		},
		cli.Int64Flag{
			Name:     "value",
			Usage:    "The value of the record.",
			Required: true,
		},
	},
	Action: newRecord,
}
