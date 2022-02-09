package circuit

import (
	"fmt"
	"io"
)

type Gate interface {
	Get(*Node) Bit
	string(*Node, io.Writer)
}

type ConstGate struct {
	Value Bit
}

func (g *ConstGate) Get(n *Node) Bit {
	return g.Value
}

func (g *ConstGate) string(n *Node, w io.Writer) {
	var s string
	if g.Value == HIGH {
		s = "High"
	} else {
		s = "Low"
	}
	w.Write([]byte(s))
}

var variableGateIndex int

type VariableGate struct {
	BitPointer *Bit
	index int
}

func NewVariableGate(b *Bit) *VariableGate {
	variableGateIndex++
	return &VariableGate{
		BitPointer: b,
		index: variableGateIndex,
	}
}

func (g *VariableGate) Get(n *Node) Bit {
	return *(g.BitPointer)
}

func (g *VariableGate) string(n *Node, w io.Writer) {
	w.Write([]byte(fmt.Sprintf("Var_%d", g.index)))
}

type NandGate struct {}

func (g *NandGate) Get(n *Node) Bit {
	return !(n.a.get() && n.b.get())
}

func (g *NandGate) string(n *Node, w io.Writer) {
	w.Write([]byte("N("))
	n.a.string(w)
	w.Write([]byte(", "))
	n.b.string(w)
	w.Write([]byte(")"))
}

var (
	nand = &NandGate{}
)
