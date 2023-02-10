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
	list.LPush(vals...)
	list1.RPush(vals...)
	list.LLen()
	list.LRange(0, -1)
	list1.LLen()
	list1.LRange(0, -1)
	/*Output
	LPush  todo   [buy some milk]
	LPush  todo   [buy some milk]
	including  1  elements
	RPush  todo-Plus   [buy some milk]
	including  1  elements
	LPush  todo   [watch tv]
	including  2  elements
	RPush  todo-Plus   [watch tv]
	including  2  elements
	LPush  todo   [finish homework]
	including  3  elements
	RPush  todo-Plus   [finish homework]
	including  3  elements
	LPush  todo   [buy some milk watch tv finish homework]
	including  6  elements
	RPush  todo-Plus   [buy some milk watch tv finish homework]
	including  6  elements
	LLen  todo
	6
	LRange  todo   0   -1
	[finish homework watch tv buy some milk finish homework watch tv buy some milk]
	LLen  todo-Plus
	6
	LRange  todo-Plus   0   -1
	[buy some milk watch tv finish homework buy some milk watch tv finish homework]
	*/
}

func TestPushPop(t *testing.T) {
	list, list1 := NewList(key), NewList(key1)

	t.Log("----------")
	list.LPushX(vals[0])
	list.LPush(vals[0])
	list.LPushX(vals[1])
	list.LPushX(vals[2])
	list.LPop()
	list.LPop()
	list.LPop()
	/*Output
	LPushX  todo   buy some milk
	including  0  elements
	LPush  todo   [buy some milk]
	including  1  elements
	LPushX  todo   watch tv
	including  2  elements
	LPushX  todo   finish homework
	including  3  elements
	LPop  todo
	finish homework
	LPop  todo
	watch tv
	LPop  todo
	buy some milk
	*/

	t.Log("----------")
	list1.RPushX(vals[0])
	list1.RPush(vals[0])
	list1.RPushX(vals[1])
	list1.RPushX(vals[2])
	list1.RPop()
	list1.RPop()
	list1.RPop()
	/*Output
	RPushX  todo-Plus   buy some milk
	including  0  elements
	RPush  todo-Plus   [buy some milk]
	including  1  elements
	RPushX  todo-Plus   watch tv
	including  2  elements
	RPushX  todo-Plus   finish homework
	including  3  elements
	RPop  todo-Plus
	finish homework
	RPop  todo-Plus
	watch tv
	RPop  todo-Plus
	buy some milk
	*/

	t.Log("----------")
	list.LPush(vals...)
	list1.RPush(vals...)
	list.LRange(0, -1)
	list1.LRange(0, -1)
	/*Output
	LPush  todo   [buy some milk watch tv finish homework]
	including  3  elements
	RPush  todo-Plus   [buy some milk watch tv finish homework]
	including  3  elements
	LRange  todo   0   -1
	[finish homework watch tv buy some milk]
	LRange  todo-Plus   0   -1
	[buy some milk watch tv finish homework]
	*/

	list.RPopLPush(list.key)
	list.LRange(0, -1)
	list.RPopLPush(list1.key)
	list.LRange(0, -1)
	list1.LRange(0, -1)
	/*Output
	RPopLPush  todo   todo
	LRange  todo   0   -1
	[buy some milk finish homework watch tv]
	RPopLPush  todo   todo-Plus
	LRange  todo   0   -1
	[buy some milk finish homework]
	LRange  todo-Plus   0   -1
	[watch tv buy some milk watch tv finish homework]
	*/
}
