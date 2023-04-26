package marshal

import (
	"encoding/json"
	"testing"
)

func TestUnMarshal(t *testing.T) {
	str := `{
		"old_send_msg_switch_key":false,
		"new_send_msg_switch_key":false,
		"diff_switch_key": true
	}`
	tccSwitch := TccSwitch{}
	err := json.Unmarshal([]byte(str), tccSwitch)
	if err != nil {
		t.Error(err)
	}
	t.Log(tccSwitch)
}

type TccSwitch struct {
	OldKey        bool `json:"old_send_msg_switch_key"`
	NewKey        bool `json:"new_send_msg_switch_key"`
	DiffSwitchKey bool `json:"diff_switch_key"`
}
