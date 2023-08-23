package key_words

import "testing"

type A struct {
	Name string
	age  int
	a    int
}

type CC struct {
	a int
}

type B struct {
	A
	CC
	F    float64
	Name string
}

func (b *B) C() {}

func TestInternalStruct(t *testing.T) {
	//a := &B{Name: "name", Age: 27, F: 3.14}   error
	// 嵌套匿名结构体，创建结构体变量时，不能直接赋值展开的字段

	a1 := &B{A{Name: "name", age: 18}, CC{}, 3.14, ""}
	a2 := &B{A: A{Name: "name", age: 18}, F: 3.14}
	a1.C()
	a2.C()

	// 结构体可以使用嵌套结构体所有的字段和方法，不管首字母大写或小写，都可以使用
	// 匿名结构体字段访问可以简化
	// 当结构体和匿名结构体有相同的字段或方法时，编译器采用就近访问原则
	t.Log(a1.Name, a1.age, a1.F, a1.Name)

	// 嵌入两个(或多个)匿名结构体，如果两个匿名结构体有相同的字段和方法(同时结构体
	//本身没有同名的字段和方法)，在访问时，就必须明确指定匿名结构体名字，否则编译报错
	t.Log(a1.A.a, a1.CC.a)
}
