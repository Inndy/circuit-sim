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

// func TestRotateRight
// func TestRotateLeft
// func (v *BitVector) ShiftRight(n int) *BitVector {

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
