package hash

import (
	"github.com/shzgithub2018/learn/Redis/dal"
	"testing"
)

func TestShortUrl(t *testing.T) {
	handler := NewShortUrl(ctx, t, dal.GetRedisCli(ctx))

	shortUrl := handler.Shorten("https://www.baidu.com")
	t.Log(shortUrl)

	targetUrl := handler.Restore(shortUrl)
	t.Log(targetUrl)
	/*Output
	gen and save success.
	7PV
	https://www.baidu.com
	*/
}
