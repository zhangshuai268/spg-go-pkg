package wechat

import (
	"errors"
	spg_go_pkg "github.com/zhangshuai268/spg-go-pkg"
)

const (
	tokenUrl      = "https://api.weixin.qq.com/cgi-bin/token?"                      //微信获取接口调用凭证
	grantTypeC    = "client_credential"                                             //grantTypeC
	appletUrl     = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send?"     //小程序订阅消息发送
	officialUrl   = "https://api.weixin.qq.com/cgi-bin/message/template/subscribe?" //公众号订阅消息发送
	templateUrl   = "https://api.weixin.qq.com/cgi-bin/message/template/send?"      //公众号模板消息发送
	templateIdUrl = "https://api.weixin.qq.com/cgi-bin/template/api_add_template?"  //获取模板消息id
)

type MessageService interface {
	// GetAccessToken 获取接口调用凭证access_token
	GetAccessToken() (*AccessTokenResponse, error)
	// SubscribeSend 订阅消息发送
	SubscribeSend(pm spg_go_pkg.ParamMap, accessToken string) (*SubscribeSendResponse, error)
	// TemplateSend 模板消息发送，仅支持公众号
	TemplateSend(pm spg_go_pkg.ParamMap, accessToken string) (*TemplateSendResponse, error)
	// TemplateId 获取模板消息id
	TemplateId(pm spg_go_pkg.ParamMap, accessToken string) (*TemplateIdResponse, error)
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
