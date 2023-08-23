package slice

import "testing"

func TestSliceAsParam(t *testing.T) {
	var s1 []int
	Add(s1, 1, 2, 3)
	t.Log(s1)

	var s2 = make([]int, 0)
	Add(s2, 1, 2, 3)
	t.Log(s2)

	var s3 = []int{}
	Add(s3, 1, 2, 3)
	t.Log(s3)
}

func Add(s []int, numList ...int) {
	for idx, num := range numList {
		s[idx] = num
	}
}
