package tx

var MinTimestampLockTime = uint32(500_000_000)
var MaxBlockHeightLockTime = MinTimestampLockTime - 1

type LockTime struct {
	value uint32
}

func NewLockTime(v uint32) LockTime {
	return LockTime{v}
}

func (l LockTime) Value() uint32 {
	return l.value
}

func (l LockTime) IsZero() bool {
	return l.value == 0
}

func (l LockTime) IsBlockHeightLock() bool {
	return l.value < MaxBlockHeightLockTime
}

func (l LockTime) IsTimestampLock() bool {
	return l.value >= MinTimestampLockTime
}
