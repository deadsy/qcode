//-----------------------------------------------------------------------------

package qfmt

import "math"

//-----------------------------------------------------------------------------

func pow2(n int) float64 {
	return math.Pow(2.0, float64(n))
}

//-----------------------------------------------------------------------------

type Q31 int32

var MAX_F32_Q31 float32 = Q31(math.MaxInt32).F32()
var MIN_F32_Q31 float32 = Q31(math.MinInt32).F32()

// return the float32 for the q1.31 value
func (q Q31) F32() float32 {
	return float32(q) * float32(pow2(-31))
}

// return the q1.31 value for the float32
func F32_to_Q31(f float32) Q31 {
	if f < MIN_F32_Q31 || f > MAX_F32_Q31 {
		panic("out of range")
	}
	return Q31(f * float32(pow2(31)))
}

//-----------------------------------------------------------------------------

type Q27 int32

var MAX_F32_Q27 float32 = Q27(math.MaxInt32).F32()
var MIN_F32_Q27 float32 = Q27(math.MinInt32).F32()

// return the float32 for the q5.27 value
func (q Q27) F32() float32 {
	return float32(q) * float32(pow2(-27))
}

// return the fractional component for the q5.27 value
func (q Q27) Frac() int32 {
	return int32(q) & ((1 << 27) - 1)
}

// return the q5.27 value for the float32
func F32_to_Q27(f float32) Q27 {
	if f < MIN_F32_Q27 || f > MAX_F32_Q27 {
		panic("out of range")
	}
	return Q27(f * float32(pow2(27)))
}

//-----------------------------------------------------------------------------

type Q20 int32

// return the fractional component for the q12.20 value
func (q Q20) Frac() int32 {
	return int32(q) & ((1 << 20) - 1)
}

// return the float32 for the q12.20 value
func (q Q20) F32() float32 {
	return float32(q) * float32(pow2(-20))
}

//-----------------------------------------------------------------------------
