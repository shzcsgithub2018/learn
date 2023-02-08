package hash

import (
	"testing"
	"time"
)

func TestGenerateToken(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(GenerateToken())
	}
}

func TestTime(t *testing.T) {
	tt := time.Now()
	tt.Unix()
	t.Log(tt, "\n", tt.Unix(), "\n", tt.Add(10).Unix())
}

func TestLoginSession(t *testing.T) {
	userId := int64(3241243241234)
	handle := NewLoginSession(ctx, userId)
	var token string

	handle.Validate(token) // SESSION_NOT_LOGIN
	token = handle.Create()
	handle.Validate(token)          // SESSION_TOKEN_CORRECT
	handle.Validate(token + "plus") // SESSION_TOKEN_INCORRECT
	time.Sleep(time.Duration(DefaultTimeout+1) * time.Second)
	handle.Validate(token) // SESSION_EXPIRED

	handle.Destroy() // Destroy
}
