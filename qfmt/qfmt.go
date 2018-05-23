package qfmt

import "math"

func pow2(n int) float64 {
	return math.Pow(2.0, float64(n))
}

type Q5_27 int32

const MAX_Q5_27 = (1 << 31) - 1
const MIN_Q5_27 = -(1 << 31)

var MAX_F32_Q5_27 float32 = Q5_27(MAX_Q5_27).Float32()
var MIN_F32_Q5_27 float32 = Q5_27(MIN_Q5_27).Float32()

var MAX_F64_Q5_27 float64 = Q5_27(MAX_Q5_27).Float64()
var MIN_F64_Q5_27 float64 = Q5_27(MIN_Q5_27).Float64()

// return the float32 for the q5.27 value
func (q Q5_27) Float32() float32 {
	if q < MIN_Q5_27 || q > MAX_Q5_27 {
		panic("out of range")
	}
	return float32(q) * float32(pow2(-27))
}

// return the float64 for the q5.27 value
func (q Q5_27) Float64() float64 {
	if q < MIN_Q5_27 || q > MAX_Q5_27 {
		panic("out of range")
	}
	return float64(q) * pow2(-27)
}

// return the q5.27 value for the float32
func Float32_to_Q5_27(f float32) Q5_27 {
	if f < MIN_F32_Q5_27 || f > MAX_F32_Q5_27 {
		panic("out of range")
	}
	return Q5_27(f * float32(pow2(27)))
}
