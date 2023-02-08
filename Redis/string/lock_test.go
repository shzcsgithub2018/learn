package String

import (
	"github.com/shzgithub2018/learn/Util/concurrency"
	"testing"
	"time"
)

// TestCache
// @Description:
// @param t
func TestLock(t *testing.T) {
	key := "article::Mary:001"
	lock := NewLock(ctx, key)
	lock.Release()
	var noLockNum, lockNum int64

	startUnLock := time.Now()
	concurrency.UnSafe(100, 100, func() {
		Incr(&noLockNum)
	})
	endUnLock := time.Now()

	startLock := time.Now()
	concurrency.UnSafe(100, 100, func() {
		for !lock.Acquire() {
		}
		Incr(&lockNum)
		lock.Release()
	})
	endLock := time.Now()

	t.Log("noLockNum:", noLockNum, "	lockNum:", lockNum)
	t.Log("unLockTime:", endUnLock.UnixMilli()-startUnLock.UnixMilli(), "ms")
	t.Log("lockTime:", endLock.UnixMilli()-startLock.UnixMilli(), "ms")
	/*Output
	noLockNum: 9555 	lockNum: 10000
	unLockTime: 1 ms
	lockTime: 9218 ms
	*/
}

func Incr(n *int64) {
	*n++
}
