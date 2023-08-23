package lib

import (
	"testing"
	"time"
)

func TestUnix(t *testing.T) {
	now := time.Now().Unix()
	date := time.Unix(now, 0)
	t.Log(now, "\n", date)

	time.Sleep(time.Second * 2)

	nowNsec := time.Now().UnixNano()
	date = time.Unix(0, nowNsec)
	t.Log(nowNsec, "\n", date)

}

func TestParseDuration(t *testing.T) {
	now := time.Now()
	m1, _ := time.ParseDuration("-24h")
	yesterday := now.Add(m1)
	t.Log(now, yesterday)
}
