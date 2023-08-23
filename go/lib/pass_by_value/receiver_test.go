package pass_by_value

import "testing"

type ITest interface {
	Add5()
	Add10()
}

type Test1 struct {
	a int
}

func (t Test1) Add5() {
	t.a += 5
}
func (t Test1) Add10() {
	t.a += 10
}

type Test2 struct {
	a int
}

func (t *Test2) Add5() {
	t.a += 5
}
func (t *Test2) Add10() {
	t.a += 10
}

func TestReceiver(t *testing.T) {
	t1, t2 := Test1{a: 10}, &Test1{a: 10}
	t11, t12 := Test2{a: 10}, &Test2{a: 10}
	var it1 ITest

	//var it2 *ITest
	// 从来不会用 接口的指针！！！！！   且 golang 不会对 interface 做解引用， golang 取消了自动解引用

	// 指针接收者的接口实现，只能用指针调用。
	//it1 = t11 Type does not implement 'ITest' as the 'Add5' method has a pointer receiver
	it1 = t12 //指针调用

	// 值接收者的接口实现，既可以用值调用，也可以用指针调用。
	it1 = t1 //值调用
	it1 = t2 //指针调用

	t11.Add5()
	it1.Add5()
}
