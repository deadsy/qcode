package main

import (
	"fmt"
	"math"
)

/*
Conversion between float and Q5.27 formats
*/

const QM = 5  // bits before the binary point
const QN = 27 // bits after the binary point

const MAX_Q = (1 << (QM + QN - 1)) - 1
const MIN_Q = -(1 << (QM + QN - 1))

var MAX_F float32 = q2f(MAX_Q)
var MIN_F float32 = q2f(MIN_Q)

func pow2(n int) float32 {
	return float32(math.Pow(2.0, float64(n)))
}

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

// return the float for the q value
func q2f(q int32) float32 {
	if q < MIN_Q || q > MAX_Q {
		panic("out of range")
	}
	return float32(q) * pow2(-QN)
}

// return the q value for the float
func f2q(f float32) int32 {
	if f < MIN_F || f > MAX_F {
		panic("out of range")
	}
	return int32(f * pow2(QN))
}

func main() {

	fmt.Printf("%.20f %.20f\n", MIN_F, MAX_F)
	fmt.Printf("%d %d\n", MIN_Q, MAX_Q)

	tests := []float32{0.0, 1.0, -1.0, -10.0, 10.0, -16.0, 15.99}
	for _, f := range tests {
		q := f2q(f)
		fmt.Printf("%f %d %f\n", f, q, q2f(q))
	}

	a := f2q(-3.0)
	b := f2q(-4.0)

	fmt.Printf("%f\n", q2f(a+b))

}
