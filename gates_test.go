package circuit

import "testing"

type BinaryOperatorTruthTable struct {
	name string
	gate BinaryGateMaker
	truth []struct{a, b, c Bit}
}

func testBinaryTruthTable(t *testing.T, tbl *BinaryOperatorTruthTable) {
	for _, row := range tbl.truth {
		n1 := GetConstNode(row.a)
		n2 := GetConstNode(row.b)
		testNode := tbl.gate(n1, n2)
		if r := Eval(testNode)[0]; r != row.c {
			t.Errorf("%s test failed: (%s, %s -> %s but expect %s)", tbl.name, row.a, row.b, r, row.c)
		}
	}
}

func TestLogicGates (t *testing.T) {
	notGateTruthTable := []struct{ a, b Bit } {
		{ LOW,  HIGH },
		{ HIGH, LOW  },
	}

	for _, row := range notGateTruthTable {
		if r := Eval(Not(GetConstNode(row.a)))[0]; r != row.b {
			t.Errorf("Not gate test failed")
		}
	}

	testBinaryTruthTable(t, &BinaryOperatorTruthTable{
		name: "Nand",
		gate: Nand,
		truth: []struct{a, b, c Bit}{
			{ LOW,  LOW,  HIGH },
			{ LOW,  HIGH, HIGH },
			{ HIGH, LOW,  HIGH },
			{ HIGH, HIGH, LOW  },
		},
	})

	testBinaryTruthTable(t, &BinaryOperatorTruthTable{
		name: "And",
		gate: And,
		truth: []struct{a, b, c Bit}{
			{ LOW,  LOW,  LOW  },
			{ LOW,  HIGH, LOW  },
			{ HIGH, LOW,  LOW  },
			{ HIGH, HIGH, HIGH },
		},
	})

	testBinaryTruthTable(t, &BinaryOperatorTruthTable{
		name: "Or",
		gate: Or,
		truth: []struct{a, b, c Bit}{
			{ LOW,  LOW,  LOW  },
			{ LOW,  HIGH, HIGH },
			{ HIGH, LOW,  HIGH },
			{ HIGH, HIGH, HIGH },
		},
	})

	testBinaryTruthTable(t, &BinaryOperatorTruthTable{
		name: "Xor",
		gate: Xor,
		truth: []struct{a, b, c Bit}{
			{ LOW,  LOW,  LOW  },
			{ LOW,  HIGH, HIGH },
			{ HIGH, LOW,  HIGH },
			{ HIGH, HIGH, LOW  },
		},
	})

	halfAdderTruthTable := []struct{ a, b, s, c Bit }{
		{ LOW,  LOW,  LOW,  LOW  },
		{ LOW,  HIGH, HIGH, LOW  },
		{ HIGH, LOW,  HIGH, LOW  },
		{ HIGH, HIGH, LOW,  HIGH },
	}

	for _, row := range halfAdderTruthTable {
		a := GetConstNode(row.a)
		b := GetConstNode(row.b)
		sumNode, carryNode := HalfAdd(a, b)
		result := Eval(sumNode, carryNode)
		if sumBit := result[0]; sumBit != row.s {
			t.Errorf("HalfAdder sumBit error: %v (expect %v)", sumBit, row.s)
		}
		if carryBit := result[1]; carryBit != row.c {
			t.Errorf("HalfAdder carryBit error: %v (expect %v)", carryBit, row.c)
		}
	}

	fullAdderTruthTable := []struct{ a, b, c0, s, c Bit }{
		{ LOW,  LOW,  LOW,  LOW,  LOW  },
		{ LOW,  LOW,  HIGH, HIGH, LOW  },
		{ LOW,  HIGH, LOW,  HIGH, LOW  },
		{ LOW,  HIGH, HIGH, LOW,  HIGH },
		{ HIGH, LOW,  LOW,  HIGH, LOW  },
		{ HIGH, LOW,  HIGH, LOW,  HIGH },
		{ HIGH, HIGH, LOW,  LOW,  HIGH },
		{ HIGH, HIGH, HIGH, HIGH, HIGH },
	}

	for _, row := range fullAdderTruthTable {
		a := GetConstNode(row.a)
		b := GetConstNode(row.b)
		c0 := GetConstNode(row.c0)
		sumNode, carryNode := FullAdd(a, b, c0)
		result := Eval(sumNode, carryNode)
		if sumBit := result[0]; sumBit != row.s {
			t.Errorf("HalfAdder sumBit error: %v (expect %v)", sumBit, row.s)
		}
		if carryBit := result[1]; carryBit != row.c {
			t.Errorf("HalfAdder carryBit error: %v (expect %v)", carryBit, row.c)
		}
	}
}
