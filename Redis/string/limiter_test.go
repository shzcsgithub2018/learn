package String

import (
	"github.com/shzgithub2018/learn/Util/concurrency"
	"testing"
)

// TestLimiterMulti
// @Description:
// @param t
func TestLimiterMulti(t *testing.T) {
	cnt1, cnt2 := NewCounter(ctx, "cnt1"), NewCounter(ctx, "cnt2")
	cnt1.Reset()
	cnt2.Reset()

	key := "cnt"
	limiter := NewLimiter(ctx, key)
	limiter.SetMaxExecuteTimes(10000)

	concurrency.UnSafe(20, 10000, func() {
		if limiter.StillValidToExecute() {
			cnt1.Incr()
		} else {
			cnt2.Incr()
		}
	})

	t.Log(limiter.RemainingExecuteTimes())
	t.Log("success:", cnt1.Get(), " fail:", cnt2.Get())
	/*Output
	-190000
	success: 10000  fail: 190000
	*/
}
