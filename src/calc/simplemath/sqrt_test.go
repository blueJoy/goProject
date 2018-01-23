package simplemath

import "testing"

func TestSqrt(t *testing.T) {
	v := Sqrt(16)
	if v != 4 {
		t.Error("Sqrt(16) failed. Got %v, expected 4.", v)
	}
}
