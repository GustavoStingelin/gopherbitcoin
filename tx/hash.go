package tx

import (
	"gopherbitcoin/sb/ba32"
)

type Hash struct {
	ba32.BA32
}

func NewHash(b [32]byte) Hash {
	return Hash{ba32.New(b)}
}

func ParseHash(b []byte) (Hash, error) {
	b32, err := ba32.Parse(b)
	return Hash{b32}, err
}

func MustParseHash(b []byte) Hash {
	return Hash{ba32.MustParse(b)}
}

func ParseHashHex(s string) (Hash, error) {
	b32, err := ba32.ParseHex(s)
	return Hash{b32}, err
}

func MustParseHashHex(s string) Hash {
	return Hash{ba32.MustParseHex(s)}
}

func (h Hash) ToID() ID {
	return ID{BA32: h.Reverse()}
}
