package wechat

import (
	"encoding/json"
	"errors"
	spg_go_pkg "github.com/zhangshuai268/spg-go-pkg"
	"github.com/zhangshuai268/spg-go-pkg/pkg/http"
	"github.com/zhangshuai268/spg-go-pkg/pkg/util"
	"github.com/zhangshuai268/spg-go-pkg/service/wechat"
)

type TemplateSendResponse struct {
	ErrCode    int    `json:"errcode"`
	ErrMessage string `json:"errmsg"`
	MsgId      int    `json:"msgid"`
}

func (m *message) TemplateSend(pm spg_go_pkg.ParamMap, accessToken string) (*TemplateSendResponse, error) {
	if m.AppType != wechat.AppTypeO {
		return nil, errors.New("模板消息仅支持公众号")
	}
	url := templateUrl + "access_token=" + accessToken
	var bytes []byte
	bytes, err := json.Marshal(pm)
	if err != nil {
		return nil, err
	}
	sendRes, err := http.HttpPostJson(url, bytes)
	if err != nil {
		return nil, err
	}
	var t TemplateSendResponse
	err = util.StructTo(&sendRes, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
