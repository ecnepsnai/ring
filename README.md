# Ring

[![Go Report Card](https://goreportcard.com/badge/github.com/ecnepsnai/ring?style=flat-square)](https://goreportcard.com/report/github.com/ecnepsnai/ring)
[![Godoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/ecnepsnai/ring)
[![Releases](https://img.shields.io/github/release/ecnepsnai/ring/all.svg?style=flat-square)](https://github.com/ecnepsnai/ring/releases)
[![LICENSE](https://img.shields.io/github/license/ecnepsnai/ring.svg?style=flat-square)](https://github.com/ecnepsnai/ring/blob/master/LICENSE)

Ring is a threadsafe rotating array in Golang. You define a maximum number of items
and for each added item, it will fill the array up to that max, then begin
to wrap around the start of the array.

## Example

```golang
ring := ring.New(15)

i := 0
for i < 50 {
    ring.Add(i)
    i++
}

// All of the last values (up to the maximum) sorted by when they were added
everything := ring.All()
// The last added item
last := ring.Last()
```
