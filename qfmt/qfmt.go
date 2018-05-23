package qfmt

import "math"

func pow2(n int) float32 {
	return float32(math.Pow(2.0, float64(n)))
}

type Q5_27 int32

const MAX_Q5_27 = (1 << 31) - 1
const MIN_Q5_27 = -(1 << 31)

var MAX_F32_Q5_27 float32 = Q5_27_to_Float32(MAX_Q5_27)
var MIN_F32_Q5_27 float32 = Q5_27_to_Float32(MIN_Q5_27)

// return the float32 for the q5.27 value
func Q5_27_to_Float32(q Q5_27) float32 {
	if q < MIN_Q5_27 || q > MAX_Q5_27 {
		panic("out of range")
	}
	return float32(q) * pow2(-27)
}

// return the q5.27 value for the float32
func Float32_to_Q5_27(f float32) Q5_27 {
	if f < MIN_F32_Q5_27 || f > MAX_F32_Q5_27 {
		panic("out of range")
	}
	return Q5_27(f * pow2(27))
}
