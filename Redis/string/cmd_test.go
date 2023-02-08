package String

import (
	"github.com/redis/go-redis/v9"
	"github.com/shzgithub2018/learn/Redis/dal"
	"testing"
	"time"
)

type TestFunc func(t *testing.T)

var FuncList = []TestFunc{
	TestSet,
	TestSetNX,
	TestSetXX,
	TestExpiration,
	TestKeepTTL,
	TestGetSet,
	TestMSetAndMGet,
	TestMSetNX,
	TestStrLen,
}

var m = map[string]string{
	"message":  "hello world",
	"homepage": "redis.io",
	"price":    "2.56",
	"number":   "10086",
	"book":     "The Design and Implementation of Redis",
}

const (
	ZERO  = 0
	One   = time.Duration(1) * time.Second
	Ten   = time.Duration(10) * time.Second
	Three = time.Duration(3) * time.Second
	Four  = time.Duration(4) * time.Second
)

// TestSet
// @Description:
// @param t
func TestSet(t *testing.T) {
	key, val, expiration := "message", m["message"], Ten
	handler := NewRedisString(ctx, t, WithKeyVal(key, val), WithExpiration(expiration))
	handler.Set()
	/*Output
	set message "hello world" EX 10
	OK
	*/
}

// TestSetNX
// @Description:
// @param t
func TestSetNX(t *testing.T) {
	key, val, expiration := "message", m["message"], Three
	handler := NewRedisString(ctx, t, WithKeyVal(key, val), WithExpiration(expiration))

	handler.SetNX()
	handler.SetNX()
	/*Output
	  set message "hello world" NX EX 3
	  success
	  set message "hello world" NX EX 3
	  fail
	*/
}

// TestSetXX
// @Description:
// @param t
func TestSetXX(t *testing.T) {
	key, val, expiration := "message", m["message"], Three
	handler := NewRedisString(ctx, t, WithKeyVal(key, val), WithExpiration(expiration))

	handler.SetXX()
	handler.Set()
	handler.SetXX()
	/*OutPut
	  	set message "hello world" XX EX 3
		fail
		set message "hello world" EX 3
		OK
		set message "hello world" XX EX 3
		success
	*/
}

// TestGet
// @Description:
// @param t
func TestGet(t *testing.T) {
	key, val, expiration := "message", m["message"], Three
	handler := NewRedisString(ctx, t, WithKeyVal(key, val), WithExpiration(expiration))

	handler.Get()
	handler.Set()
	handler.Get()
	/*Output
	redis: nil
	set message "hello world" EX 3
	OK
	"hello world"
	*/
}

// TestExpiration
// @Description:
// @param t
func TestExpiration(t *testing.T) {
	key, val, expiration := "message", m["message"], Three
	handler := NewRedisString(ctx, t, WithKeyVal(key, val), WithExpiration(expiration))

	handler.Get()
	handler.Set()
	handler.Get()
	time.Sleep(Four)
	t.Log("After 4 second……")
	handler.Get()
	/*OutPut
	get  message
	redis: nil

	set message "hello world" EX 3
	OK
	get  message
	"hello world"
	After 4 second……
	get  message
	redis: nil
	*/
}

// TestKeepTTL
// @Description:
// @param t
func TestKeepTTL(t *testing.T) {
	key, val, expiration := "message", m["message"], Three
	handler := NewRedisString(ctx, t, WithKeyVal(key, val), WithExpiration(expiration))

	handler.Set()

	val, expiration = "Hello World Plus", redis.KeepTTL
	handler = NewRedisString(ctx, t, WithKeyVal(key, val), WithKeepTTL())

	handler.Set()
	handler.Get()
	time.Sleep(Four)
	t.Log("After 4 second……")
	handler.Get()
	/*Output
	set message "hello world" EX 3
	OK
	set message "Hello World Plus" EX  KeepTTL
	OK
	get  message
	"Hello World Plus"
	After 4 second……
	get  message
	redis: nil
	*/
}

// TestGetSet
// @Description:
// @param t
func TestGetSet(t *testing.T) {
	key, val := "message", m["message"]
	handler := NewRedisString(ctx, t, WithKeyVal(key, val))
	handler.GetSet()

	val += " Plus"
	handler = NewRedisString(ctx, t, WithKeyVal(key, val))
	handler.GetSet()

	val += " Plus"
	handler = NewRedisString(ctx, t, WithKeyVal(key, val))
	handler.GetSet()
	/* Output
	GetSet message "hello world"
	redis: nil
	GetSet message "hello world Plus"
	hello world
	GetSet message "hello world Plus Plus"
	hello world Plus
	*/
}

// TestMSetAndMGet
// @Description: 一次网络通信，复杂度 O(N)
// @param t
func TestMSetAndMGet(t *testing.T) {
	//   - MSet("key1", "value1", "key2", "value2")
	//   - MSet([]string{"key1", "value1", "key2", "value2"})
	//   - MSet(map[string]interface{}{"key1": "value1", "key2": "value2"})

	var (
		key1, val1 = "message", m["message"]
		key2, val2 = "homepage", m["homepage"]

		list = []string{key1, val1 + " list", key2, val2 + " list"}

		kvMap = map[string]interface{}{key1: val1 + "map", key2: val2 + "map"}
	)

	MGet := func() {
		cmd := dal.GetRedisCli().MGet(ctx, key1, key2)
		if err := cmd.Err(); err != nil {
			t.Error(err)
		}
		t.Log(key1, cmd.Val()[0])
		t.Log(key2, cmd.Val()[1])
	}

	cmd := dal.GetRedisCli(ctx).MSet(ctx, key1, val1, key2, val2)
	if err := cmd.Err(); err != nil {
		t.Error(err)
	}
	t.Log(cmd.Val())
	MGet()

	cmd = dal.GetRedisCli(ctx).MSet(ctx, list)
	if err := cmd.Err(); err != nil {
		t.Error(err)
	}
	t.Log(cmd.Val())
	MGet()

	cmd = dal.GetRedisCli(ctx).MSet(ctx, kvMap)
	if err := cmd.Err(); err != nil {
		t.Error(err)
	}
	t.Log(cmd.Val())
	MGet()

	key1 = "what"
	MGet()

	/*Output
	OK
	message hello world
	homepage redis.io
	OK
	message hello world list
	homepage redis.io list
	OK
	message hello worldmap
	homepage redis.iomap
	what <nil>  // 不存在的键返回 nil 值
	homepage redis.iomap
	*/
}

// TestMSetNX
// @Description:
// @param t
func TestMSetNX(t *testing.T) {
	var (
		key1, val1 = "message2", m["message"]
		key2, val2 = "homepage2", m["homepage"]

		kvMap = map[string]interface{}{key1: val1, key2: val2}
	)

	handler := NewRedisString(ctx, t, WithMKeyVal(kvMap))

	handler.MSetNX()
	handler.MSetNX()

	kvMap["a"] = "b"
	handler = NewRedisString(ctx, t, WithMKeyVal(kvMap))
	handler.MSetNX()
	/*Output
	MSetNX  message2 "hello world" homepage2 "redis.io"
	fail
	MSetNX  message2 "hello world" homepage2 "redis.io"
	fail
	MSetNX  a "b" message2 "hello world" homepage2 "redis.io"
	fail
	*/
}

// TestStrLen
// @Description:
// @param t
func TestStrLen(t *testing.T) {
	key, val, expiration := "number", "5reafsf", Three
	handler := NewRedisString(ctx, t, WithKeyVal(key, val), WithExpiration(expiration))
	handler.Set()
	handler.StrLen()
	/*Output
	OK
	7
	*/
}

// TestGetRange
// @Description:
// @param t
func TestGetRange(t *testing.T) {
	key, val, expiration := "number", "5reafsf", Three
	handler := NewRedisString(ctx, t, WithKeyVal(key, val), WithExpiration(expiration))
	handler.Set()
	handler.StrLen() // strlen: 7

	handler.GetRange(0, 6)    // 5reafsf
	handler.GetRange(-7, -1)  // 5reafsf
	handler.GetRange(-1, 6)   // f
	handler.GetRange(1, -1)   // reafsf
	handler.GetRange(4, -1)   // fsf
	handler.GetRange(-1, 4)   //
	handler.GetRange(-10, 10) // 5reafsf
	/*
		 5  r  e  a  f  s  f
		 0  1  2  3  4  5  6
		-7 -6 -5 -4 -3 -2 -1
	*/
}

// TestSetRange
// @Description: 不支持负数索引
// @param t
func TestSetRange(t *testing.T) {
	key, val, expiration := "number", "5reafsf", Three
	handler := NewRedisString(ctx, t, WithKeyVal(key, val), WithExpiration(expiration))
	handler.Set()
	handler.StrLen()

	handler.SetRange(0, "hello world") // 自动拓展被修改的字符串
	handler.Get()
	/*Output
	11
	get number
	"hello world"
	*/

	handler.SetRange(-1, "plus") // ERR offset is out of range
	handler.Get()

	handler.SetRange(15, "plus") // 在值里边填充空字节
	handler.Get()
	/*Output
	19
	get number
	"hello world    plus"
	*/
}

// TestAppend
// @Description:
// @param t
func TestAppend(t *testing.T) {
	key, val, expiration := "number", "5reafsf", Three
	handler := NewRedisString(ctx, t, WithKeyVal(key, val), WithExpiration(expiration))
	handler.Set()
	handler.Append("plus")
	handler.Get()
	/*Output
	set number "5reafsf" EX 3
	OK
	11
	get number
	"5reafsfplus"
	*/

	// 处理不存在的键
	key, val, expiration = "number1", "5reafsf", Three
	handler = NewRedisString(ctx, t, WithKeyVal(key, val), WithExpiration(expiration))
	//handler.Set()
	handler.Append("plus")
	handler.Get()
	/*Output
	4
	get number1
	174: "plus"
	*/
}

var kvMap = map[string]string{
	"number1": "10086",
	"number2": "+894",
	"number3": "-123",
	"number4": "+2.56",
	"number5": "-5.12",
	"number6": "12345678901234567890",
	"number7": "3.14e5",
	"number8": "one",
	"number9": "123abc",
}

func TestIncrByAndDecrBy(t *testing.T) {
	expiration := One
	for key, val := range kvMap {
		handler := NewRedisString(ctx, t, WithKeyVal(key, val), WithExpiration(expiration))
		handler.Set()

		handler.IncrBy(10)
		handler.DecrBy(10)
		t.Log()
	}
	/*Output
	set number2 "+894" EX 1
	OK
	ERR value is not an integer or out of range
	ERR value is not an integer or out of range

	set number3 "-123" EX 1
	OK
	-113
	-123

	set number5 "-5.12" EX 1
	OK
	ERR value is not an integer or out of range
	ERR value is not an integer or out of range

	set number6 "12345678901234567890" EX 1
	OK
	ERR value is not an integer or out of range
	ERR value is not an integer or out of range

	set number1 "10086" EX 1
	OK
	10096
	10086

	set number7 "3.14e5" EX 1
	OK
	ERR value is not an integer or out of range
	ERR value is not an integer or out of range

	set number8 "one" EX 1
	OK
	ERR value is not an integer or out of range
	ERR value is not an integer or out of range

	set number9 "123abc" EX 1
	OK
	ERR value is not an integer or out of range
	ERR value is not an integer or out of range

	set number4 "+2.56" EX 1
	OK
	ERR value is not an integer or out of range
	ERR value is not an integer or out of range

	*/
}

// TestIncrByAndDecrByForIllegal
// @Description:
// @param t
func TestIncrByAndDecrByForIllegal(t *testing.T) {
	key := "number"
	handler := NewRedisString(ctx, t, WithKeyVal(key, "null"))
	handler.Get()

	handler.IncrBy(12) // 处理不存在的键
}

// TestIncrAndDecr
// @Description:
// @param t
func TestIncrAndDecr(t *testing.T) {
	key, val, expiration := "num", "10", One
	handler := NewRedisString(ctx, t, WithKeyVal(key, val), WithExpiration(expiration))
	handler.Set()

	handler.Incr()
	handler.Decr()
	/*Output
	set num "10" EX 1
	OK
	11
	10
	*/
}

func TestIncrByFloat(t *testing.T) {
	key, val, expiration := "n", "3.14", One
	handler := NewRedisString(ctx, t, WithKeyVal(key, val), WithExpiration(expiration))

	// 即可用于浮点数、也可用于整数
	// 增量可以是浮点数，也可以是整数
	handler.IncrByFloat(1) // 不存在的键 默认初始化为 0
	handler.IncrByFloat(-1)
	handler.IncrByFloat(3.14)
	handler.IncrByFloat(-3.14)

	handler = NewRedisString(ctx, t, WithKeyVal(key, "0.0123456789123456789123456789"), WithExpiration(expiration))
	handler.Set()
	handler.IncrByFloat(0) // 保留计算结果小数点后 17 位数字，超过截断
	/*Output
	1
	0
	3.14
	0
	set n "0.0123456789123456789123456789" EX 1
	OK
	0.01234567891234568
	*/
}
