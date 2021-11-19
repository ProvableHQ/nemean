# Nemean
Nemean is a CLI and SDK that provides tooling for custodians and engineers to interact with the [Aleo](https://aleo.org/) network.

The library provides a Go wrapper around Aleo objects while relying on Rust for the underlying cryptography.

This includes:
* An RPC client compatible with [SnarkOS](https://github.com/AleoHQ/snarkOS)
* Basic send and receive support.

## Getting Started

The cli provides an RPC client to communicate with SnarkOS: 

    make build

    ./nemean

Nemean can also be used as a library to handle wallet management.

### Examples
    # create a new aleo account
    ./nemean create

This software is in active development. Do not use for production.
