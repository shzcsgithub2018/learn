package list

import "testing"

var (
	key  = "todo"
	key1 = "todo-Plus"
	vals = []string{
		"buy some milk",
		"watch tv",
		"finish homework",
	}
)

func TestPush(t *testing.T) {
	list, list1 := NewList(key), NewList(key1)
	for _, val := range vals {
		list.LPush(val)
		list1.RPush(val)
	}
	list.LRange(0, -1)
	list1.LRange(0, -1)
}
