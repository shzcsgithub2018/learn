package string

import (
	"github.com/redis/go-redis/v9"
	"github.com/shzgithub2018/learn/Redis/dal"
	"testing"
	"time"
)

var m = map[string]string{
	"message":  "hello world",
	"homepage": "redis.io",
	"price":    "2.56",
	"number":   "10086",
	"book":     "The Design and Implementation of Redis",
}

const ZERO = 0

/**
 * TestSet
 * @Description: "SET",
 * @param t
 * @Output:
	* key:message	val：hello world
	* hello world
*/
func TestSetGET(t *testing.T) {
	for key, val := range m {
		t.Log("key:" + key + "	val：" + val)
		if err := dal.GetRedisCli(ctx).Set(ctx, key, val, ZERO).Err(); err != nil {
			t.Error(err)
		}
		if str, err := dal.GetRedisCli(ctx).Get(ctx, key).Result(); err != nil {
			t.Error(err)
		} else {
			t.Log(str)
		}
		t.Log()

		break
	}
}

/**
 * TestExpiration
 * @Description:
 * @param t
 * @Output:
 	* key:message	val：hello world
	* hello world
	* After 1 second……
	* redis: nil
*/
func TestExpiration(t *testing.T) {
	expiration := time.Duration(1) * time.Second
	for key, val := range m {
		t.Log("key:" + key + "	val：" + val)
		if err := dal.GetRedisCli(ctx).Set(ctx, key, val, expiration).Err(); err != nil {
			t.Error(err)
		}
		if str, err := dal.GetRedisCli(ctx).Get(ctx, key).Result(); err != nil {
			t.Error(err)
		} else {
			t.Log(str)
		}
		t.Log("After 1 second……")
		time.Sleep(expiration * 2)
		if str, err := dal.GetRedisCli(ctx).Get(ctx, key).Result(); err != nil {
			t.Error(err)
		} else {
			t.Log(str)
		}
		t.Log()
		break
	}
}

/**
 * TestKeepTTL
 * @Description:
 * @param t
 * @Output:
   	* key:name	val1：Peter
	* Peter
	* set expiration is KeepTTL
	* key:name	val2：Bob
	* Bob
	* After 4 second……
	* redis: nil
*/
func TestKeepTTL(t *testing.T) {
	expiration := time.Duration(3) * time.Second
	var (
		key  = "name"
		val1 = "Peter"
		val2 = "Bob"
	)

	t.Log("key:" + key + "	val1：" + val1)
	if err := dal.GetRedisCli(ctx).Set(ctx, key, val1, expiration).Err(); err != nil {
		t.Error(err)
	}
	if str, err := dal.GetRedisCli(ctx).Get(ctx, key).Result(); err != nil {
		t.Error(err)
	} else {
		t.Log(str)
	}

	expiration = redis.KeepTTL
	t.Log("set expiration is KeepTTL")
	t.Log("key:" + key + "	val2：" + val2)
	if err := dal.GetRedisCli(ctx).Set(ctx, key, val2, expiration).Err(); err != nil {
		t.Error(err)
	}
	if str, err := dal.GetRedisCli(ctx).Get(ctx, key).Result(); err != nil {
		t.Error(err)
	} else {
		t.Log(str)
	}
	t.Log("After 4 second……")
	time.Sleep(time.Duration(4) * time.Second)
	if str, err := dal.GetRedisCli(ctx).Get(ctx, key).Result(); err != nil {
		t.Error(err)
	} else {
		t.Log(str)
	}
	t.Log()
}

/**
* TestSetNx
* @Description:
* @param
* @Output:
   * key:name1	val1：Bob
   * setNX success
   * setNX fail
*/
func TestSetNX(t *testing.T) {
	var (
		key        = "name1"
		val        = "Bob"
		expiration = time.Duration(3) * time.Second
	)

	t.Log("key:" + key + "	val1：" + val)
	cmd := dal.GetRedisCli(ctx).SetNX(ctx, key, val, expiration)
	if cmd.Err() != nil {
		t.Error(cmd.Err())
	}
	if cmd.Val() {
		t.Log("setNX success")
	} else {
		t.Log("setNX fail")
	}

	// double setNX
	cmd = dal.GetRedisCli(ctx).SetNX(ctx, key, val, expiration)
	if cmd.Err() != nil {
		t.Error(cmd.Err())
	}
	if cmd.Val() {
		t.Log("setNX success")
	} else {
		t.Log("setNX fail")
	}
}

/*
*
  - TestSetXX
  - @Description:
  - @param t
  - @Output:
    key:name1	val1：Bob
    setXX fail
*/
func TestSetXX(t *testing.T) {
	var (
		key        = "name1"
		val        = "Bob"
		expiration = time.Duration(3) * time.Second
	)

	t.Log("key:" + key + "	val1：" + val)
	cmd := dal.GetRedisCli(ctx).SetXX(ctx, key, val, expiration)
	if cmd.Err() != nil {
		t.Error(cmd.Err())
	}
	if cmd.Val() {
		t.Log("setXX success")
	} else {
		t.Log("setXX fail")
	}
}
