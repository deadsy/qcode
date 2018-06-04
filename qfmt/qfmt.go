//-----------------------------------------------------------------------------

package qfmt

import "math"

//-----------------------------------------------------------------------------

func pow2(n int) float64 {
	return math.Pow(2.0, float64(n))
}

//-----------------------------------------------------------------------------

type Q1_31 int32

var MAX_F32_Q1_31 float32 = Q1_31(math.MaxInt32).F32()
var MIN_F32_Q1_31 float32 = Q1_31(math.MinInt32).F32()

// return the float32 for the q1.31 value
func (q Q1_31) F32() float32 {
	return float32(q) * float32(pow2(-31))
}

// return the q1.31 value for the float32
func F32_to_Q1_31(f float32) Q1_31 {
	if f < MIN_F32_Q1_31 || f > MAX_F32_Q1_31 {
		panic("out of range")
	}
	return Q1_31(f * float32(pow2(31)))
}

//-----------------------------------------------------------------------------

type Q5_27 int32

var MAX_F32_Q5_27 float32 = Q5_27(math.MaxInt32).F32()
var MIN_F32_Q5_27 float32 = Q5_27(math.MinInt32).F32()

// return the float32 for the q5.27 value
func (q Q5_27) F32() float32 {
	return float32(q) * float32(pow2(-27))
}

// return the q5.27 value for the float32
func F32_to_Q5_27(f float32) Q5_27 {
	if f < MIN_F32_Q5_27 || f > MAX_F32_Q5_27 {
		panic("out of range")
	}
	return Q5_27(f * float32(pow2(27)))
}

//-----------------------------------------------------------------------------

type Q12_20 int32

//-----------------------------------------------------------------------------
