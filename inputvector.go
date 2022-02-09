package circuit

type InputVector struct {
	BitVector
	Bits BitArray
}

func NewInputVector(n int) *InputVector {
	bits := make([]Bit, n)
	nodes := make([]*Node, n)
	for i := range nodes {
		nodes[i] = NewNode(NewVariableGate(&bits[i]), nil, nil)
	}
	return &InputVector{BitVector{nodes}, bits}
}

func (v *InputVector) GetBytes() []byte {
	return v.Bits.GetBytes()
}

func (v *InputVector) SetBytes(arr []byte) {
	v.Bits.SetBytes(arr)
}

func (v *InputVector) GetUint64() uint64 {
	return v.Bits.GetUint64()
}

func (v *InputVector) SetUint64(val uint64) {
	v.Bits.SetUint64(val)
}
