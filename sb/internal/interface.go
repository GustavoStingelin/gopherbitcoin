package internal

import (
	"gopherbitcoin/sb/ba32"
	"gopherbitcoin/sb/bslice"
)

// iBytes is an internal interface just to check the implementations, and won't
// be exported to avoid usage and interface boxing.
type iBytes[SubType, SuperType any] interface {
	Bytes() SubType
	SHA256() ba32.BA32
	Reverse() SuperType
	HexString() string
	Size() int
	Equal(other SuperType) bool
}

var _ iBytes[[32]byte, ba32.BA32] = new(ba32.BA32)

var _ iBytes[[]byte, bslice.BSlice] = new(bslice.BSlice)
