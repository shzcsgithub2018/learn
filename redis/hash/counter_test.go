package hash

import "testing"

func TestCounterBasic(t *testing.T) {
	var (
		key         = "page_view_counters"
		hashKeyList = []string{
			"/user/peter",
			"/product/256",
			"/product/512",
		}
	)
	counter := NewCounter(ctx, key)

	for _, hashKey := range hashKeyList {
		counter.Incr(hashKey)
		counter.Decr(hashKey)
		counter.Incr(hashKey, WithModifyNum(10))
	}

	counter.Incr(hashKeyList[0], WithModifyNum(100))
	counter.Reset(hashKeyList[0])
	counter.Get(hashKeyList[0])
	/*Output
	del  page_view_counters
	HIncrBy page_view_counters /user/peter  1
	1
	HIncrBy page_view_counters /user/peter  -1
	0
	HIncrBy page_view_counters /user/peter  10
	10
	HIncrBy page_view_counters /product/256  1
	1
	HIncrBy page_view_counters /product/256  -1
	0
	HIncrBy page_view_counters /product/256  10
	10
	HIncrBy page_view_counters /product/512  1
	1
	HIncrBy page_view_counters /product/512  -1
	0
	HIncrBy page_view_counters /product/512  10
	10
	HIncrBy page_view_counters /user/peter  100
	110
	HGet  page_view_counters /user/peter
	110
	HSet page_view_counters /user/peter 0
	HGet  page_view_counters /user/peter
	0
	*/
}
