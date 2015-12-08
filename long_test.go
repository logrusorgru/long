package long_test

import (
	"github.com/logrusorgru/long"
	"testing"
)

type rng struct {
	f, t uint64
}

var ranges = []uint64{
	0, 1, // 1, 1
	63, 64, // 1, 2
	8191, 8192, // 2, 3
	1048575, 1048576, // 3, 4
	134217727, 134217728, // 4, 5
	17179869183, 17179869184, // 5, 6
	2199023255551, 2199023255552, // 6, 7
	281474976710655, 281474976710656, // 7, 8
	36028797018963967, 36028797018963968, // 8, 9
	4611686018427387903, 4611686018427387904, // 9, 10
	18446744073709551614, 18446744073709551615, // 1, 1
	0x0fffffffffffffff, // 9
	0x1fffffffffffffff, // 9
	0x2fffffffffffffff, // 9
	0x3fffffffffffffff, // 9
	0x4fffffffffffffff, // 10
	0x5fffffffffffffff, // 10
	0x6fffffffffffffff, // 10
	0x7fffffffffffffff, // 10
	0x8fffffffffffffff, // 10
	0x9fffffffffffffff, // 10
	0xafffffffffffffff, // 10
	0xbfffffffffffffff, // 10
	0xcfffffffffffffff, // 9
	0xdfffffffffffffff, // 9
	0xefffffffffffffff, // 9
	0xffffffffffffffff, // 1
}

func isClear(p []byte, n int) bool {
	for i := n; i < len(p); i++ {
		if p[i] != 0 {
			return false
		}
	}
	return true
}

func clear(p []byte) {
	for i := range p {
		p[i] = 0
	}
}

var maxUint64 = ^uint64(0)

func TestRanges(t *testing.T) {
	p := make([]byte, 10)
	for _, u := range ranges {
		n, err := long.Encode(u, p)
		if err != nil {
			t.Error("Encode error:", err)
		}
		if n > 10 {
			t.Error("Encode: len too big")
		}
		y, m, err := long.Decode(p[:n])
		if err != nil {
			t.Error("Decode error:", err)
		}
		if y != u {
			t.Errorf("Encode/Decode: wrong value, expected %d, got %d\n", u, y)
		}
		if m > 10 {
			t.Error("Decode: len too big")
		}
		if !isClear(p, n) {
			t.Error("extra data in buffer")
		}
		clear(p)
	}
}

// errors test

func TestEncodeShortBuffer(t *testing.T) {
	u := uint64(0x7fffffffffffffff)
	for i := 0; i < 10; i++ {
		p := make([]byte, i)
		_, err := long.Encode(u, p)
		if err == nil {
			t.Error("missing error")
		}
	}
}

func TestDecodeShortBuffer(t *testing.T) {
	p := make([]byte, 10)
	u := uint64(0x7fffffffffffffff)
	long.Encode(u, p)
	for i := 0; i < 10; i++ {
		_, _, err := long.Decode(p[:i])
		if err == nil {
			t.Error("missing error")
		}
	}
}

func TestDecodeTooLong(t *testing.T) {
	p := make([]byte, 10)
	u := uint64(0x7fffffffffffffff)
	long.Encode(u, p)
	p = append(p[:5], p...)
	_, _, err := long.Decode(p)
	if err == nil {
		t.Error("missing error")
	}
}
