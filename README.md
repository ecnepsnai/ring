# Ring

[![Go Report Card](https://goreportcard.com/badge/github.com/ecnepsnai/ring?style=flat-square)](https://goreportcard.com/report/github.com/ecnepsnai/ring)
[![Godoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/ecnepsnai/ring)
[![Releases](https://img.shields.io/github/release/ecnepsnai/ring/all.svg?style=flat-square)](https://github.com/ecnepsnai/ring/releases)
[![LICENSE](https://img.shields.io/github/license/ecnepsnai/ring.svg?style=flat-square)](https://github.com/ecnepsnai/ring/blob/master/LICENSE)

Package ring is a threadsafe rotating slice. You define a maximum number of items and for each added item, it will fill
the slice up to that max, then begin to wrap around the start of the slice.

# Usage & Examples

Examples can be found on the [documentation for the library](https://pkg.go.dev/github.com/ecnepsnai/ring)
