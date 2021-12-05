# Integrating Nemean

Nemean is both a CLI and library. The main user of the library is for a wallet engineer that wants to integrate Aleo into their stack using Go. Alternatively, one might use Nemean as a reference for understanding how the wallet concepts fit together. 

## Architecture
Nemean is written in Go, but uses Rust for the underlying cryptographic library functionality. To do so, the snarkvm Rust crates are exposed using Rust's Foreign Function Interface (FFI). The Rust code is exposed with a C interface found in `aleo/`. See `aleo/aleo.h` for the interface. The end result is a `libaleo.so` shared library that is used by Go (Nemean). 

Nemean communciates with Rust using CGO and the shared library. See the `make build` job for a more detailed understanding.

To use this in your environment, you will need to create the shared library and have a c header file on your host machine. You can alternatively pass the ldflags to `go build` for granularity.

## Building on top of Nemean
Nemean provides types for basic wallet concepts for the Aleo network. This includes transactions, records, and accounts. To support receiving Aleo tokens, the `account/` dir is the starting point. To support sending transactions, see `record/` and `transaction`. 

Outside of send & receive, Nemean includes RPC parity with SnarkOS. With `rpc/`, you can build a service to ingest data from the network and build indexers, event producers, and other useful data tools to help query for chain and network state.

## Custody
The Nemean CLI does not use a BIP-32 (HD Wallet) scheme, and instead is a JBOK implementation. Nemean provides creating keys using Go's `crypto/rand` or by providing your own uniformly random byte slice.