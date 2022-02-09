package circuit

var EliminatedNodeCount = 0

func isKnownConst(n *Node) (Bit, bool) {
	switch n {
	case High:
		return HIGH, true
	case Low:
		return LOW, true
	default:
		return LOW, false
	}
}

func isKnownHigh(n *Node) bool {
	b, ok := isKnownConst(n)
	return ok && b == HIGH
}

func isKnownLow(n *Node) bool {
	b, ok := isKnownConst(n)
	return ok && b == LOW
}

func isKnownNotGate(n *Node) (*Node, bool) {
	if n.g == nand {
		if n.a == n.b {
			return n.a, true
		}
		if isKnownHigh(n.a) {
			return n.b, true
		}
		if isKnownHigh(n.b) {
			return n.a, true
		}
	}

	return nil, false
}

func eliminateRepeatedNotGate(n **Node) {
	for {
		if g, ok := isKnownNotGate(*n); ok {
			if g2, ok := isKnownNotGate(g); ok {
				EliminatedNodeCount += 2
				*n = g2
				continue
			}
		}

		break
	}
}

func Optimize(n *Node) *Node {
	eliminateRepeatedNotGate(&n)

	if g, ok := isKnownNotGate(n); ok {
		if b, ok := isKnownConst(g); ok {
			n = GetConstNode(!b)
		}
	}

	if n.g == nand {
		l_a := isKnownLow(n.a)
		l_b := isKnownLow(n.b)
		if l_a || l_b {
			EliminatedNodeCount++
			return High
		}
	}

	if n.a != nil {
		n.a = Optimize(n.a)
	}
	if n.b != nil {
		n.b = Optimize(n.b)
	}

	return n
}
