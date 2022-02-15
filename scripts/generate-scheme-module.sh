#!/usr/bin/env sh

set -e

REPO_DIR="$(git rev-parse --show-toplevel)"
MODULE_DIR="$REPO_DIR/internal/blueprint/internal/scheme"

# Install modules
if [ ! -d "$REPO_DIR/node_modules" ]; then
  npm ci
fi

# Generate schemas.go file
(
  cd "$REPO_DIR"
  echo 'package scheme'
  echo
  echo 'var SCHEMAS = map[string][]byte{'
  for FILEPATH in assets/schemas/public/*/*/*.yaml; do
    API_VERSION="$(echo "$FILEPATH" | cut -d / -f 4-5)"
    BASENAME=$(basename -s ".yaml" "$FILEPATH")
    KIND=`echo ${BASENAME:0:1} | tr  '[a-z]' '[A-Z]'`${BASENAME:1}

    if [ "$KIND" = Index ]; then
      continue
    fi

    echo "\t\"$API_VERSION/$KIND\": []byte(\`"
    "$REPO_DIR/scripts/normalize-schema.mjs" $FILEPATH | sed -e 's/^/\t\t/;'
    echo "\t\`),"
  done

  echo "\t\"internal/Object\": []byte(\`"
  "$REPO_DIR/scripts/normalize-schema.mjs" assets/schemas/internal/object.yaml | sed -e 's/^/\t\t/;'
  echo "\t\`),"

  echo '}'
) > "$MODULE_DIR/schemas.go"
