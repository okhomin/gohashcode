package gohashcode_test

import (
	"fmt"
	"github.com/okhomin/gohashcode"
)

func ExampleHashcode() {
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
}
