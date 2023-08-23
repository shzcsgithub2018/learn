package key_words

import "testing"

func Test_Int(t *testing.T) {
	numPtr := new(int64)
	if numPtr == nil {
		t.Error("numPtr is nil")
	} else {
		t.Log(*numPtr) // 0
	}

	numPtrPtr := new(*int64) // new 了 一个指向指针的指针，*int64 初始化为 nil
	if numPtrPtr == nil {
		t.Error("numPtrPtr is nil")
	} else {
		t.Log(*numPtrPtr) // <nil>
	}
}

func Test_String(t *testing.T) {
	strPtr := new(string)
	if strPtr == nil {
		t.Error("strPtr is nil")
	} else {
		t.Log(*strPtr) // ""
	}

	strPtrPtr := new(*string) // new 了 一个指向指针的指针，*int64 初始化为 nil
	if strPtrPtr == nil {
		t.Error("strPtrPtr is nil")
	} else {
		t.Log(*strPtrPtr) // <nil>
	}
}

func Test_Float(t *testing.T) {
	floatPtr := new(float64)
	if floatPtr == nil {
		t.Error("floatPtr is nil")
	} else {
		t.Log(*floatPtr) // 0
	}

	floatPtrPtr := new(*float64) // new 了 一个指向指针的指针，*int64 初始化为 nil
	if floatPtrPtr == nil {
		t.Error("floatPtrPtr is nil")
	} else {
		t.Log(*floatPtrPtr) // <nil>
	}
}

func Test_all(t *testing.T) {
	type Tmp struct {
		a    int
		aPtr *int

		b    string
		bPtr *string

		c    float64
		cPtr *float64

		f    func()
		fPtr *func()

		s    []int
		sPtr *[]int

		m    map[int]string
		mPtr *map[int]string

		ch    chan int
		chPtr *chan int
	}

	tmp := new(Tmp)
	t.Log(tmp)

	// int
	t.Log(new(int), new(*int))

	// string

	// float

	// func

	// slice

	// map

	// chan
}
