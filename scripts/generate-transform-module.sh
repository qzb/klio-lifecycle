#!/usr/bin/env sh

REPO_DIR="$(git rev-parse --show-toplevel)"
MODULE_DIR="$REPO_DIR/internal/blueprint/internal/scheme/transform"

# Install modules
if [ ! -d "$REPO_DIR/node_modules" ]; then
  npm ci
fi

# Generate types
(
  cd "$REPO_DIR"
  for DIR in assets/schemas/public/*/*; do
    API_VERSION="$(echo "$DIR" | cut -d / -f 4-5)"
    mkdir -p "$MODULE_DIR/src/types/$API_VERSION"
    "$REPO_DIR/scripts/generate-types.mjs" "$DIR/object.yaml" > "$MODULE_DIR/src/types/$API_VERSION/index.ts";
  done
  "$REPO_DIR/scripts/generate-types.mjs" "assets/schemas/internal/object.yaml" > "$MODULE_DIR/src/types/internal.ts";
)

# Compile typescript
(
  cd "$MODULE_DIR"
  "$(npm bin)/tsc"
)

# Put JS code into golang file
(
  cd "$MODULE_DIR"
  echo 'package transform'
  echo
  echo 'const SCRIPT = `'
  echo 'var exports = {};'
  cat dist/index.js
  echo '`'
) > script.go

