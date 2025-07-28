package bslice

import (
	"encoding/hex"
	"gopherbitcoin/sb"
	"gopherbitcoin/sb/internal"
)

func New(b []byte) sb.BSlice {
	return sb.New(b)
}

func Zero() sb.BSlice {
	return sb.Zero[internal.ByteSlice]()
}

func ParseHex(s string) (sb.BSlice, error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		return Zero(), err
	}
	return New(b), nil
}

func MustParseHex(s string) sb.BSlice {
	out, err := ParseHex(s)
	if err != nil {
		panic(err)
	}
	return out
}
