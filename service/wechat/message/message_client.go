package wechat

import (
	"errors"
	spg_go_pkg "github.com/zhangshuai268/spg-go-pkg"
)

const (
	tokenUrl   = "https://api.weixin.qq.com/cgi-bin/token?" //微信获取接口调用凭证
	grantTypeC = "client_credential"                        //grantTypeC
)

type MessageService interface {
	// GetAccessToken 获取接口调用凭证access_token
	GetAccessToken() (*AccessTokenResponse, error)
	// SubscribeSend 订阅消息发送
	SubscribeSend(pm spg_go_pkg.ParamMap) (*SubscribeSendResponse, error)
}

type message struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	AppType   string `json:"app_type"`
}

func NewMessageClient(appId, appSecret, appType string) (MessageService, error) {
	if appId == "" {
		return nil, errors.New("缺少app_id")
	}

	if appSecret == "" {
		return nil, errors.New("缺少app_secret")
	}

	if appType == "" {
		return nil, errors.New("缺少app_type")
	}

	return &message{
		AppId:     appId,
		AppSecret: appSecret,
		AppType:   appType,
	}, nil
}
