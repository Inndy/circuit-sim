package circuit

type BitVector struct {
	Nodes []*Node
}

func NewBitVector(n int) *BitVector {
	return &BitVector{make([]*Node, n)}
}

func FillVector(n int, node *Node) *BitVector {
	r := NewBitVector(n)
	for i := range(r.Nodes) {
		r.Nodes[i] = node
	}
	return r
}

func ZeroVector(n int) *BitVector {
	return FillVector(n, Low)
}

func OneVector(n int) *BitVector {
	return FillVector(n, High)
}

func (v *BitVector) optimize() {
	for i, n := range v.Nodes {
		v.Nodes[i] = Optimize(n)
	}
}

func (v *BitVector) Optimize() int {
	begin := EliminatedNodeCount
	e := EliminatedNodeCount + 1
	for e != EliminatedNodeCount {
		e = EliminatedNodeCount
		v.optimize()
	}
	return EliminatedNodeCount - begin
}

func (v *BitVector) makeUnaryOperate(f UnaryGateMaker) *BitVector {
	ret := NewBitVector(len(v.Nodes))
	for i := range(ret.Nodes) {
		ret.Nodes[i] = f(v.Nodes[i])
	}

	return ret
}

func (v *BitVector) makeBinaryOperate(v2 *BitVector, f BinaryGateMaker) *BitVector {
	n := len(v.Nodes)
	if n > len(v2.Nodes) {
		n = len(v2.Nodes)
	}

	ret := NewBitVector(n)
	for i := len(ret.Nodes) - 1; i >= 0; i-- {
		ret.Nodes[i] = f(v.Nodes[i], v2.Nodes[i])
	}

	return ret
}

func (v *BitVector) Copy() *BitVector {
	nodes := make([]*Node, len(v.Nodes))
	copy(nodes, v.Nodes)
	return &BitVector{Nodes: nodes}
}

func (v *BitVector) Not() *BitVector {
	return v.makeUnaryOperate(func(n *Node) *Node {
		return Not(n)
	})
}

func (v *BitVector) RotateRight(n int) *BitVector {
	r := v.Copy()
	l := len(v.Nodes)

	for n < 0 {
		n += l;
	}
	for n >= l {
		n -= l;
	}

	tail := make([]*Node, n)
	copy(tail, r.Nodes[l-n:])
	copy(r.Nodes[n:], r.Nodes)
	copy(r.Nodes, tail)
	return r
}

func (v *BitVector) RotateLeft(n int) *BitVector {
	return v.RotateRight(-n)
}

func (v *BitVector) ShiftRight(n int) *BitVector {
	r := v.Copy()
	l := len(v.Nodes)

	for n >= l {
		n -= l;
	}

	copy(r.Nodes[n:], r.Nodes)
	return r
}

func (v *BitVector) And(v2 *BitVector) *BitVector {
	return v.makeBinaryOperate(v2, func(n1, n2 *Node) *Node {
		return And(n1, n2)
	})
}

func (v *BitVector) Or(v2 *BitVector) *BitVector {
	return v.makeBinaryOperate(v2, func(n1, n2 *Node) *Node {
		return Or(n1, n2)
	})
}

func (v *BitVector) Xor(v2 *BitVector) *BitVector {
	return v.makeBinaryOperate(v2, func(n1, n2 *Node) *Node {
		return Xor(n1, n2)
	})
}

func (v *BitVector) add(v2 *BitVector, carry *Node) *BitVector {
	return v.makeBinaryOperate(v2, func(a, b *Node) (sum *Node) {
		sum, carry = FullAdd(a, b, carry)
		return
	})
}

func (v *BitVector) Add(v2 *BitVector) *BitVector {
	return v.add(v2, Low)
}

func (v *BitVector) Sub(v2 *BitVector) *BitVector {
	return v.add(v2.Not(), High)
}

func (v *BitVector) Eval() BitArray {
	return Eval(v.Nodes...)
}
