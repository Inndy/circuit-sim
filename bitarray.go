package circuit

type BitArray []Bit

func NewBitArray(n int) BitArray {
	return make([]Bit, n)
}

func setBit(a []byte, o int, b Bit) {
	byteOffset := o / 8
	bitOffset := o & 7
	byte_bit := byte(1 << (7 - bitOffset))
	if b {
		a[byteOffset] |= byte_bit
	} else {
		a[byteOffset] &= ^byte_bit
	}
}

func (a BitArray) GetBytes() []byte {
	n := (len(a) + 7) / 8
	ret := make([]byte, n)

	bOff := len(a) - 1
	for i := len(ret) - 1; i >= 0; i-- {
		b := byte(0)
		for j := 0; j < 8; j++ {
			b >>= 1
			if bOff >= 0 && a[bOff] {
				b |= 0x80
			}
			bOff--
		}
		ret[i] = b
	}

	return ret
}

func (a BitArray) SetBytes(arr []byte) BitArray {
	bIndex := len(a) - 1
	for byteIdx := len(arr) - 1; byteIdx >= 0; byteIdx-- {
		for i := 0; i < 8; i++ {
			bit := LOW
			if ((arr[byteIdx] >> i) & 1) == 1 {
				bit = HIGH
			}

			a[bIndex] = bit
			bIndex--
			if bIndex < 0 {
				return a
			}
		}
	}

	return a
}

func (a BitArray) GetUint64() (val uint64) {
	for i, b := range a {
		if i > 0 {
			val = val << 1
		}
		if b {
			val |= 1
		}
	}

	return
}

func (a BitArray) SetUint64(val uint64) BitArray {
	n := len(a) - 1
	for i := range a {
		b := LOW
		if ((val >> (n - i)) & 1) != 0 {
			b = HIGH
		}
		a[i] = b
	}

	return a
}
