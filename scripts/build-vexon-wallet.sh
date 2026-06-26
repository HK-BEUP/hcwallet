#!/usr/bin/env bash
set -euo pipefail

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
out_dir="${1:-"$repo_root/bin"}"
gopath="${VEXON_GOPATH:-/tmp/vexon-gopath}"
pkg_dir="$gopath/src/github.com/HcashOrg"
wallet_link="$pkg_dir/hcwallet"

mkdir -p "$out_dir" "$pkg_dir"

if [ -e "$wallet_link" ] && [ "$(readlink "$wallet_link" || true)" != "$repo_root" ]; then
	echo "Refusing to replace existing $wallet_link" >&2
	exit 1
fi

ln -sfn "$repo_root" "$wallet_link"

go_bin="${GO:-go}"
gocache="${GOCACHE:-/tmp/vexon-gocache}"

echo "Building vexonwallet into $out_dir"
(
	cd "$wallet_link"
	GO111MODULE=off GOCACHE="$gocache" GOPATH="$gopath" "$go_bin" build -o "$out_dir/vexonwallet" .
)

"$out_dir/vexonwallet" --version || true

echo "Done."
