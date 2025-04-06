#!/usr/bin/env bash

set -e

SRC="../cf-api-spec/gen/openapi.yaml"
DST="./spec"

if [[ ! -f "$SRC" ]]; then
  echo "OpenAPI spec not found at '$SRC'"
  exit 1
fi

echo "📄 Copying OpenAPI spec from '$SRC' to '$DST'..."
cp "$SRC" "$DST"

echo "OpenAPI spec copied successfully."
echo "To generate a new server, run: ./script/generate.sh"