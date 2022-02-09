package circuit

import (
	"bytes"
	"testing"
)

func TestNewBitArray(t *testing.T) {
	w := 64 + 9
	A := NewBitArray(w)
	if len(A) != w {
		t.Error("NewBitArray returned wrong legnth of slice")
	}

	for _, b := range A {
		if b != LOW {
			t.Error("NewBitArray returned non-empty slice")
		}
	}

	for _, b := range A.GetBytes() {
		if b != 0 {
			t.Error("BitArray.GetBytes returned non-empty bytes")
		}
	}

	if A.GetUint64() != 0 {
		t.Error("BitArray.GetUint64 returned non-zero")
	}

	v0 := uint64(0x0123456789abcdef)
	A.SetUint64(v0)
	if v := A.GetUint64(); v != v0 {
		t.Errorf("BitArray.SetUint64 then BitArray.GetUint64 returned different value (%x, %x expected)", v, v0)
	}

	b0 := []byte{0xfe, 0xcd, 0xab, 0x01, 0x23, 0x45, 0x67, 0x89}
	A.SetBytes(b0)
	if b := A.GetBytes(); !bytes.HasSuffix(b, b0) {
		t.Errorf("BitArray.SetBytes then BitArray.GetBytes returned different value (%x, %x expected)", b, b0)
	}

	b1 := make([]byte, w / 8)
	v1 := uint64(1 << 63)
	for i := 0; i < 0x400; i++ {
		v1 ^= (v1 >> 1) + uint64((i * 0x1fdf3f9) << 47)
		for i := range b1 {
			b := byte(v1 >> 9)
			b ^= byte(v1 >> i)
			b ^= byte(i * 0x37)
			if i > 0 {
				b ^= b1[i - 1]
			}
			b1[i] ^= b
		}
		A.SetUint64(v1)
		if v := A.GetUint64(); v != v1 {
			t.Errorf("BitArray.SetUint64 then BitArray.GetUint64 returned different value (%x, %x expected)", v, v1)
		}

		A.SetBytes(b1)
		if b := A.GetBytes(); !bytes.HasSuffix(b, b1) {
			t.Errorf("BitArray.SetBytes then BitArray.GetBytes returned different value (%x, %x expected)", b, b0)
		}
	}
}
