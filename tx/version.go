package tx

// VersionOne is the first version of the Bitcoin Transaction.
var VersionOne = NewVersion(1)

// VersionTwo was introduced with BIP 68, BIP 112, and BIP 113.
// It enables relative timelocks and CHECKSEQUENCEVERIFY.
var VersionTwo = NewVersion(2)

type Version struct {
	value uint32
}

func NewVersion(v uint32) Version {
	return Version{v}
}

func (v Version) Value() uint32 {
	return v.value
}

func (v Version) IsOne() bool {
	return v.value == 1
}

func (v Version) IsTwo() bool {
	return v.value == 2
}
