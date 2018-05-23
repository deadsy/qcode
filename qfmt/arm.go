package qfmt

func smmul(op1, op2 int32) int32 {
	x := int64(op1) * int64(op2)
	return (int32)(x >> 32)
}

func smmla(op1, op2, op3 int32) int32 {
	x := (int64(op1) * int64(op2)) + (int64(op3) << 32)
	return (int32)(x >> 32)
}

func smmls(op1, op2, op3 int32) int32 {
	x := (int64(op1) * int64(op2)) - (int64(op3) << 32)
	return (int32)(x >> 32)
}
