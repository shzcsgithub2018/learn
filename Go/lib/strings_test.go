package main

import (
	"strings"
	"testing"
)

func TestStrings(t *testing.T) {
	url := "https://smart-customer-1258344696.cos.ap-beijing.myqcloud.com/knowledge-library/dd98c818-698d-42a6-8862-371a5fce1c39/test.txt"
	list := strings.Split(url, ".myqcloud.com/")
	t.Log(list)
}
