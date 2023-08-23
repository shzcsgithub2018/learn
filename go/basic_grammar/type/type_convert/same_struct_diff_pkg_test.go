// Package type_convert
// @Description: 不同包相同结构转换
package type_convert

import (
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/shzgithub2018/learn/go/basic_grammar/type/type_convert/test_pkg_1"
	"github.com/shzgithub2018/learn/go/basic_grammar/type/type_convert/test_pkg_2"
	"testing"
	"unsafe"
)

var Bob = &test_pkg_1.Student{
	Age:  15,
	Name: "Bob",
	ScoreList: []*test_pkg_1.Score{
		{
			Num:  90,
			Name: "math",
		},
		{
			Num:  80,
			Name: "chinese",
		},
	},
}

// Test_UnsafePinter
// @Description: 只适用于内存布局相同(struct不一样当然不行，但是如果只是struct里面的字段名字不同，类型排列顺序完全相同的一样可以unsafe.Pointer)的 struct
// @param t
func Test_UnsafePinter(t *testing.T) {
	type Score struct {
		Name string
		Num  int
	}
	type Student struct {
		Age  int
		Name string
		//Height    int
		ScoreList []*Score
	}

	Mary := (*Student)(unsafe.Pointer(Bob))
	t.Log(Mary)

	Lisa := (*test_pkg_2.Student)(unsafe.Pointer(Bob))
	t.Log(Lisa)
	//type StudentA struct {
	//	Age       int
	//	Name      string
	//	Height    int
	//	ScoreList []*Score
	//}
	//Kitty := (*StudentA)(unsafe.Pointer(Bob)) // 内存布局不同会 panic
	//t.Log(Kitty)
}

// Test_MarshalUnMarshal
// @Description: 序列化 反序列化  结构不同也可
// @param t
func Test_MarshalUnMarshal(t *testing.T) {
	type Score struct {
		Name string
		Num  int
	}
	type Student struct {
		Age  int
		Name string
		//Height    int
		ScoreList []*Score
	}
	b, _ := json.Marshal(Bob)
	var c Student
	_ = json.Unmarshal(b, &c)
	t.Log(c)

	// 结构不同也支持
	type StudentA struct {
		Age       int
		Name      string
		Height    int
		ScoreList []*Score
	}
	b, _ = json.Marshal(Bob)
	var d StudentA
	_ = json.Unmarshal(b, &d)
	t.Log(d)
}

// Test_TypeCovert
// @Description: golang 结构体直接转化（编译优化的东西，所以不支持跨 pkg）
// @param t
func Test_TypeCovert(t *testing.T) {
	type Score struct {
		Name string
		Num  int
	}
	type Student struct {
		Age  int
		Name string
		//Height    int
		ScoreList []*Score
	}

	type StudentA struct {
		Age  int
		Name string
		//Height    int
		ScoreList []*Score
	}

	//b := Student(Bob)  不支持 convert

	//同一个 pkg 下支持 convert
	a := Student{Name: "Bob", Age: 12}
	b := StudentA(a)
	t.Log(b)
}

//todo https://github.com/jinzhu/copier
// 待学习

func Test_Copier(t *testing.T) {
	type Score struct {
		Name string
		Num  int
	}
	type Student struct {
		Age  int
		Name string
		//Height    int
		ScoreList []*Score
	}
	var c Student
	if err := copier.Copy(&c, *Bob); err != nil {
		t.Error(err)
	}

	t.Log(c)
}
