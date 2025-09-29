#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROTO_DIR="${ROOT_DIR}/proto"
OUT_DIR="${ROOT_DIR}/gen/go"

if ! command -v buf >/dev/null 2>&1; then
  echo "buf CLI is required. Install it from https://buf.build/docs/installation." >&2
  exit 1
fi

mkdir -p "${OUT_DIR}"

buf --config "${PROTO_DIR}/buf.yaml" generate --template "${PROTO_DIR}/buf.gen.yaml" "${PROTO_DIR}"
