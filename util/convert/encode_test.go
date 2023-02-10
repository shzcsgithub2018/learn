package convert

import (
	"strings"
	"testing"
)

func TestBase10To36(t *testing.T) {
	dataMap := map[int64]string{
		0:                "0",
		10:               "A",
		13123123:         "7T9V7",
		1034124231431432: "A6KE433ITK",
	}

	for key, val := range dataMap {
		if strings.Compare(Base10To36(key), val) != 0 {
			t.Error(Base10To36(key), val, "ERR")
		}
	}
}
