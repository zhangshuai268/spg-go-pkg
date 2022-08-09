package wechat

import (
	"encoding/json"
	spg_go_pkg "github.com/zhangshuai268/spg-go-pkg"
	"github.com/zhangshuai268/spg-go-pkg/pkg/http"
	"github.com/zhangshuai268/spg-go-pkg/pkg/util"
)

type TemplateIdResponse struct {
	ErrCode    int    `json:"errcode"`
	ErrMessage string `json:"errmsg"`
	TemplateId string `json:"template_id"`
}

func (m *message) TemplateId(pm spg_go_pkg.ParamMap, accessToken string) (*TemplateIdResponse, error) {
	url := templateIdUrl + "access_token=" + accessToken
	var bytes []byte
	bytes, err := json.Marshal(pm)
	if err != nil {
		return nil, err
	}
	sendRes, err := http.HttpPostJson(url, bytes)
	if err != nil {
		return nil, err
	}
	var t TemplateIdResponse
	err = util.StructTo(&sendRes, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
