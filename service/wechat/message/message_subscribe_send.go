package wechat

import (
	"encoding/json"
	spg_go_pkg "github.com/zhangshuai268/spg-go-pkg"
	"github.com/zhangshuai268/spg-go-pkg/pkg/http"
	"github.com/zhangshuai268/spg-go-pkg/pkg/util"
	"github.com/zhangshuai268/spg-go-pkg/service/wechat"
)

type SubscribeSendResponse struct {
	ErrCode    int    `json:"errcode"`
	ErrMessage string `json:"errmsg"`
}

func (m *message) SubscribeSend(pm spg_go_pkg.ParamMap, accessToken string) (*SubscribeSendResponse, error) {
	url := ""
	//判断客户端类型
	switch m.AppType {
	//小程序
	case wechat.AppTypeA:
		url = appletUrl
	//公众号
	case wechat.AppTypeO:
		url = officialUrl
	default:
	}
	//发送
	res, err := doSend(pm, accessToken, url)
	if err != nil {
		return nil, err
	}
	var s SubscribeSendResponse
	err = util.StructTo(&res, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil

}

// DoSend 发送
func doSend(pm spg_go_pkg.ParamMap, accessToken string, sendUrl string) (map[string]interface{}, error) {
	sendUrl = sendUrl + "access_token=" + accessToken
	var bytes []byte
	bytes, err := json.Marshal(pm)
	if err != nil {
		return nil, err
	}
	sendRes, err := http.HttpPostJson(sendUrl, bytes)
	if err != nil {
		return nil, err
	}
	return sendRes, nil
}
