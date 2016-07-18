package utils

import (
	"testing"
)

var base62TestCases = map[int64]string{
	10: `a`,
	42: `G`,
	1000: `g8`,
	10000: `2Bi`,
	65536: `h32`,
	10000000000000: `2Q3rKTOE`,
}

func TestEncodeBase62(t *testing.T) {
	for i, a := range base62TestCases {
		r := EncodeBase62(i)
		if a != r {
			t.Errorf("wrong result. got=%d, want=%s", r, a)
		}
	}
}

func TestDecodeBase62(t *testing.T) {
	for i, a := range base62TestCases {
		r := DecodeBase62(a)
		if i != r {
			t.Errorf("wrong result. got=%s, want=%d", r, i)
		}
	}
}