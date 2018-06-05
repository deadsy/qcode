//-----------------------------------------------------------------------------

package main

import (
	"fmt"
	"math"

	"github.com/deadsy/qcode/asm"
	. "github.com/deadsy/qcode/qfmt"
)

//-----------------------------------------------------------------------------

const SINE2TSIZE = 4096

var sine2t [SINE2TSIZE + 1]Q31

func sin_init() {
	for i := 0; i < SINE2TSIZE+1; i++ {
		f := (float64(i) * 2.0 * math.Pi) / SINE2TSIZE
		sine2t[i] = Q31(math.MaxInt32 * math.Sin(f))
	}
}

func sin(x Q20) Q31 {

	k := uint32(x) >> 20
	k0 := x.Frac() << 11
	k1 := math.MaxInt32 - k0

	y0 := sine2t[k]
	y1 := sine2t[k+1]

	r := asm.SMMUL(int32(y0), k0)
	r = asm.SMMLA(int32(y1), k1, r)

	return Q31(r << 1)
}

//-----------------------------------------------------------------------------

func interpolate(x, y0, y1 Q27) Q27 {

	k0 := x.Frac() << 4
	k1 := math.MaxInt32 - k0

	r := asm.SMMUL(int32(y0), k0)
	r = asm.SMMLA(int32(y1), k1, r)

	return Q27(r << 1)
}

//-----------------------------------------------------------------------------

func test1() {
	sin_init()
	var i int32
	for i = -2048; i <= 2047; i++ {
		x := Q20(i << 20)
		y := sin(x)
		fmt.Printf("%f: %f\n", x.F32(), y.F32())
	}
}

func test2() {
	y0 := F32_to_Q27(0.0)
	y1 := F32_to_Q27(10.0)
	x := F32_to_Q27(0.5)
	y := interpolate(x, y0, y1)
	fmt.Printf("%f\n", y.F32())
}

//-----------------------------------------------------------------------------

func main() {
	test1()
}

//-----------------------------------------------------------------------------
