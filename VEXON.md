# Vexon wallet notes

This branch starts the `hcwallet` fork work for Vexon (VEX).

## Current changes

- Default wallet app data directory: `vexonwallet`
- Default wallet config file: `vexonwallet.conf`
- Default wallet log file: `vexonwallet.log`
- Default node certificate source: `vexond/rpc.cert`
- Mainnet node RPC port: `18556`
- Mainnet wallet JSON-RPC server port: `18557`
- Mainnet wallet gRPC server port: `18558`
- Testnet node RPC port: `28556`
- Testnet wallet JSON-RPC server port: `28557`
- Testnet wallet gRPC server port: `28558`
- Vendored `hcd` chain parameters were synced from the Vexon `hcd` branch so
  wallet address encoding uses the Vexon network magics and V address prefixes.

## Omni build switch

The upstream wallet always pulled in `omnilib`, which requires old static C++
libraries such as Boost, Berkeley DB, OpenSSL 1.0.2, and libevent. Vexon first
stage does not need Omni, so Omni is disabled by default.

Normal wallet build:

```sh
chmod +x scripts/build-vexon-wallet.sh
GO=/tmp/go117/bin/go scripts/build-vexon-wallet.sh
```

Omni-enabled build, only after the required C++ static libraries are installed:

```sh
GO111MODULE=off go build -tags omni .
```

## Verification

Using Go 1.17.13 in a temporary GOPATH, the wallet main package builds without
the Omni tag.

Address encoding was verified against the vendored Vexon `hcd` params:

```text
P2PKH starts with Vs
P2SH starts with Vc
```

Do not use a randomly generated verification address for premine funds. Generate
the final premine address from a backed-up Vexon wallet seed.

## Premine wallet steps

Use a clean machine or a locked-down server session for the premine wallet. The
seed controls the planned 8,400,000 VEX premine.

```sh
cd /Users/minxiangcai/Documents/HcashOrg/hcwallet
chmod +x scripts/build-vexon-wallet.sh
GO=/tmp/go117/bin/go scripts/build-vexon-wallet.sh

bin/vexonwallet --create --appdata /secure/path/vexonwallet
```

After the seed is written down and the node is running, start the wallet against
local `vexond` and ask it for a receiving address:

```sh
bin/vexonwallet --appdata /secure/path/vexonwallet
/Users/minxiangcai/Documents/HcashOrg/hcd/bin/vexonctl --wallet getnewaddress
```

Only the returned public `Vs...` address should be copied into the Vexon node
premine ledger. Never put wallet files, seed words, TLS keys, or RPC passwords
in Git.
