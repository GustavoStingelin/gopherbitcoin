package ba32

import (
	"encoding/hex"
	"gopherbitcoin/sb"
	"gopherbitcoin/sb/internal"
)

func New(b [32]byte) sb.BA32 {
	return sb.New(b)
}

func Zero() sb.BA32 {
	return sb.Zero[internal.ByteArray32]()
}

func Parse(b []byte) (sb.BA32, error) {
	if len(b) != 32 {
		return Zero(), sb.ErrInvalidLength
	}
	return sb.New([32]byte(b)), nil
}

func MustParse(b []byte) sb.BA32 {
	out, err := Parse(b)
	if err != nil {
		panic(sb.ErrInvalidLength)
	}
	return out
}

func ParseHex(s string) (sb.BA32, error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		return Zero(), err
	}
	return Parse(b)
}

func MustParseHex(s string) sb.BA32 {
	out, err := ParseHex(s)
	if err != nil {
		panic(err)
	}
	return out
}
