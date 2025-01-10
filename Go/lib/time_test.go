package main

import (
	"encoding/json"
	"testing"
	"time"
)

func TestTimeMarshal(t *testing.T) {
	var unixStamp CustomTime
	unixStamp.Time = time.Now()

	str, _ := unixStamp.MarshalJSON()
	t.Log(string(str))

	var tmp CustomTime
	tmp.UnmarshalJSON(str)

	str1, _ := tmp.MarshalJSON()
	t.Log(string(str1))
}

type CustomTime struct {
	time.Time
}

// MarshalJSON 自定义序列化方法
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(ct.Format("2006-01-02 15:04:05"))
}

// UnmarshalJSON 自定义反序列化方法
func (ct *CustomTime) UnmarshalJSON(data []byte) error {
	var strTime string
	if err := json.Unmarshal(data, &strTime); err != nil {
		return err
	}
	parsedTime, err := time.Parse("2006-01-02 15:04:05", strTime)
	if err != nil {
		parsedTime, err = time.Parse("2006-01-02 12:00", strTime)
		if err != nil {
			return err
		}
	}
	ct.Time = parsedTime
	return nil
}
