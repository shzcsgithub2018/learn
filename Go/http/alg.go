package api

import (
	"context"
	"encoding/json"
	"fmt"
	"git.woa.com/ams-ai-biz/ai-customer/internal/errs"
	"git.woa.com/ams-ai-biz/ai-customer/internal/log"
	"git.woa.com/ams-ai-biz/ai-customer/internal/rainbow"
	"github.com/go-resty/resty/v2"
	"time"
)

type VectorReqBody struct {
	QueryID  string `json:"query_id"`
	Text     string `json:"text"`
	K        int    `json:"k"`
	EmbIndex string `json:"emb_index"`
}

type VectorRespBody struct {
	Retcode int    `json:"retcode"`
	UUID    string `json:"uuid"`
	Msg     string `json:"msg"`
	Results []struct {
		Index  string  `json:"index"`
		Value  string  `json:"value"`
		Metric float64 `json:"metric"`
	} `json:"results"`
}

type VectorSearchConfig struct {
	Url                 string  `json:"url"`
	EmbIndex            string  `json:"emb_index"`
	Token               string  `json:"token"`
	Wsid                string  `json:"wsid"`
	K                   int     `json:"k"`
	SetConfidenceDegree float64 `json:"set_confidence_degree"`
}

func GetVector(ctx context.Context, text string) (*VectorRespBody, *VectorSearchConfig, error) {
	var config VectorSearchConfig
	err := rainbow.UnmarshalConfig(ctx, rainbow.KnowledgeVectorConfig, &config)
	if err != nil {
		return nil, nil, errs.WrapSystemErr(err)
	}

	body := &VectorReqBody{
		QueryID:  fmt.Sprintf("%d", time.Now().UnixMilli()),
		Text:     text,
		K:        config.K,
		EmbIndex: config.EmbIndex,
	}
	bodyStr, _ := json.Marshal(body)

	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		SetAuthToken(config.Token).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"Wsid":         config.Wsid,
		}).
		SetBody(bodyStr).
		Post(config.Url)
	if err != nil {
		return nil, nil, errs.WrapSystemErr(err)
	}

	if resp.StatusCode() != 200 {
		return nil, nil, errs.BuildSystemErr(fmt.Sprintf("Http call fail. code:%d", resp.StatusCode()))
	}

	log.ErrorContextf(ctx, "GetVector http resp body:%s", string(resp.Body()))

	res := VectorRespBody{}
	if err = json.Unmarshal(resp.Body(), &res); err != nil {
		return nil, nil, errs.WrapSystemErr(err)
	}

	return &res, &config, nil
}
