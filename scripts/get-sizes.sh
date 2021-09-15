#!/bin/sh

set -e

TEMP_DIR=$(mktemp -d)
trap "rm -rf \"$TEMP_DIR\"" EXIT

for DIR in cmd/*; do
  NAME=$(basename "$DIR")
  mkdir "$TEMP_DIR/$NAME"
  go build -o "$TEMP_DIR/$NAME/bin" "$DIR/"*
  (
    cd "$TEMP_DIR/$NAME"
    tar -czvf tgz bin 2> /dev/null
  )
done

echo "Extra size of commands based on used interpreter"
echo
echo "| Name | Compressed | Uncompressed |"
echo "| - | - | - |"

BASE_TGZ_SIZE=$(stat -f%z "$TEMP_DIR/no-interpreter/tgz")
BASE_BIN_SIZE=$(stat -f%z "$TEMP_DIR/no-interpreter/bin")

for DIR in "$TEMP_DIR"/*; do
  NAME=$(basename "$DIR")
  if [ $NAME = "no-interpreter" ]; then
    continue
  fi
  TGZ_SIZE=$(stat -f%z "$DIR/tgz")
  BIN_SIZE=$(stat -f%z "$DIR/bin")
  TGZ_SIZE_DIFF=$(python -c "print(round(($TGZ_SIZE - $BASE_TGZ_SIZE) / 1000000., 2))")
  BIN_SIZE_DIFF=$(python -c "print(round(($BIN_SIZE - $BASE_BIN_SIZE) / 1000000., 2))")
  echo "| $NAME | ${TGZ_SIZE_DIFF}MB | ${BIN_SIZE_DIFF}MB | "
done


