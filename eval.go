package circuit

func Eval(n... *Node) BitArray {
	tick += 1
	ret := make([]Bit, len(n))
	for i, node := range n {
		ret[i] = node.get()
	}
	return ret
}

func EvalBit(n *Node) Bit {
	return Eval(n)[0]
}
