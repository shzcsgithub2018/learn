package hash

import (
	"context"
	"github.com/redis/go-redis/v9"
	String "github.com/shzgithub2018/learn/Redis/string"
	"github.com/shzgithub2018/learn/Util/convert"
	"testing"
)

/*
what: 比较短的网址
	- 过程
		当我们访问A时，会经过mtw.so的域名解析最终会请求该域名下的某一个http接口（😎后段程序员都懂，就不做多解释了）
		这个接口会拿到url后的参数 6kK03S
		通过这个参数后端会定位到一个长链接，也就是原始链接，最后通过重定向到原始链接（301/302按需使用），至此就结束了
		（重定向是浏览器自动完成的，区别与转发）

	- 应用
		- 新浪微博发微博有字数限制，微博会自动把我们的长网址转成短网址
		- 三发发给用户的手机短信，网址也都是转换后的短网址，否则会影响排版，让人感觉不专业，降低点击率。
        - 分享的一些链接，也是短网址（抖音，淘宝， 京东竟然不是。。。）
why：
	- 缩短网址
	- 收集有关点击数据
	- 收集信任度
	- 链接重定向
	- 防红防屏蔽
	- A/B 测试
*/

type ShortUrl struct {
	ctx         context.Context
	t           *testing.T
	cli         *redis.Client
	IdGenerator *String.IdGenerator
}

const (
	ID_COUNTER = "ShortUrl::id_counter"
	URL_HASH   = "ShortUrl:url_hash"
)

func NewShortUrl(ctx context.Context, t *testing.T, cli *redis.Client) *ShortUrl {
	IdGenerator := String.NewIdGenerator(ctx, ID_COUNTER)
	IdGenerator.Init(10000)
	return &ShortUrl{
		ctx:         ctx,
		t:           t,
		cli:         cli,
		IdGenerator: IdGenerator,
	}
}

func (s *ShortUrl) Shorten(targetUrl string) string {
	// 生成全局 id
	id := s.IdGenerator.GenId()
	// 缩短 id
	shortId := convert.Base10To36(id)
	// 保存
	val := s.cli.HSet(s.ctx, URL_HASH, shortId, targetUrl).Val()
	if val == 1 {
		s.t.Log("gen and save success.")
	} else {
		s.t.Log("gen and update success.")
	}
	return shortId
}

func (s *ShortUrl) Restore(shortId string) string {
	return s.cli.HGet(s.ctx, URL_HASH, shortId).Val()
}
