# Integrating Nemean

Nemean is both a CLI and library.The main use of the library is to make it easy for a wallet engineer that wants to integrate Aleo into their tech stack using Go. Alternatively, one might use Nemean as a reference for understanding how the wallet concepts fit together. 

## Architecture
Nemean is written in Go, but uses Rust for the underlying cryptographic library functionality. To do so, the snarkvm Rust crates are exposed using Rust's Foreign Function Interface (FFI). The Rust code is exposed with a C interface found in `aleo/`. See `aleo/aleo.h` for the interface. The end result is a `libaleo.so` shared library that is used by Go (Nemean). 

Nemean communciates with Rust using CGO and the shared library. See the `make build` job for a more detailed understanding.

To use this in your environment, you will need to create the shared library and have a c header file on your host machine. You can alternatively pass the ldflags to `go build` for granularity.

## Building on top of Nemean
Nemean provides types for basic wallet concepts for the Aleo network. This includes transactions, records, and accounts. To support receiving Aleo tokens, the `account/` dir is the starting point. To support sending transactions, see `record/` and `transaction`. 

Outside of send & receive, Nemean includes RPC parity with SnarkOS. With `rpc/`, you can build a service to ingest data from the network and build indexers, event producers, and other useful data tools to help query for chain and network state.

## Custody
The Nemean CLI does not use a BIP-32 (HD Wallet) scheme, and instead is a JOBK implementation. Nemean provides creating keys using Go's `crypto/rand` or by providing your own uniformly random byte slice.

Bring your own randomness:
```console
$ SEED=$(openssl rand -base64 32) && nemean create --from=$SEED
```

Generate accounts using a private key string:
```console
$ SECRET=$(nemean create | jq -r .privatekey) && nemean account --from=$SECRET
```

## Audits
Records are encrypted, but you can use a view key for the purposes of audits or consuming a record.

For an audit, it might be necessary to review all the records owned by your Aleo account.

First, find a record that belongs to your account:
```console
nemean -rpc=127.0.0.1:3035  gettransaction -id=at1knlnsyzxqrqwuxhe7ptp7sjumt3fvj3fkjgaelnslc8r9qj555zs8smdyn
```

Notice the ciphertexts associated with the transaction. Provide the ciphertext and viewkey associated with the account to decrypt.
```console
$ nemean decrypt_record --ciphertext="7e404cc875851b1c1b9de886767bc4f773fe8d6a13461d27f7da9b8b71907f04a38c16650c9d68a987a05727a3468f429fbff9032e20a467ec5e397d500e9806ccf66ac69f7ea6400fc3f932cc9abd86c75add6bdf5547f0a1e23b93b4ddf50bd285150f14abb2b3b3207c5d61975c1a66b4afa059a6d1ad49c476be39ef40129b13a0b2447bd835275b46e912c6428767fb10bb7d32155069e20e9162ff3c089d38bb0cafe2b727ec0dc92f1231392f412f8999fcbd927d5dd601703b49cf0ab7c483804a294be29a796b1a0cf6210a387cc5aabbe68884ccdcc3a5a6fa9b00b95ef09d1df62c89451792be91506041e64a22c0554c6d85fa4972be10c7870e24080e4489c67ee09e8789895ee39cfb80d95eecede36b394d9d225289d17a0bfdf9aeba41b4e15a38bc65330090ddd675b9ca14ae8b6a49c9d1a412e32e3409" --viewkey="AViewKey1nNE7ZmaY3gsynD8WfDGcVHpxHYmwtfzPFWKymQjuwHTm" | jq tostring
"{\"owner\":\"aleo1qnj20ajacfwf5wfs7h48zvr6gfudj92gs0ehr2z4ev24thcugyys0xegj4\",\"value\":150000000,\"payload\":\"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=\",\"program_id\":\"ap108dg24pwmezwu7hd9gt0dhrp759stge4sq4jecsg066usnclepfnhwn9a0xl5zv5spt7vvgwfqfsqt3dlw4\",\"serial_number_nonce\":\"sn1c9lz0g7nkhlsx5gtlj09d72sged0u334p6v49r5cpkkaqllvxursacerz2\",\"commitment_randomness\":\"cr1jq0cy4e56v0ch5snvzj5qqa3fga0cmk9tlgewnge5y728zxleuqssvgtex\"}"
```