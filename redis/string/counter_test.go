package String

import (
	"github.com/shzgithub2018/learn/util/concurrency"
	"testing"
)

// TestBasic
// @Description:
// @param t
func TestCounterBasic(t *testing.T) {
	key := "cnt"
	counter := NewCounter(ctx, key)
	counter.Reset()

	t.Log(counter.Incr())                  // 1
	t.Log(counter.Incr(WithModifyNum(10))) // 11

	t.Log(counter.Decr())                  // 10
	t.Log(counter.Decr(WithModifyNum(10))) // 0

	counter.Incr(WithModifyNum(100))
	t.Log(counter.Reset()) // 100
	t.Log(counter.Get())   // 0
}

// TestMulti
// @Description:  并发测试 redis 计数器的读写原子性
// @param t
func TestCounterMulti(t *testing.T) {
	key, cnt := "cnt", 0
	counter := NewCounter(ctx, key)
	counter.Reset()

	concurrency.UnSafe(250, 1000, func() {
		counter.Incr()
		cnt++
	})
	concurrency.UnSafe(250, 1000, func() {
		counter.Decr()
		cnt--
	})

	t.Log(counter.Get(), cnt)
	/*Output
	0
	-31
	*/
}
