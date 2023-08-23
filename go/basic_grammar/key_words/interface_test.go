package key_words

import "testing"

func TestInterface(t *testing.T) {
	type component interface {
		A() bool
	}

	type composite struct {
	}
}

type composite struct {
}

func (composite2 composite) A() bool {
	return false
}

// 接口的本质
// https://chenhe.me/post/pointer-and-interface-in-go/#%E6%8E%A5%E5%8F%A3%E7%9A%84%E6%9C%AC%E8%B4%A8

// 可以路由到  pass_by_value.receiver_test.go
