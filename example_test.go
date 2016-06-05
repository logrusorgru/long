//
// Copyright (c) 2016 Konstanin Ivanov <kostyarin.ivanov@gmail.com>.
// All rights reserved. This program is free software. It comes without
// any warranty, to the extent permitted by applicable law. You can
// redistribute it and/or modify it under the terms of the Do What
// The Fuck You Want To Public License, Version 2, as published by
// Sam Hocevar. See LICENSE.md file for more details or see below.
//

//
//        DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
//                    Version 2, December 2004
//
// Copyright (C) 2004 Sam Hocevar <sam@hocevar.net>
//
// Everyone is permitted to copy and distribute verbatim or modified
// copies of this license document, and changing it is allowed as long
// as the name is changed.
//
//            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
//   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION
//
//  0. You just DO WHAT THE FUCK YOU WANT TO.
//

package long

import (
	"fmt"
)

func ExampleEncode() {
	u := ^uint64(0)
	p := make([]byte, 10)
	n, err := Encode(u, p)
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
	n, err := Encode(u, p)
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
	n, err := Encode(uint64(i), p)
	if err != nil {
		// handle error
	}
	// decode it back
	u, _, err := Decode(p[:n])
	if err != nil {
		// handle error
	}
	fmt.Println(int(u) == i)
	// Output: true
}
