package String

import (
	"github.com/shzgithub2018/learn/util/concurrency"
	"testing"
)

func TestIdGenerator(t *testing.T) {
	handler := NewIdGenerator(ctx, "articleId")
	handler.Reset()
	handler.Init(0)

	idList := [10001]int64{}

	concurrency.UnSafe(100, 100, func() {
		idList[handler.GenId()] = 1
	})

	for i := int64(1); i <= 100*100; i++ {
		if idList[i] != 1 {
			t.Error("Warning")
		}
	}
	/*Output
	 */
}
