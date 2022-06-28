package generic

import (
	"testing"
)


func TestMuitType(t *testing.T) {
	a := 1
	b := 2
	t.Log(GetMaxNum(a, b))


	c := "123"
	d := "456"
	t.Log(GetMaxNum(c, d))

	e := int64(100)
	f := int64(0)
	t.Log(GetMaxNum(e, f))
}

func GetMaxNum[T any](a, b T) T {
	return b
}