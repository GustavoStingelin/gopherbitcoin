package tx

import (
	"gopherbitcoin/sb"
	"gopherbitcoin/sb/ba32"
)

type ID struct {
	sb.BA32
}

func NewID(b [32]byte) ID {
	return ID{ba32.New(b)}
}

func ParseID(b []byte) (ID, error) {
	b32, err := ba32.Parse(b)
	return ID{b32}, err
}

func MustParseID(b []byte) ID {
	return ID{ba32.MustParse(b)}
}

func ParseIDHex(s string) (ID, error) {
	b32, err := ba32.ParseHex(s)
	return ID{b32}, err
}

func MustParseIDHex(s string) ID {
	return ID{ba32.MustParseHex(s)}
}

func (id ID) ToHash() Hash {
	return Hash{BA32: id.Reverse()}
}
