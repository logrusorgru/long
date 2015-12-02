package long_test

import (
	"fmt"
	"github.com/logrusorgru/long"
)

func ExampleEncode() {
	u := ^uint64(0)
	p := make([]byte, 10)
	n, err := long.Encode(u, p)
	if err != nil {
		// handle error
	}
	// print result
	for i, v := range p[:n] {
		fmt.Printf("#%d: %08b\n", i, v)
	}
	// Output: #0: 00000010
}

func ExampleEncode_negative() {
	i := -500
	u := ^uint64(i)
	p := make([]byte, 10)
	n, err := long.Encode(u, p)
	if err != nil {
		// handle error
	}
	// print result
	for i, v := range p[:n] {
		fmt.Printf("#%d: %08b\n", i, v)
	}
	// Output:
	// #0: 11001101
	// #1: 00001110
}

func ExampleDecode() {
	i := -8193
	p := make([]byte, 10)
	n, err := long.Encode(uint64(i), p)
	if err != nil {
		// handle error
	}
	// decode it back
	u, _, err := long.Decode(p[:n])
	if err != nil {
		// handle error
	}
	fmt.Println(int(u) == i)
	// Output: true
}
