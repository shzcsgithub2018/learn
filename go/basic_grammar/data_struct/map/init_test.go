package map_test

import "testing"

func TestMap(t *testing.T) {
	m := map[int64]string{}
	m[1] = "efaf"
	t.Log(m)
}
