package string

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/shzgithub2018/learn/Redis/dal"
	"testing"
	"time"
)

type RedisString struct {
	ctx        context.Context
	t          *testing.T
	key, val   string
	kvMap      map[string]interface{}
	expiration time.Duration
	keepTTL    bool
}

type StrOption func(r *RedisString)

func NewRedisString(ctx context.Context, t *testing.T, opts ...StrOption) *RedisString {
	redisString := &RedisString{
		ctx:   ctx,
		t:     t,
		kvMap: make(map[string]interface{}, 1),
	}

	for _, opt := range opts {
		opt(redisString)
	}
	return redisString
}

func WithKeyVal(key, val string) StrOption {
	return func(r *RedisString) {
		r.key = key
		r.val = val
	}
}

func WithExpiration(expiration time.Duration) StrOption {
	return func(r *RedisString) {
		r.expiration = expiration
	}
}

func WithKeepTTL() StrOption {
	return func(r *RedisString) {
		r.keepTTL = true
	}
}

func WithMKeyVal(kvMap map[string]interface{}) StrOption {
	return func(r *RedisString) {
		for key, val := range kvMap {
			r.kvMap[key] = val
		}
	}
}

// Set
// @Description:
// @receiver r
func (r *RedisString) Set() {
	buildCmd := func() string {
		cmd := "set " + r.key + " \"" + r.val + "\""
		var ex string
		if r.keepTTL {
			ex = " KeepTTL"
		} else {
			if r.expiration != time.Duration(0) {
				ex = r.expiration.String()
				ex = ex[:len(ex)-1]
			}
		}
		return cmd + " EX " + ex
	}

	r.t.Log(buildCmd())
	var cmd *redis.StatusCmd
	if r.keepTTL {
		cmd = dal.GetRedisCli(r.ctx).Set(r.ctx, r.key, r.val, redis.KeepTTL)
	} else {
		cmd = dal.GetRedisCli(r.ctx).Set(r.ctx, r.key, r.val, r.expiration)
	}
	if err := cmd.Err(); err != nil {
		r.t.Log(err)
	} else {
		r.t.Log(cmd.Val())
	}
}

// SetNX
// @Description:
// @receiver r
func (r *RedisString) SetNX() {
	buildCmd := func() string {
		cmd := "set " + r.key + " \"" + r.val + "\" NX"
		var ex string
		if r.keepTTL {
			ex = " KeepTTL"
		} else {
			if r.expiration != time.Duration(0) {
				ex = r.expiration.String()
				ex = ex[:len(ex)-1]
			}
		}
		return cmd + " EX " + ex
	}

	r.t.Log(buildCmd())
	var cmd *redis.BoolCmd
	if r.keepTTL {
		cmd = dal.GetRedisCli(r.ctx).SetNX(r.ctx, r.key, r.val, redis.KeepTTL)
	} else {
		cmd = dal.GetRedisCli(r.ctx).SetNX(r.ctx, r.key, r.val, r.expiration)
	}
	if err := cmd.Err(); err != nil {
		r.t.Log(err)
		return
	}
	if cmd.Val() {
		r.t.Log("success")
	} else {
		r.t.Log("fail")
	}
}

// SetXX
// @Description:
// @receiver r
func (r *RedisString) SetXX() {
	buildCmd := func() string {
		cmd := "set " + r.key + " \"" + r.val + "\" XX"
		var ex string
		if r.keepTTL {
			ex = " KeepTTL"
		} else {
			if r.expiration != time.Duration(0) {
				ex = r.expiration.String()
				ex = ex[:len(ex)-1]
			}
		}
		return cmd + " EX " + ex
	}

	r.t.Log(buildCmd())
	var cmd *redis.BoolCmd
	if r.keepTTL {
		cmd = dal.GetRedisCli(r.ctx).SetXX(r.ctx, r.key, r.val, redis.KeepTTL)
	} else {
		cmd = dal.GetRedisCli(r.ctx).SetXX(r.ctx, r.key, r.val, r.expiration)
	}
	if err := cmd.Err(); err != nil {
		r.t.Log(err)
		return
	}
	if cmd.Val() {
		r.t.Log("success")
	} else {
		r.t.Log("fail")
	}
}

// Get
// @Description:
// @receiver r
func (r *RedisString) Get() {
	r.t.Log("get " + r.key)
	cmd := dal.GetRedisCli(r.ctx).Get(r.ctx, r.key)
	if err := cmd.Err(); err != nil {
		r.t.Log(err)
	} else {
		r.t.Log("\"" + cmd.Val() + "\"")
	}
}

// MSetNX
// @Description:
// @receiver r
func (r *RedisString) MSetNX() {
	cmdStr := "MSetNX  "
	for key, val := range r.kvMap {
		cmdStr += key + " \"" + val.(string) + "\" "
	}
	r.t.Log(cmdStr)
	cmd := dal.GetRedisCli().MSetNX(r.ctx, r.kvMap)
	if err := cmd.Err(); err != nil {
		r.t.Error(err)
	}
	if cmd.Val() {
		r.t.Log("success")
	} else {
		r.t.Log("fail")
	}
}

// GetSet
// @Description:
// @receiver r
func (r *RedisString) GetSet() {
	r.t.Log("getSet " + r.key + " \"" + r.val + "\"")
	cmd1 := dal.GetRedisCli(r.ctx).GetSet(r.ctx, r.key, r.val)
	if err := cmd1.Err(); err != nil {
		r.t.Error(err)
		return
	}
	r.t.Log(cmd1.Val())
}

// StrLen
// @Description:
// @receiver r
func (r *RedisString) StrLen() {
	cmd := dal.GetRedisCli(r.ctx).StrLen(r.ctx, r.key)
	if err := cmd.Err(); err != nil {
		r.t.Log(err)
	} else {
		r.t.Log("strlen:", cmd.Val())
	}
}

// GetRange
// @Description: [start, end] O(N)
// @receiver r
// @param start
// @param end
func (r *RedisString) GetRange(start, end int64) {
	cmd := dal.GetRedisCli(r.ctx).GetRange(r.ctx, r.key, start, end)
	if err := cmd.Err(); err != nil {
		r.t.Log(err)
		return
	}
	r.t.Log(cmd.Val())
}

// SetRange
// @Description:
// @receiver r
// @param start
// @param substr
func (r *RedisString) SetRange(start int64, substr string) {
	cmd := dal.GetRedisCli(r.ctx).SetRange(r.ctx, r.key, start, substr)
	if err := cmd.Err(); err != nil {
		r.t.Log(err)
		return
	}
	r.t.Log(cmd.Val())
}

// Append
// @Description:
// @receiver r
// @param suffix
func (r *RedisString) Append(suffix string) {
	cmd := dal.GetRedisCli(r.ctx).Append(r.ctx, r.key, suffix)
	if err := cmd.Err(); err != nil {
		r.t.Log(err)
		return
	}
	r.t.Log(cmd.Val())
}

// IncrBy
// @Description:
// @receiver r
// @param val
func (r *RedisString) IncrBy(val int64) {
	cmd := dal.GetRedisCli(r.ctx).IncrBy(r.ctx, r.key, val)
	if err := cmd.Err(); err != nil {
		r.t.Log(err)
		return
	}
	r.t.Log(cmd.Val())
}

// DecrBy
// @Description:
// @receiver r
// @param val
func (r *RedisString) DecrBy(val int64) {
	cmd := dal.GetRedisCli(r.ctx).DecrBy(r.ctx, r.key, val)
	if err := cmd.Err(); err != nil {
		r.t.Log(err)
		return
	}
	r.t.Log(cmd.Val())
}

// Incr
// @Description:
// @receiver r
func (r *RedisString) Incr() {
	cmd := dal.GetRedisCli(r.ctx).Incr(r.ctx, r.key)
	if err := cmd.Err(); err != nil {
		r.t.Log(err)
		return
	}
	r.t.Log(cmd.Val())
}

// Decr
// @Description:
// @receiver r
func (r *RedisString) Decr() {
	cmd := dal.GetRedisCli(r.ctx).Decr(r.ctx, r.key)
	if err := cmd.Err(); err != nil {
		r.t.Log(err)
		return
	}
	r.t.Log(cmd.Val())
}

// IncrByFloat
// @Description:
// @receiver r
// @param val
func (r *RedisString) IncrByFloat(val float64) {
	cmd := dal.GetRedisCli(r.ctx).IncrByFloat(r.ctx, r.key, val)
	if err := cmd.Err(); err != nil {
		r.t.Log(err)
		return
	}
	r.t.Log(cmd.Val())
}
