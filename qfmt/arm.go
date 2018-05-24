package qfmt

// ARM SMMUL operation
func smmul(op1, op2 int32) int32 {
	x := int64(op1) * int64(op2)
	return (int32)(x >> 32)
}

// ARM SMMULR operation
func smmulr(op1, op2 int32) int32 {
	x := 0x80000000 + int64(op1)*int64(op2)
	return (int32)(x >> 32)
}

// ARM SMMLA operation
func smmla(op1, op2, op3 int32) int32 {
	x := (int64(op3) << 32) + (int64(op1) * int64(op2))
	return (int32)(x >> 32)
}

// ARM SMMLAR operation
func smmlar(op1, op2, op3 int32) int32 {
	x := 0x80000000 + (int64(op3) << 32) + (int64(op1) * int64(op2))
	return (int32)(x >> 32)
}

// ARM SMMLS operation
func smmls(op1, op2, op3 int32) int32 {
	x := (int64(op3) << 32) - (int64(op1) * int64(op2))
	return (int32)(x >> 32)
}

// ARM SMMLSR operation
func smmlsr(op1, op2, op3 int32) int32 {
	x := 0x80000000 + (int64(op3) << 32) - (int64(op1) * int64(op2))
	return (int32)(x >> 32)
}

//ASR #s where s is in the range 1 to 31
//LSL #s where s is in the range 0 to 31.

// ARM SSAT operation
func ssat(rd int32, n, s int) int32 {
	if n < 1 || n > 32 {
		panic("n out of range")
	}
	if s < -31 || s > 31 {
		panic("s out of range")
	}
	// TODO
	return rd
}

// ARM USAT operation
func usat(rd int32, n, s int) uint32 {
	if n < 0 || n > 31 {
		panic("n out of range")
	}
	if s < -31 || s > 31 {
		panic("s out of range")
	}
	// TODO
	return uint32(rd)
}
