package v1

import (
	"errors"
	"github.com/zhangshuai268/spg-go-pkg/pkg"
)

const (
	oauthUrl    = "https://api.weixin.qq.com/sns/oauth2/access_token?"
	userinfoUrl = "https://api.weixin.qq.com/sns/userinfo?"
	authUrl     = "https://api.weixin.qq.com/sns/auth?"
	jscode2Url  = "https://api.weixin.qq.com/sns/jscode2session?"
	grantType   = "authorization_code"
)

type LoginService interface {
	Oauth(code string) (*OauthResponse, error)
	UserInfo(accessToken, openId, lang string) (*UserInfoResponse, error)
	Auth(accessToken, openId string) (*AuthResponse, error)
	JsCode2(code string) (*JsCode2Response, error)
	Mobile(mobile, iv, sessionKey string) (*MobileResponse, error)
}

type login struct {
	Appid     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	GrantType string `json:"grant_type"`
}

func NewLogin(param interface{}) (LoginService, error) {
	var l login
	err := pkg.StructTo(&param, &l)
	if err != nil {
		return nil, err
	}
	if l.Appid == "" {
		return nil, errors.New("缺少app_id")
	}
	if l.AppSecret == "" {
		return nil, errors.New("缺少app_secret")
	}
	if l.GrantType == "" {
		l.GrantType = grantType
	}
	return &l, nil
}
