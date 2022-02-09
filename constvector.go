package circuit

func NewConstVector(b []Bit) *BitVector {
	ret := &BitVector{make([]*Node, len(b))}
	for i, v := range b {
		n := Low
		if v {
			n = High
		}
		ret.Nodes[i] = n
	}
	return ret
}
