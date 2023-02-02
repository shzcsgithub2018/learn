package gorounti

import (
	"fmt"
	"testing"
	"time"
)

func TestTmp(t *testing.T) {
	go func() {
		type aa struct {
			a int32
			b int64
		}
		tt := &aa{a: 10, b: 100}
		go func() {
			for tt.a < 30 {
				tt.a++
				t.Log(tt.a, tt.b)
				fmt.Println(&tt)
			}
		}()
		fmt.Println(&tt)
	}()
	time.Sleep(time.Duration(10) * time.Second)
}
