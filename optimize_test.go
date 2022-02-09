package circuit

import (
	"math/rand"
	"testing"
)

func optimizeAndAddConst(a, b uint64) uint64 {
	input := NewInputVector(64)
	input.SetUint64(a)
	adder := input.Add(NewConstVector(NewBitArray(64).SetUint64(b)))
	adder.Optimize()
	return adder.Eval().GetUint64()
}

func testOptimizeAndAddConst(a, b, c uint64, t *testing.T) {
	for i := a; i <= b; i++ {
		for j := a; j <= b; j += c {
			s := i + j
			if n := optimizeAndAddConst(i, j); n != s {
				t.Errorf("optimizeAndAddConst(%d, %d) -> %d but expect %d", i, j, n, s)
			}
		}
	}
}

func TestOptimizeBitVector(t *testing.T) {
	rand.Seed(0)

	testOptimizeAndAddConst(0, 10, 1, t)
	for i := 0; i < 10; i++ {
		a := rand.Uint64()
		b := rand.Uint64() % 0x20 + 8 + a
		c := rand.Uint64() % 9 + 5
		testOptimizeAndAddConst(a, b, c, t)
	}
}
