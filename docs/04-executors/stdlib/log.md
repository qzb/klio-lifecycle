---
title: Log
---

A simple logging package.

```go
log := import("log")
```

{{% notice note %}}
Since `error` is Tengo's keyword, calling `log.error("message")` ends with syntax error. Instead use one of the following:

- `log["error"]("message")`
- `log.err("message")`
{{% /notice %}}

## Functions

### `print(...args)`

* `...args` *any*

Formats message using the default formats for its operands and writes to standard output.

### `printf(format, ...args)`

* `format` *string*
* `...args` *any*

Formats message according to a format specifier and writes to standard output.

### `fatal(...args)`

* `...args` *any*

Same as [print][], but prints at fatal level.

### `fatalf(...args)`

* `format` *string*
* `...args` *any*

Same as [printf][], but prints at fatal level.

### `err(...args)`

* `...args` *any*

Same as [print][], but prints at error level.

### `error(...args)`

* `...args` *any*

Same as [print][], but prints at error level.

### `errorf(...args)`

* `format` *string*
* `...args` *any*

Same as [printf][], but prints at error level.

### `warn(...args)`

* `...args` *any*

Same as [print][], but prints at warn level.

### `warnf(...args)`

* `format` *string*
* `...args` *any*

Same as [printf][], but prints at warn level.

### `info(...args)`

* `...args` *any*

Same as [print][], but prints at info level.

### `infof(...args)`

* `format` *string*
* `...args` *any*

Same as [printf][], but prints at info level.

### `verbose(...args)`

* `...args` *any*

Same as [print][], but prints at verbose level.

### `verbosef(...args)`

* `format` *string*
* `...args` *any*

Same as [printf][], but prints at verbose level.

### `debug(...args)`

* `...args` *any*

Same as [print][], but prints at debug level.

### `debugf(...args)`

* `format` *string*
* `...args` *any*

Same as [printf][], but prints at debug level.

### `spam(...args)`

* `...args` *any*

Same as [print][], but prints at spam level.

### `spamf(...args)`

* `format` *string*
* `...args` *any*

Same as [printf][], but prints at spam level.

[print]: #printargs
[printf]: #printfformat-args
