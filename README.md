# klio-lifecycle

Build, test and deploy your application based on declarative configs.

## Build

To build you need to install following dependencies:

- go (v1.16)
- node (v1.16)

Additionally you have to install node dependencies and generate schema validators:

```sh
npm ci
go generate ./...
```

Finally, you can build commands:

```sh
go build ./cmd/build
go build ./cmd/deploy
```

## Docs

To build docs you need to install following dependencies:

- hugo (v0.92)
- node (v1.16)

First, you have to download theme and install node dependencies:

```sh
git submodule update --init
npm ci
```

Next, generate schemas for examples:

```sh
./scripts/generate-schemas-for-docs.sh
```

Finally, start local documentation server:

```sh
cd ./assets/docs
hugo server -D
```
