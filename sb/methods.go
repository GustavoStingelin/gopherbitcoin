package sb

import (
	"crypto/sha256"
	"encoding/hex"
	"gopherbitcoin/sb/internal"
	"unsafe"
)

func (b Bytes[T]) Bytes() T {
	return b.value
}

// Slice returns a byte slice representation of the Bytes[T] value.
// It escapes to heap, because slices are pointers to arrays, so use it
// carefully.
func (b Bytes[T]) Slice() []byte {
	return unsafe.Slice(
		(*byte)(unsafe.Pointer(&b.value)),
		len(b.value),
	)
}

func (b Bytes[T]) SHA256() BA32 {
	return New(sha256.Sum256(b.Slice()))
}

func (b Bytes[T]) Size() int {
	return len(b.value)
}

// Reverse returns a new Bytes[T] with the bytes in reverse order.
// This creates a new slice, so it escapes to heap, use it carefully.
func (b Bytes[T]) Reverse() Bytes[T] {
	reversed := make([]byte, len(b.value))
	for i := 0; i < len(b.value); i++ {
		reversed[i] = b.value[len(b.value)-1-i]
	}
	return Bytes[T]{T(reversed)}
}

func (b Bytes[T]) HexString() string {
	return hex.EncodeToString(b.Slice())
}

func New[T internal.SupportedBytes](value T) Bytes[T] {
	return Bytes[T]{value: value}
}

func Zero[T internal.SupportedBytes]() Bytes[T] {
	return Bytes[T]{}
}
