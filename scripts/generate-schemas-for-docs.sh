#!/usr/bin/env sh

set -e

REPO_DIR="$(git rev-parse --show-toplevel)"
STATIC_DIR="$REPO_DIR/assets/docs/static"

# Install modules
if [ ! -d "$REPO_DIR/node_modules" ]; then
  npm ci
fi

# Generate json files
cd "$REPO_DIR"
for FILEPATH in assets/schemas/public/*/*/*.yaml; do
  API_VERSION="$(echo "$FILEPATH" | cut -d / -f 4-5)"
  KIND=$(basename -s ".yaml" "$FILEPATH")

  mkdir -p "$STATIC_DIR/schemas/$API_VERSION"
  "$REPO_DIR/scripts/bundle-schema.mjs" $FILEPATH > "$STATIC_DIR/schemas/$API_VERSION/$KIND.json"
done
