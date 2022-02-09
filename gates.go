package circuit

type UnaryGateMaker func (n1 *Node) *Node
type BinaryGateMaker func (n1, n2 *Node) *Node

func Nand(a, b *Node) *Node {
	return NewNode(nand, a, b)
}

func Not(n *Node) *Node {
	return Nand(n, n)
}

func And(a, b *Node) *Node {
	return Not(Nand(a, b))
}

func Or(a, b *Node) *Node {
	return Nand(Not(a), Not(b))
}

func Xor(a, b *Node) *Node {
	mid := Nand(a, b)
	l := Nand(a, mid)
	r := Nand(b, mid)
	return Nand(l, r)
}

func HalfAdd(a, b *Node) (sum, carry *Node) {
	return Xor(a, b), And(a, b)
}

func FullAdd(a, b, c *Node) (sum, carry *Node) {
	s1 := Xor(a, b)
	sum = Xor(s1, c)
	carry = Or(And(a, b), And(c, s1))
	return
}
