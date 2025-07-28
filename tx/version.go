package tx

type Version struct {
	value uint
}

func NewVersion(v uint) Version {
	return Version{v}
}
