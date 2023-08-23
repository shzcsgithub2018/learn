package channel

import (
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTicker(time.Second * 5)
	timer2 := time.NewTimer(time.Second * 10)

	for {
		flag := false
		select {
		case <-timer.C:
			t.Log(time.Now())
		case <-timer2.C:
			t.Log("timer timer!")
			flag = true
		}
		if flag {
			break
		}
	}
}
