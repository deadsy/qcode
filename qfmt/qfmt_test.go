package qfmt

import "testing"

func Test_Float32(t *testing.T) {
	tests := []struct {
		q Q27
		f float32
	}{
		{0, 0.0},
		{Q27(1 << 27), 1.0},
		{Q27((2 << 27) + (1 << 26)), 2.5},
		{Q27(0x7fffffff), MAX_F32_Q27},
		{Q27(-0x80000000), MIN_F32_Q27},
	}
	for _, v := range tests {
		f := v.q.F32()
		if f != v.f {
			t.Logf("%+v %f (expected) %f (actual)\n", v.q, v.f, f)
			t.Error("FAIL")
		}
	}
}
