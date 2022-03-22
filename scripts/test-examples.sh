#!/usr/bin/env bash

set -e
set -o pipefail

echo "Building commands..."
go build ./cmd/build
go build ./cmd/deploy

diff() {
	git diff -- ./*
}

cd examples

for DIR in *; do
	echo "Testing \"${DIR}\"..."
	cd "${DIR}"
	../../build --param param=value >build-output.txt
	../../deploy -t latest -e local --param param=value >deploy-output.txt
	if [ -n "$(diff)" ]; then
		echo
		diff
		echo
		echo "Result or output of the commands was changed. If those changes are O.K. - stage changed files."
		exit 1
	fi
	cd ..
done
