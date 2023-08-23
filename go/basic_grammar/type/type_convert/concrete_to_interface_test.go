package type_convert

import "testing"

type ITest interface {
	Hello()
}

type Yesterday struct {
}

func (y *Yesterday) Hello() {}

func TestNotPoint(t *testing.T) {
	//var p ITest
	//p = new(Yesterday)
}
