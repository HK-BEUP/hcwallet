# Vexon Wallet

Vexon Wallet is the reference wallet daemon for Vexon (VEX). It is derived from
the Hcash wallet codebase and is being adapted to work with Vexon Core.

The wallet manages encrypted private keys, derives Vexon addresses, connects to
`vexond`, and exposes wallet RPC services for local tools.

## Current Status

This wallet is in private development and should not be used with public funds
until the Vexon launch checklist is complete.

Completed so far:

- Vexon app data and config names.
- Mainnet and testnet Vexon RPC ports.
- Vexon address prefixes from vendored chain parameters.
- Omni disabled by default for local development builds.
- Build script for `vexonwallet`.
- Premine wallet workflow documented.

## Build

This legacy wallet is built in GOPATH mode. On this development machine the
supported Go binary is installed at `/tmp/go117/bin/go`.

```sh
cd /Users/minxiangcai/Documents/HcashOrg/hcwallet
GO=/tmp/go117/bin/go scripts/build-vexon-wallet.sh
```

The build writes:

```text
bin/vexonwallet
```

## Create A Wallet

Use a clean machine or locked-down session. The seed controls the wallet funds.

```sh
bin/vexonwallet --create --appdata "$HOME/.vexonwallet"
```

Write the seed down offline. Never paste it into chat, logs, GitHub issues, or
commits.

## Run Locally

Start `vexond` first. Then run the wallet against the local node:

```sh
bin/vexonwallet \
  --configfile "$HOME/.vexonwallet/vexonwallet.conf" \
  --appdata "$HOME/.vexonwallet" \
  --rpcconnect 127.0.0.1:18556 \
  --cafile "$HOME/.vexond/rpc.cert" \
  --noservertls
```

For detailed node and wallet startup commands, see the Vexon Core runbook:

```text
/Users/minxiangcai/Documents/HcashOrg/hcd/docs/vexon-localnet.md
```

## Premine Address

The current public premine address is:

```text
VsbyU7TV7FeK54Jk5WkGHoYLfGu4fc78cQV
```

Only the public address belongs in source code. The wallet seed and private
passphrase must stay offline and private.

## Development Notes

More implementation notes are in [VEXON.md](VEXON.md).

## License

This code is derived from Hcash/Decred components and remains under the ISC
license used by the upstream project.
