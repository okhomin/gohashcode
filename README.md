# gohashcode [![GoDoc](https://godoc.org/github.com/okhomin/go-hashcode?status.svg)](https://godoc.org/github.com/okhomin/go-hashcode)

gohashcode is a Go library for generating a unique hash value for arbitrary values in Go.

This can be used to key values in a hash (for use in a map, set, etc.)
that are complex. The most common use case is comparing two values without particular fields, caching values,
de-duplication, and so on.

## Features

* Hash any arbitrary Go value, including complex types.

* Tag a struct field to ignore it and not affect the hash value.

* Optionally, override the hashing process by implementing `Hashcoder` to optimize for speed, collision
  avoidance for your data set, etc.

## Installation
```
$ go get github.com/okhomin/gohashcode
```
## Usage & Example

For usage and examples see the [Godoc](http://godoc.org/github.com/okhomin/gohashcode)

A quick code example is shown below:

```go
type S struct {
	A int
	B string
	C string `hash:"-"`     // ignore
	E string `hash:"false"` // ignore
	F []uint
	G map[string]uint
}

hash := gohashcode.Hashcode(S{
	A: 123,
	B: "Hello",
	C: "Ignore",
	E: "Ignore",
	F: []uint{1, 2, 3},
	G: map[string]uint{
		"one": 1,
		"two": 2,
	},
})

fmt.Println(hash)
// Output:
// 2377125169852
```
