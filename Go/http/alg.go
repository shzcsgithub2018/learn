package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type ReqBody struct {
	SessionList [][]string `json:"sessions"`
}

type RespBody struct {
	Results []string `json:"results"`
}

func SQUClass(ctx context.Context, sessionList [][]string) ([]string, error) {
	body := &ReqBody{
		SessionList: sessionList,
	}
	bodyStr, _ := json.Marshal(body)

	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
		}).
		SetBody(bodyStr).
		Post("ai-customer-squ-recognition.production.polaris:9002/squ_class")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, errors.New(fmt.Sprintf("Http call fail. code:%d", resp.StatusCode()))
	}

	fmt.Println(ctx, "GetVector http resp body:%s", string(resp.Body()))

	res := RespBody{}
	if err = json.Unmarshal(resp.Body(), &res); err != nil {
		return nil, err
	}

	return res.Results, nil
}
