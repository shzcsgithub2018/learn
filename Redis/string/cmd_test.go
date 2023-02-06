package string

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
	handler := dal.NewRedisString(ctx, t, dal.WithKeyVal(key, val), dal.WithExpiration(expiration))
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
	handler := dal.NewRedisString(ctx, t, dal.WithKeyVal(key, val), dal.WithExpiration(expiration))

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
	handler := dal.NewRedisString(ctx, t, dal.WithKeyVal(key, val), dal.WithExpiration(expiration))

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
	handler := dal.NewRedisString(ctx, t, dal.WithKeyVal(key, val), dal.WithExpiration(expiration))

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
	handler := dal.NewRedisString(ctx, t, dal.WithKeyVal(key, val), dal.WithExpiration(expiration))

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
	handler := dal.NewRedisString(ctx, t, dal.WithKeyVal(key, val), dal.WithExpiration(expiration))

	handler.Set()

	val, expiration = "Hello World Plus", redis.KeepTTL
	handler = dal.NewRedisString(ctx, t, dal.WithKeyVal(key, val), dal.WithKeepTTL())

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
	handler := dal.NewRedisString(ctx, t, dal.WithKeyVal(key, val))
	handler.GetSet()

	val += " Plus"
	handler = dal.NewRedisString(ctx, t, dal.WithKeyVal(key, val))
	handler.GetSet()

	val += " Plus"
	handler = dal.NewRedisString(ctx, t, dal.WithKeyVal(key, val))
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

	handler := dal.NewRedisString(ctx, t, dal.WithMKeyVal(kvMap))

	handler.MSetNX()
	handler.MSetNX()

	kvMap["a"] = "b"
	handler = dal.NewRedisString(ctx, t, dal.WithMKeyVal(kvMap))
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
	handler := dal.NewRedisString(ctx, t, dal.WithKeyVal(key, val), dal.WithExpiration(expiration))
	handler.Set()
	handler.StrLen()
	/*Output
	OK
	7
	*/
}
