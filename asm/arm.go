package asm

// ARM SMMUL operation
func SMMUL(op1, op2 int32) int32 {
	x := int64(op1) * int64(op2)
	return (int32)(x >> 32)
}

// ARM SMMULR operation
func SMMULR(op1, op2 int32) int32 {
	x := 0x80000000 + int64(op1)*int64(op2)
	return (int32)(x >> 32)
}

// ARM SMMLA operation
func SMMLA(op1, op2, op3 int32) int32 {
	x := (int64(op3) << 32) + (int64(op1) * int64(op2))
	return (int32)(x >> 32)
}

// ARM SMMLAR operation
func SMMLAR(op1, op2, op3 int32) int32 {
	x := 0x80000000 + (int64(op3) << 32) + (int64(op1) * int64(op2))
	return (int32)(x >> 32)
}

// ARM SMMLS operation
func SMMLS(op1, op2, op3 int32) int32 {
	x := (int64(op3) << 32) - (int64(op1) * int64(op2))
	return (int32)(x >> 32)
}

// ARM SMMLSR operation
func SMMLSR(op1, op2, op3 int32) int32 {
	x := 0x80000000 + (int64(op3) << 32) - (int64(op1) * int64(op2))
	return (int32)(x >> 32)
}

// ARM ASR operation
func ASR(op1 int32, n int) int32 {
	if n < 0 || n > 31 {
		panic("n out of range")
	}
	return op1 >> uint(n)
}

// ARM LSL operation
func LSL(op1 int32, n int) int32 {
	if n < 0 || n > 31 {
		panic("n out of range")
	}
	return op1 << uint(n)
}

// ARM SSAT operation
func SSAT(op1 int32, n, s int) int32 {
	if n < 1 || n > 32 {
		panic("n out of range")
	}
	// optional shifts
	x := op1
	if s < 0 {
		x = ASR(x, -s)
	}
	if s > 0 {
		x = LSL(x, s)
	}
	// saturate
	lower := int32(-(1 << uint(n-1)))
	if x < lower {
		return lower
	}
	upper := int32((1 << uint(n-1)) - 1)
	if x > upper {
		return upper
	}
	return x
}

// ARM USAT operation
func USAT(op1 int32, n, s int) uint32 {
	if n < 0 || n > 31 {
		panic("n out of range")
	}
	// optional shifts
	x := op1
	if s < 0 {
		x = ASR(x, -s)
	}
	if s > 0 {
		x = LSL(x, s)
	}
	// saturate
	if x < 0 {
		return 0
	}
	upper := int32((1 << uint(n)) - 1)
	if x > upper {
		return uint32(upper)
	}
	return uint32(x)
}
