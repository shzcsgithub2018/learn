package hash

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/shzgithub2018/learn/Redis/dal"
	"log"
	"math/rand"
	"strings"
	"time"
)

/*
 登录会话：
	为方便用户，网站一般都会为已登录的用户生成一个加密令牌，然后把这个令牌分别保存在服务端和客户端，
	之后每当用户再次访问该网站的时候，网站就可以通过验证客户端提交的令牌来确认用户的身份，从而使得用户不必重复地执行登录操作。
   	另外，防止用户长时间不输入密码而遗忘密码，令牌需要有过期期限

*/

const (
	DefaultTimeout = 10

	SessionTokenHash    = "session::token"
	SessionExpireTsHash = "session::expire_timestamp"

	SessionNotLogin       = "SESSION_NOT_LOGIN"
	SessionExpired        = "SESSION_EXPIRED"
	SessionTokenCorrect   = "SESSION_TOKEN_CORRECT"
	SessionTokenIncorrect = "SESSION_TOKEN_INCORRECT"
)

func GenerateToken() string {
	randStr := fmt.Sprint(rand.Int63())
	return string(sha256.New().Sum([]byte(randStr)))
}

type LoginSession struct {
	ctx context.Context
	key string
}

func NewLoginSession(ctx context.Context, userId int64) *LoginSession {
	key := fmt.Sprint(userId)
	dal.GetRedisCli(ctx).Del(ctx, key)
	return &LoginSession{
		ctx: ctx,
		key: key,
	}
}

func (h *LoginSession) Create() string {
	token := GenerateToken()
	expireTime := fmt.Sprint(time.Now().Unix() + DefaultTimeout)
	log.Println("token:", token, " ExpireTime:", expireTime, "s")

	dal.GetRedisCli(h.ctx).HSet(h.ctx, h.key, SessionTokenHash, token, SessionExpireTsHash, expireTime)

	return token
}

func (h *LoginSession) Validate(token string) bool {
	isExists := dal.GetRedisCli(h.ctx).HExists(h.ctx, h.key, SessionTokenHash).Val()
	if !isExists {
		log.Println(SessionNotLogin)
		return false
	}
	serverToken := dal.GetRedisCli(h.ctx).HGet(h.ctx, h.key, SessionTokenHash).Val()
	expireTime := dal.GetRedisCli(h.ctx).HGet(h.ctx, h.key, SessionExpireTsHash).Val()
	currTime := fmt.Sprint(time.Now().Unix())

	if strings.Compare(serverToken, token) != 0 {
		log.Println(SessionTokenIncorrect)
		return false
	}
	if strings.Compare(expireTime, currTime) < 0 {
		log.Println(SessionExpired)
		return false
	}
	log.Println(SessionTokenCorrect)
	return true
}

func (h *LoginSession) Destroy() {
	dal.GetRedisCli(h.ctx).HDel(h.ctx, h.key, SessionTokenHash)
	dal.GetRedisCli(h.ctx).HDel(h.ctx, h.key, SessionExpireTsHash)
	log.Println("Destroy")
}
