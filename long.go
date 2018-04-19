//
// Copyright (c) 2016 Konstanin Ivanov <kostyarin.ivanov@gmail.com>.
// All rights reserved. This program is free software. It comes without
// any warranty, to the extent permitted by applicable law. You can
// redistribute it and/or modify it under the terms of the Do What
// The Fuck You Want To Public License, Version 2, as published by
// Sam Hocevar. See LICENSE file for more details or see below.
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

// Package long represents a method of encoding integers.
// It's similar to varint. But it optimised for negative
// numbers too.
package long

import (
	"errors"
)

const maxBytesLength = 10

var (
	// ErrShortBuffer describes that a buffer is short
	ErrShortBuffer = errors.New("short buffer")
	// ErrTooLong describes that a buffer contains bad data
	ErrTooLong = errors.New("too long")
)

// helpers

func shouldInverse(u uint64) bool { return u > ^u }
func inverse(u uint64) uint64     { return ^u }

// head helpers
func setInverseBit(c uint8) uint8 { return c | 0x2 }
func hasInverseBit(c uint8) bool  { return c&0x2 != 0 }

// body helpers
func setNextByteBit(c uint8) uint8   { return c | 0x1 }
func hasNextByteBit(c uint8) bool    { return c&0x1 != 0 }
func clearNextByteBit(c uint8) uint8 { return c & 0xfe }

// clear the last bit and compare
func nextLoop(u uint64) bool { return u&0xfffffffffffffffe > 0xfe }

func shifted(c uint8, shift uint) uint64 { return uint64(c) << shift }

// Encode uint64 to the buffer. It returns number of
// bytes and error if any. The error can only be the ErrShortBuffer
func Encode(u uint64, p []byte) (n int, err error) {
	// value length at list 1 byte
	if len(p) < 1 {
		err = ErrShortBuffer
		return
	}
	n = 1
	var c uint8 // current
	if shouldInverse(u) {
		u = inverse(u)       // inverse
		c = uint8(u)         // low byte
		c = c << 2           // the only 6 bit (clear 2 lower bits)
		c = setInverseBit(c) // set inverse bit
	} else {
		c = uint8(u) // low byte
		c = c << 2   // the only 6 bit (clear 2 lower bits)
	}
	if u <= 0x3f { // enough (6bit = 63 = 0x3f)
		p[0] = c
		return // n = 1, err = nil
	}
	// one more byte
	c = setNextByteBit(c)
	p[0] = c   // store current byte
	u = u >> 5 // shift 5 lower bit (1 lower bit is next byte bit)
	for ; nextLoop(u); n++ {
		if len(p) < n+1 {
			err = ErrShortBuffer
			return
		}
		c = uint8(u)
		c = setNextByteBit(c)
		p[n] = c
		u = u >> 7 // shift 7 lower bit
	}
	if len(p) < n+1 {
		err = ErrShortBuffer
		return
	}
	c = uint8(u)
	c = clearNextByteBit(c)
	p[n] = c
	n++
	return
}

// Decode decodes buffer and returns uint64, number of bytes,
// and error if any.
func Decode(p []byte) (u uint64, n int, err error) {
	var c uint8 // current
	if len(p) < 1 {
		err = ErrShortBuffer
		return
	}
	n = 1
	c = p[0]
	inv := hasInverseBit(c)
	defer func() {
		if inv {
			u = inverse(u)
		}
	}()
	if hasNextByteBit(c) {
		c = c >> 2 // drop the next byte bit and the inverse bit
		u = uint64(c)
		// go to the for loop
	} else {
		c = c >> 2 // drop the next byte bit and the inverse bit
		u = uint64(c)
		return // n = 1, err = nil
	}
	var shift uint = 5 // c's shift
	for {
		if n > maxBytesLength {
			err = ErrTooLong
			return
		}
		if len(p) < n+1 {
			err = ErrShortBuffer
			return
		}
		c = p[n]
		if hasNextByteBit(c) {
			c = clearNextByteBit(c)
			u = u | shifted(c, shift)
			shift = shift + 7
			n++
		} else {
			// the next byte bit already cleared
			u = u | shifted(c, shift)
			// there aren't reasons to increase the shift
			n++
			break
		}
	}
	return
}
