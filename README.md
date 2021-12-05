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

### Docker

    docker build -t nemean -f Dockerfile .

    # create an account
    docker run  nemean /bin/bash -c "./nemean create"

### Examples
    # create a new aleo account
    ./nemean create

    # get the latest blockheight
    ./nemean -rpchost=127.0.0.1:3035 latestblockheight

## Documentation

Documentation can be found in `/doc`.

* [Overview of receiving and sending.](doc/getting_started.md)
* [Guide for wallet engineers.](doc/integration.md)
* [Setting up Nemean on an airgapped machine.](doc/airgapped.md)
* [Taking advantage of the Aleo network.](doc/uses.md)

This software is in active development. Do not use for production.
