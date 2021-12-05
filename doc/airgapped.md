# Nemean in an airgapped environment

This document describes how to use Nemean in an airgapped environment. In a production setting, private key material should never be on/next to a network attached device. Fortunately, this library and cli can be configured to be run in a non-networked environment.

## Setting up
Prepare an airgapped machine that will be used to send and receive transactions. A regular retail computer w/o networking can be used. For better security, consider dedicated hardware as well as hardened operating systems.

Clone the repo onto a data storage device and build locally on the airgapped machine. Alternatively, build the Dockerfile and move it to the airgapped machine.

To create an account, simply use the nemean cli.
```console
$ nemean create
```

The above command will use Go's `crypto/rand`. Otherwise, one might consider using a seed generated elsewhere.
```console
$ SEED=$(openssl rand -base64 32) && nemean create --from=$SEED
```

For convenience, you might want to transmit the payload using QR. For example, the following will generate a QR code for an Aleo address.
```console
$ nemean create | jq .Address | qrencode  -t utf8
█████████████████████████████████████████
█████████████████████████████████████████
████ ▄▄▄▄▄ █▀ █▀▀▀▀▀▄▄▄▄ ▀▄▄▄█ ▄▄▄▄▄ ████
████ █   █ █▀ ▄ █▀▄  ▀ ▄▀▄█▄▀█ █   █ ████
████ █▄▄▄█ █▀█ █▄█ ▀▀▀▀▄█▄ ▄▀█ █▄▄▄█ ████
████▄▄▄▄▄▄▄█▄█▄█ █▄▀▄█▄▀ ▀ █ █▄▄▄▄▄▄▄████
████ ▄ ▄▄█▄ ▄ ▄█▄▄ ▄ ▄▀▀▄▀▀ █ ▀ █ █▄▀████
████ ▀█ ▄ ▄▀  ▀ ▄▄▄▄███▀▄▀▄▀▄▄███▄▀ █████
█████▄▀ ▄█▄▄▄▄▀▄▀▄▄▄  ▀ ▀▀▄▀▄ ▀▄▄▄▀▄▀████
████▀▄██▀▄▄▄▄ ██▀█   ▄▄██ ▄█ █▀ ▀  ▄█████
████ █▄▀█ ▄ █▄ █▄▄ ▄  ▀ ▀▀▄ ▄ ▀▄▄ ▀▄ ████
██████ ▀▄ ▄ ▄█  ▄▄▄▄█▄ ██▀▀█▀█ ▄▄ ▀ █████
████▀█▄█  ▄▄▀█▀▄▀▄█▄▄▄█▀▄▀▄▀█▀▀▄▄ ▀█▀████
████ █▄▀ ▄▄▀▄ ██▀▄█▄▀▄▀█▀▀▄█ ▀▀██▄ ▄▀████
████▄██▄▄█▄▄ █▀█▄▀█▄▄▄▀ ▀█▄  ▄▄▄ ▄▀█▀████
████ ▄▄▄▄▄ █▄▀▄ ▄▄█▄ ▄▀▀ ▀▄▄ █▄█ ▀▀█▀████
████ █   █ █ ▄█▄▀▀█▄  ▀▀▄█▄█  ▄▄ ▀▀██████
████ █▄▄▄█ █   █▀ ▄ █▄▄█▀  ███ ▀▀▀ ▄█████
████▄▄▄▄▄▄▄█▄▄▄█▄█▄▄▄▄██▄██▄▄█▄▄▄██▄█████
█████████████████████████████████████████
█████████████████████████████████████████
```

## Payloads
To craft a transaction, there is several fields of stateful information that must be provided to the airgapped machine.

This includes the following information:
```console
$ nemean send
NAME:
   nemean send - Create a basic transfer transaction.

USAGE:
   nemean send [command options] [arguments...]

CATEGORY:
   wallet

DESCRIPTION:
   
  The send command creates a single transfer transaction that consumes
  a single record and returns a serialized transaction in hex.
  

OPTIONS:
   --to value            recipient address
   --ledger_proof value  list of ledger proofs for input record
   --amount value        amount to send (default: 0)
   --fee value           network fee (default: 0)
   --private_key value   private key to sign transaction
   --record value        JSON input record to consume
```

Depending on the available hardware capabilities of the machine, one might consider QR or USB. There is not yet a standard developed by the community for airgapped operations, so a custom payload structure must be considered.
