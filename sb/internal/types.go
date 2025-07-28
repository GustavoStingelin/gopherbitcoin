package internal

type ByteSlice = []byte
type ByteArray32 = [32]byte
type ByteArray64 = [64]byte

type SupportedBytes interface {
	~ByteSlice | ~ByteArray32 | ~ByteArray64
}
