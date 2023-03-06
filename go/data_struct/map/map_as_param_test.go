package map_test

import "testing"

// TestMapAsParam
// @Description: 无论是 make 还是 字面值初始化，实际运行中都是一个 hmap 指针，函数里面对map进行修改时，会同时修改源map的值，但是将map修改为nil时，则达不到预期效果
// @param t
func TestMapAsParam(t *testing.T) {
	var m = make(map[int64]string, 0)
	Add(m, "a", "b", "c")
	t.Log(m)

	mm := map[int64]string{}
	Add(mm, "a", "b", "c")
	t.Log(mm)
}

func Add(m map[int64]string, strList ...string) {
	for idx, str := range strList {
		m[int64(idx)] = str
	}
}
