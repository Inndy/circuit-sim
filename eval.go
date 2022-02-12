package circuit

func eval(n... *Node) BitArray {
	ret := make([]Bit, len(n))
	for i, node := range n {
		ret[i] = node.get()
	}
	return ret
}

func Eval(n... *Node) BitArray {
	tick += 1
	return eval(n...)
}

func EvalVector(v... *BitVector) []BitArray {
	tick += 1
	a := make([]BitArray, len(v))
	for i, n := range v {
		a[i] = eval(n.Nodes...)
	}
	return a
}

func EvalBit(n *Node) Bit {
	return Eval(n)[0]
}
