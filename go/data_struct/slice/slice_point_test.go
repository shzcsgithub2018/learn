package slice

import "testing"

// TestSlicePoint
// @Description: https://blog.csdn.net/m0_46251547/article/details/123233179
// nil切片 VS 空切片
// @param t
func TestSlicePoint(t *testing.T) {
	type tmp struct{}
	a := make([]*tmp, 0) // len:0  cap: 0

	var b []*tmp
	if b == nil {
		t.Log("nil 切片")
	} // nil

	if a == nil && b == nil {
		t.Log("SSSS")
	}

	for _, aa := range a {
		t.Log("fasdfas", aa)
	}
	for _, bb := range b {
		t.Log("fasfdsfas", bb)
	}

	t.Log(len(a), len(b))
}
