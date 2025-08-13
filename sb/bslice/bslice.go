package bslice

import (
	"crypto/sha256"
	"encoding/hex"
	"gopherbitcoin/sb/ba32"
)

type bs []byte

type BSlice struct {
	bs
}

func New(b []byte) BSlice {
	return BSlice{b}
}

func Zero() BSlice {
	return BSlice{}
}

func ParseHex(s string) (BSlice, error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		return Zero(), err
	}
	return New(b), nil
}

func MustParseHex(s string) BSlice {
	out, err := ParseHex(s)
	if err != nil {
		panic(err)
	}
	return out
}

func (b BSlice) Bytes() []byte {
	return b.bs
}

func (b BSlice) Size() int {
	return len(b.bs)
}

func (b BSlice) SHA256() ba32.BA32 {
	return ba32.New(sha256.Sum256(b.bs))
}

func (b BSlice) Reverse() BSlice {
	l := len(b.bs)
	reversed := make([]byte, l)
	for i := 0; i < l; i++ {
		reversed[i] = b.bs[l-1-i]
	}
	return New(reversed)
}

func (b BSlice) HexString() string {
	return hex.EncodeToString(b.bs)
}

func (b BSlice) Equal(other BSlice) bool {
	if len(b.bs) != len(other.bs) {
		return false
	}

	for i := range b.bs {
		if b.bs[i] != other.bs[i] {
			return false
		}
	}
	return true
}
