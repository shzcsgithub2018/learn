package key_words

import "testing"

func TestSwitch(t *testing.T) {
	a := 10
	switch a {
	case 0:
		t.Log("This is 0. ")
		return
	case 1:
		t.Log("This is 1. ")
		return
	}
	t.Log("Not")
}
