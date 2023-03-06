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
