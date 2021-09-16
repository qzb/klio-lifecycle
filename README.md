# klio-lifecycle

This branch is a testing ground for interpreters.

| Package                                               | Stars | Watchers | Forks | Last Commit |
| ----------------------------------------------------- | ----- | -------- | ----- | ----------- |
| [otto](https://github.com/robertkrimen/otto)          | 6222  | 189      | 521   | 14 Jun 2021 |
| [gopher-lua](https://github.com/yuin/gopher-lua)      | 4289  | 149      | 477   | 29 May 2021 |
| [yaegi](https://github.com/traefik/yaegi)             | 3816  | 50       | 183   | 15 Sep 2021 |
| [tengo](https://github.com/d5/tengo)                  | 2422  | 57       | 141   | 5 Sep 2021  |
| [goja](https://github.com/dop251/goja)                | 2386  | 66       | 201   | 12 Sep 2021 |
| [go-lua](https://github.com/Shopify/go-lua)           | 2152  | 312      | 156   | 2 Mar 2021  |
| [starlark-go](https://github.com/google/starlark-go)  | 1367  | 46       | 122   | 1 Sep 2021  |
| [go-python](https://github.com/sbinet/go-python)      | 1303  | 43       | 128   | 14 Apr 2021 |
| [anko](https://github.com/mattn/anko)                 | 1167  | 47       | 111   | 21 May 2020 |
| [go-php](https://github.com/deuill/go-php)            | 42    | 811      | 92    | 1 Oct 2018  |
| [go-ducktape](https://github.com/olebedev/go-duktape) | 778   | 28       | 90    | 26 Mar 2021 |
| [golua](https://github.com/aarzilli/golua)            | 555   | 34       | 193   | 7 May 2021  |
| [gpython](https://github.com/go-python/gpython)       | 517   | 21       | 58    | 18 Nov 2019 |

Stats as at 16 Sep 2021.

## Additional size of commands

Following numbers are based on interpreters integrated with build command, compared with
interpreter-less version.

| Package    | Compressed | Uncompressed |
| ---------- | ---------- | ------------ |
| goja       | 3.19MB     | 7.63MB       |
| gopher-lua | 1.13MB     | 2.47MB       |
| otto       | 1.21MB     | 2.81MB       |
| tengo      | 0.36MB     | 0.86MB       |

## Standard libraries

| Package    | File System | Exec | HTTP | Regex | JSON | YAML | Print | Date & Time |
| ---------- | ----------- | ---- | ---- | ----- | ---- | ---- | ----- | ----------- |
| goja       | no          | no   | no   | yes   | yes  | no   | yes\* | yes         |
| gopher-lua | yes         | no   | no   | no    | no   | yes  | yes   | yes         |
| otto       | no          | no   | no   | yes   | yes  | no   | yes   | yes         |
| tengo      | yes         | yes  | no   | yes   | yes  | no   | yes   | yes         |

\* Provided by an additional library

## Notes

### goja

- Simplest to integrate.
- Has partial support for ECMA 2015+, but it's poorly documented which features are supported.

### gopher-lua

- Using userdata as a table froze command without any errors.
- Awful API.
- IO API is too barebone.

### otto

- Substantially harder to work with than goja.
- No support for ES2015 or later.
- Errors tend to be a little bit cryptic.

### gpython

- There is no documentation how to integrate it into golang app.
- Seems to be no longer actively maintained.
- Passing data into VM is not supported.
- Seems to be no way to extract data from a VM.

### tengo

- Nobody knows it, documentation could be better.
- Has some issues with converting Golang types to Tengo, but they can be circumvented by
  marshalling/unmarshalling.

### yaegi

- Full-blown Golang interpreter, not suitable for scripting.
