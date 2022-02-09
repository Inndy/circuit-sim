package circuit

import (
	"math/rand"
	"testing"
)

const TestLoopCount = 1000

func TestNot(t *testing.T) {
	rand.Seed(0)

	input1 := NewInputVector(64)
	compute := input1.Not()
	for i := 0; i < TestLoopCount; i++ {
		a := rand.Uint64()
		s := ^a
		input1.SetUint64(a)
		if v := compute.Eval().GetUint64(); s != v {
			t.Errorf("TestNot ^%x = %x (expect %x)", a, v, s)
		}
	}
}

func rotateRightU64(v uint64, s, w int) uint64 {
	mask := uint64(0xffffffffffffffff) >> (64 - w)
	return ((v >> s) | (v << (w - s))) & mask
}

func TestRotateRight(t *testing.T) {
	rand.Seed(0)

	input := NewInputVector(64)
	for i := 0; i < TestLoopCount; i++ {
		a := rand.Uint64()
		s := i % 64
		v := rotateRightU64(a, s, 64)
		compute := input.RotateRight(s)
		input.SetUint64(a)

		if r := compute.Eval().GetUint64(); r != v {
			t.Errorf("TestRotateRight %x >> %d = %x (expect %x)", a, s, r, v)
		}
	}
}

func TestRotateLeft(t *testing.T) {
	rand.Seed(0)

	input := NewInputVector(60)
	for i := 0; i < TestLoopCount; i++ {
		a := rand.Uint64() >> 4
		s := i % 60
		v := rotateRightU64(a, 60 - s, 60)
		compute := input.RotateLeft(s)
		input.SetUint64(a)

		if r := compute.Eval().GetUint64(); r != v {
			t.Errorf("TestRotateLeft %x >> %d = %x (expect %x)", a, s, r, v)
		}
	}
}

func TestShiftRight(t *testing.T) {
	rand.Seed(0)

	input := NewInputVector(64)
	for i := 0; i < TestLoopCount; i++ {
		a := rand.Uint64()
		s := i % 64
		v := a >> s
		compute := input.ShiftRight(s)
		input.SetUint64(a)

		if r := compute.Eval().GetUint64(); r != v {
			t.Errorf("TestShiftRight %x >> %d = %x (expect %x)", a, s, r, v)
		}
	}
}

func TestAnd(t *testing.T) {
	rand.Seed(0)

	input1 := NewInputVector(64)
	input2 := NewInputVector(64)
	compute := input1.And(&input2.BitVector)
	for i := 0; i< TestLoopCount; i++ {
		a := rand.Uint64()
		b := rand.Uint64()
		s := a & b
		input1.SetUint64(a)
		input2.SetUint64(b)
		if v := compute.Eval().GetUint64(); s != v {
			t.Errorf("TestAnd %x ^ %x = %x (expect %x)", a, b, v, s)
		}
	}
}

func TestOr(t *testing.T) {
	rand.Seed(0)

	input1 := NewInputVector(64)
	input2 := NewInputVector(64)
	compute := input1.Or(&input2.BitVector)
	for i := 0; i< TestLoopCount; i++ {
		a := rand.Uint64()
		b := rand.Uint64()
		s := a | b
		input1.SetUint64(a)
		input2.SetUint64(b)
		if v := compute.Eval().GetUint64(); s != v {
			t.Errorf("TestOr %x ^ %x = %x (expect %x)", a, b, v, s)
		}
	}
}

func TestXor(t *testing.T) {
	rand.Seed(0)

	input1 := NewInputVector(64)
	input2 := NewInputVector(64)
	compute := input1.Xor(&input2.BitVector)
	for i := 0; i< TestLoopCount; i++ {
		a := rand.Uint64()
		b := rand.Uint64()
		s := a ^ b
		input1.SetUint64(a)
		input2.SetUint64(b)
		if v := compute.Eval().GetUint64(); s != v {
			t.Errorf("TestXor %x ^ %x = %x (expect %x)", a, b, v, s)
		}
	}
}

func TestAdd(t *testing.T) {
	rand.Seed(0)

	input1 := NewInputVector(64)
	input2 := NewInputVector(64)
	compute := input1.Add(&input2.BitVector)
	for i := 0; i< TestLoopCount; i++ {
		a := rand.Uint64()
		b := rand.Uint64()
		s := a + b
		input1.SetUint64(a)
		input2.SetUint64(b)
		if v := compute.Eval().GetUint64(); s != v {
			t.Errorf("TestAdd %x + %x = %x (expect %x)", a, b, v, s)
		}
	}
}

func TestSub(t *testing.T) {
	rand.Seed(0)

	input1 := NewInputVector(64)
	input2 := NewInputVector(64)
	compute := input1.Sub(&input2.BitVector)
	for i := 0; i< TestLoopCount; i++ {
		a := rand.Uint64()
		b := rand.Uint64()
		s := a - b
		input1.SetUint64(a)
		input2.SetUint64(b)
		if v := compute.Eval().GetUint64(); s != v {
			t.Errorf("TestSub %x - %x = %x (expect %x)", a, b, v, s)
		}
	}
}
