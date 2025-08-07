package tx

type Tx struct {
	version  Version
	lockTime LockTime

	// the fields aren't follow the order that they appear in a raw transaction
	// to efficiently use the memory block layout
	inputs  byte // mocked for now, TODO: implement inputs
	outputs byte // mocked for now, TODO: implement outputs
}
