package sb

import (
	"errors"
	"gopherbitcoin/sb/internal"
)

var ErrInvalidLength = errors.New("invalid length")

type Bytes[T internal.SupportedBytes] struct {
	value T
}

// BSlice is a byteSlice of bytes, should be used carefully because slices are
// pointers to arrays, and pointers need garbage collection.
type BSlice = Bytes[internal.ByteSlice]

// BA32 is an array of 32 bytes.
type BA32 = Bytes[internal.ByteArray32]

// BA64 is an array of 64 bytes.
// Just by example, don't know yet if it is useful or not.
type BA64 = Bytes[internal.ByteArray64]
