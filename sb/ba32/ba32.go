package ba32

import (
	"crypto/sha256"
	"encoding/hex"
	"gopherbitcoin/sb"
)

type b32 [32]byte

type BA32 struct {
	b32
}

func New(b [32]byte) BA32 {
	return BA32{b}
}

func Zero() BA32 {
	return BA32{}
}

func Parse(b []byte) (BA32, error) {
	if len(b) != 32 {
		return Zero(), sb.ErrInvalidLength
	}
	return New(b32(b)), nil
}

func MustParse(b []byte) BA32 {
	out, err := Parse(b)
	if err != nil {
		panic(err)
	}
	return out
}

func ParseHex(s string) (BA32, error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		return Zero(), err
	}
	return Parse(b)
}

func MustParseHex(s string) BA32 {
	out, err := ParseHex(s)
	if err != nil {
		panic(err)
	}
	return out
}

func (b BA32) Bytes() [32]byte {
	return b.b32
}

func (b BA32) SHA256() BA32 {
	return BA32{sha256.Sum256(b.b32[:])}
}

func (b BA32) Reverse() BA32 {
	var reversed [32]byte
	for i := 0; i < 32; i++ {
		reversed[i] = b.b32[31-i]
	}
	return New(reversed)
}

func (b BA32) Size() int {
	return 32
}

func (b BA32) HexEncode() [64]byte {
	// Stack-allocated buffer - minimal heap escape
	var buf [64]byte // 32 * 2 for hex encoding
	hex.Encode(buf[:], b.b32[:])
	return buf
}

func (b BA32) HexString() string {
	buf := b.HexEncode()
	return string(buf[:])
}

func (b BA32) String() string {
	return string(b.b32[:])
}

func (b BA32) Equal(other BA32) bool {
	return b.b32 == other.b32
}
