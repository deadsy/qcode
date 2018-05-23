package qfmt

import "testing"

func Test_Float32(t *testing.T) {
	tests := []struct {
		q Q5_27
		f float32
	}{
		{0, 0.0},
		{Q5_27(1 << 27), 1.0},
		{Q5_27((2 << 27) + (1 << 26)), 2.5},
		{Q5_27(0x7fffffff), MAX_F32_Q5_27},
		{Q5_27(-0x80000000), MIN_F32_Q5_27},
	}
	for _, v := range tests {
		f := v.q.Float32()
		if f != v.f {
			t.Logf("%+v %f (expected) %f (actual)\n", v.q, v.f, f)
			t.Error("FAIL")
		}
	}
}

func Test_Float64(t *testing.T) {
	tests := []struct {
		q Q5_27
		f float64
	}{
		{0, 0.0},
		{Q5_27(1 << 27), 1.0},
		{Q5_27((2 << 27) + (1 << 26)), 2.5},
		{Q5_27(0x7fffffff), MAX_F64_Q5_27},
		{Q5_27(-0x80000000), MIN_F64_Q5_27},
	}
	for _, v := range tests {
		f := v.q.Float64()
		if f != v.f {
			t.Logf("%+v %.30f (expected) %.30f (actual)\n", v.q, v.f, f)
			t.Error("FAIL")
		}
	}
}
