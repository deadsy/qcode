//-----------------------------------------------------------------------------

package main

import (
	"fmt"
	"math"

	"github.com/deadsy/qcode/asm"
	"github.com/deadsy/qcode/qfmt"
)

//-----------------------------------------------------------------------------

const SINE2TSIZE = 4096

var sine2t [SINE2TSIZE + 1]int32

func sin_init() {
	for i := 0; i < SINE2TSIZE+1; i++ {
		f := (float64(i) * 2.0 * math.Pi) / SINE2TSIZE
		sine2t[i] = int32(math.MaxInt32 * math.Sin(f))
	}
}

func sin_q31(phase int32) qfmt.Q1_31 {

	p := uint32(phase)
	pi := p >> 20

	y1 := sine2t[pi]
	y2 := sine2t[1+pi]
	pf := int32((p & 0xfffff) << 11)
	pfc := math.MaxInt32 - pf

	rr := asm.SMMUL(y1, pfc)
	rr = asm.SMMLA(y2, pf, rr)

	return qfmt.Q1_31(rr << 1)
}

//-----------------------------------------------------------------------------

func main() {
	sin_init()
  var i int32
  for i = -2048; i < 2047; i ++ {
    fmt.Printf("%d: %f\n", i, sin_q31(i << 20).F32())
  }
}

//-----------------------------------------------------------------------------
