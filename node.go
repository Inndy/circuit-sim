package circuit

import (
	"bytes"
	"fmt"
	"io"
)

var tick int

type Node struct {
	Name string
	g Gate
	a *Node
	b *Node
	tick int
	cache Bit
	i int
}

var nodeIndex = 1 // start from 2

func NewNode(g Gate, a, b *Node) *Node {
	nodeIndex++
	return &Node{
		g: g,
		a: a,
		b: b,
		i: nodeIndex,
	}
}

func (n *Node) get() Bit {
	if n.tick == tick {
		return n.cache
	}
	n.cache = n.g.Get(n)
	n.tick = tick
	return n.cache
}

func (n *Node) string(w io.Writer) {
	if n.tick == tick && n.i >= 2 {
		fmt.Fprintf(w, "&{_%d}", n.i)
		return
	}

	n.tick = tick
	if n.Name != "" {
		w.Write([]byte(n.Name))
	} else {
		fmt.Fprintf(w, "_%d:", n.i)
		n.g.string(n, w)
	}
}

func (n *Node) String() string {
	tick++
	buf := &bytes.Buffer{}
	n.g.string(n, buf)
	return buf.String()
}

var DumpIndex int
func (n *Node) dump(name string, w io.Writer) {
	if n.tick == tick {
		return
	}

	if n.Name == "" {
		n.Name = name
	}
	n.tick = tick

	if n.a == nil || n.b == nil {
		fmt.Fprintf(w, "/*%s*/\n", n.Name)
		return
	}

	na := "nil"
	nb := "nil"

	if n.a != nil {
		if n.a.Name != "" {
			na = n.a.Name
		} else {
			na = fmt.Sprintf("n[%d]", DumpIndex)
			DumpIndex++
			n.a.dump(na, w)
		}
	}

	if n.b != nil {
		if n.b.Name != "" {
			nb = n.b.Name
		} else {
			nb = fmt.Sprintf("n[%d]", DumpIndex)
			DumpIndex++
			n.b.dump(nb, w)
		}
	}

	fmt.Fprintf(w, "%s = circuit.Nand(%s, %s)\n", name, na, nb)
}

func (n *Node) Dump(name string, w io.Writer) {
	tick++
	n.dump(name, w)
}

func GetConstNode(b Bit) *Node {
	if b {
		return High
	} else {
		return Low
	}
}

var (
	Low = &Node{g: &ConstGate{LOW}, Name: "Low", i: 0}
	High = &Node{g: &ConstGate{HIGH}, Name: "High", i: 1}
)
