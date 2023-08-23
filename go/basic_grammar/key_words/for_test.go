package key_words

import "testing"

//TestForNil
// for nil 不会有 panic 哈
func TestForNil(t *testing.T) {
	var a []int64
	if a == nil {
		t.Log("a is nil. ")
	}
	for _, aa := range a {
		t.Log("hhhhhh", aa)
	}
	/*Output
	a is nil.
	*/
}

//TestForNilDirect
// 直接用 nil 会编译报错
func TestForNilDirect(t *testing.T) {
	//for _, a := range nil {
	//	t.Log("fafa")
	//}
	///*Output
	//annot range over ni
	//*/
}
