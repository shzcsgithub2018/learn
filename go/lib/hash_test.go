package lib

import (
	"hash/fnv"
	"testing"
)

func hashString(s string, mod int) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32() % uint32(mod))
}

func TestHashFnv(t *testing.T) {
	str := "dsfasdfsafsdfsfsafsb"
	t.Log(hashString(str, 50))
}
