---
title: Exec
---

Package exec runs external commands.

```go
exec := import("exec")
```
## Functions

### `run(name, ...args)`

* `name` *string* – Name of the command to run.
* `...args` *string* – Arguments for the command.
* Returns: *[ExecResult][]*

Starts the specified command and waits for it to complete. If command cannot be run, or exits with non-zero status - execution of the script is aborted.

### `run_silently(name, ...args)`

* `name` *string* – Name of the command to run.
* `...args` *string* – Arguments for the command.
* Returns: *[ExecResult][]*

Starts the specified command and waits for it to complete, any output printed by command to the stdout is logged at *debug* level. If command cannot be run, or exits with non-zero status - execution of the script is aborted.

### `command(opts)`

* `opts` *map*
* `opts.name` *string* – Name of the command to run.
* `opts.args` *Array\<string>* – Arguments for the command.
* `opts.dir` *string* – Working directory of the command, by default project's directory.
* `opts.env` *Array\<string>* – List of env variables, each entry is of the form `key=value`.
* `opts.stdout_level` *string* – Logging level for logs printed to stdout.
* `opts.stderr_level` *string* – Logging level for logs printed to stderr.
* `opts.ignore_errors` *bool* – Unless set to true, if error occurs, `command.run()` aborts script execution.
* Returns: *[Command][]*

Prepares command to run.

Options `stdout_level` and `stderr_level` accept following log levels: fatal, error, warn, info, verbose, debug, spam and disable. Disable completely suppress logging - you should't use it unless you need to handle sensitive data. Providing invalid level will cause [run](#commandrun) method to fail.

## Command

Command represents an external command being prepared or run.

### `command.run()`

* Returns: *[ExecResult][]*

Run starts the specified command and waits for it to complete. Unless `ignore_errors` option is enabled, if command cannot be run, or exits with non-zero status - execution of the script is aborted.

This method can be called multiple times.

## ExecResult

### `exec_result.stdout_text`

* *string*

Text printed by the command to the stdout. It isn't affected by the `stdout_level` option.

### `exec_result.stderr_text`

* *string*

Text printed by the command to the stderr. It isn't affected by the `stderr_level` option.

### `exec_result.exit_code`

* *int*

The exit code of the command, or -1 if command couldn't be run, or was terminated by a signal. Codes other than 0 indicates error. To access this property command must be run with `ignore_errors` option enabled.
### `exec_result.error`

* *Error* | *undefined*

Error which occurred during command execution. To access this property command must be run with `ignore_errors` option enabled.

[Command]: #command
[ExecResult]: #execresult
[log levels]: test

