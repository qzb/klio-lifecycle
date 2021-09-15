#!/bin/sh

set -e -o pipefail

TEMP_DIR=$(mktemp -d)
#trap "rm -rf \"$TEMP_DIR\"" EXIT

echo "\033[1;34mOUTPUT DIR:\033[0m $TEMP_DIR"
for DIR in cmd/*; do
  NAME=$(basename "$DIR")
  mkdir "$TEMP_DIR/$NAME"
  echo "\n\033[1;34mRUNNING $NAME\033[0m"
  go run "$DIR/"* -f "examples/$NAME/project.yaml" --result-file "$TEMP_DIR/$NAME/result.json" 2>&1 | tee "$TEMP_DIR/$NAME/logs.txt"
  diff "examples/$NAME/result.json" "$TEMP_DIR/$NAME/result.json"
done
