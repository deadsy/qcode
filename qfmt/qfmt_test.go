package qfmt

import (
	"fmt"
	"testing"
)

func Test_Q2F(t *testing.T) {

	tests := []struct {
		q Q5_27
		f float32
	}{
		{0, 0.0},
	}

	for _, v := range tests {
		f := Q5_27_to_Float32(v.q)
		if f != v.f {
			fmt.Printf("%+v %f (expected) %f (actual)\n", v.q, v.f, f)
			t.Error("FAIL")
		}
	}

}
