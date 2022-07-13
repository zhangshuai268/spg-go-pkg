package wechat

import (
	"errors"
)

const (
	oauthUrl    = "https://api.weixin.qq.com/sns/oauth2/access_token?" //微信公共号网页授权登录
	userinfoUrl = "https://api.weixin.qq.com/sns/userinfo?"            //获取用户信息
	authUrl     = "https://api.weixin.qq.com/sns/auth?"                //access_token有效期
	jscode2Url  = "https://api.weixin.qq.com/sns/jscode2session?"      //微信小程序授权登录
	tokenUrl    = "https://api.weixin.qq.com/cgi-bin/token?"           //微信获取接口调用凭证
	grantTypeA  = "authorization_code"                                 //grantTypeA
	grantTypeC  = "client_credential"                                  //grantTypeC
)

type LoginService interface {
	// Oauth 微信公众号网页授权登录
	Oauth(code string) (*OauthResponse, error)
	// UserInfo 获取用户信息
	UserInfo(accessToken, openId, lang string) (*UserInfoResponse, error)
	// Auth 判断access_token有效期
	Auth(accessToken, openId string) (*AuthResponse, error)
	// JsCode2 微信小程序登录
	JsCode2(code string) (*JsCode2Response, error)
	// Mobile 授权手机号
	Mobile(mobile, iv, sessionKey string) (*MobileResponse, error)
	// GetAccessToken 获取接口调用凭证access_token,与登录模块不同
	GetAccessToken() (*AccessTokenResponse, error)
}

type login struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	GrantType string `json:"grant_type"`
}

// NewLoginClient 初始化登录客户端
func NewLoginClient(appId, appSecret string) (LoginService, error) {
	if appId == "" {
		return nil, errors.New("缺少app_id")
	}
	if appSecret == "" {
		return nil, errors.New("缺少app_secret")
	}
	return &login{
		AppId:     appId,
		AppSecret: appSecret,
		GrantType: grantTypeA,
	}, nil
}
