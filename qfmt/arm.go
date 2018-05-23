package qfmt

// ARM SMMUL operation
func smmul(op1, op2 int32) int32 {
	x := int64(op1) * int64(op2)
	return (int32)(x >> 32)
}

// ARM SMMLA operation
func smmla(op1, op2, op3 int32) int32 {
	x := (int64(op3) << 32) + (int64(op1) * int64(op2))
	return (int32)(x >> 32)
}

// ARM SMMLS operation
func smmls(op1, op2, op3 int32) int32 {
	x := (int64(op3) << 32) - (int64(op1) * int64(op2))
	return (int32)(x >> 32)
}
