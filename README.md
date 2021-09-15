# klio-lifecycle

Build, test and deploy your application based on declarative configs.

## Build

```sh
go generate ./...
go build cmd/build/*
```

## Docs

```
git submodule update --init
./scripts/generate-schemas-for-docs.sh
cd ./assets/docs
hugo server -D
```

## Changes/Features

- Introduced 4 new kinds of config files defining executors: Tagger, Builder, Pusher, Deployer. Each
  of those files defines schemas used for corresponding entries in a Service configuration and
  contains JS script.
- New configuration format (old one is still supported).
- Files at the same project can use different API versions.
- Config files don't have to follow naming conventions anymore.
- All config files are validated against the schemas. Additionally commands validate consistency of
  the confguration. All validations are run on the start.
- Various changes affecting placeholders:
  - Some of the placeholders were renamed.
  - Added support for recursive placeholders.
  - Placeholders are now case-insensitive.

## TODO

- Design propper API for executor scripts.
- Handle multiple versions of executor scripts API.
- Prevent using paths outside of the project directory.
- Are string arrays as executors output sufficient for all use cases?
- Make names case insensitive?
- Properly handle printing JS exceptions using "Error" instance (instead of plain string).
- Use internal schemas to validate data after transformation.
- Enable placing JS code outside of executor .yaml file (support both `js` and `jsFile`).
- Warnings about deprecated parameters passed to executors.
- Remote files:
  - Prevent using branches (and tags?) as a revision?
  - Fail if fetch requires username/password.
- Validate uniqueness of variables and params placeholders (taking into account that they are
  case-insensitive).
- Improve error messages for invalid placeholders (missing file path).
- Nice error when required --param is not specified.
- Change placeholders examples to lower-case.
- Make --environment a required parameter.
- Nice error when --tag not specified, but required.
- Enable bundling executors with command.
- Create klio-independent version? How to name it?
