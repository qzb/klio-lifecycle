# klio-lifecycle

This branch is a testing ground for interpreters.

| Name       | Compressed | Uncompressed |
| ---------- | ---------- | ------------ |
| goja       | 3.19MB     | 7.63MB       |
| gopher-lua | 1.13MB     | 2.47MB       |
| otto       | 1.21MB     | 2.81MB       |
| tengo      | 0.36MB     | 0.86MB       |

## Notes

### Otto

- Substantially harder to work with than goja.
- No support for ES2015 or later.
- Errors tend to be a little bit cryptic.

### Tengo

- Nobody knows it, documentation could be better.
- Has some issues with converting Golang types to Tengo, but can be circumvented by
  marshalling/unmarshalling.

### Lua

- Using userdata as a table froze command without any errors
- Awful API
