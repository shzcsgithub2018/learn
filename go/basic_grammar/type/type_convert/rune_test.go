package type_convert

import "testing"

func TestRune(t *testing.T) {
	str := "我和我的祖国，一刻也不能分割"
	t.Log(len(str))
	t.Log(len([]byte(str)))
	t.Log(len([]rune(str)))
}
