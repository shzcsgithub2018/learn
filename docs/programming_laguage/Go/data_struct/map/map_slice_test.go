package map_test

import (
	"fmt"
	"testing"
)

func TestMapSlice(t *testing.T) {
	mm := map[int64][]*string{}
	fmt.Println(len(mm[1]), cap(mm[1]))
	funcTest(mm)
	fmt.Println(len(mm[1]), cap(mm[1]))
	t.Log(mm)
}

func funcTest(mm map[int64][]*string) {
	str := "41414"
	fmt.Println(len(mm[1]), cap(mm[1]))
	mm[1] = append(mm[1], &str)
	fmt.Println(len(mm[1]), cap(mm[1]))
	mm[1] = append(mm[1], &str)
	fmt.Println(len(mm[1]), cap(mm[1]))
	mm[1] = append(mm[1], &str)
	fmt.Println(len(mm[1]), cap(mm[1]))

	funcTest2(mm[2])
}

func funcTest2(ss []*string) {
	str := "aaaa"
	ss = append(ss, &str)
}
