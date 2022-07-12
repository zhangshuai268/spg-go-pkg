package v1

import (
	"spg-go-pkg/pkg"
)

type OauthResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenId       string `json:"open_id"`
	Scope        string `json:"scope"`
	ErrCode      int    `json:"errcode"`
	ErrMessage   string `json:"errmsg"`
}

func (l *login) Oauth(code string) (*OauthResponse, error) {
	//请求链接
	url := oauthUrl + "appid=" + l.Appid + "&secret=" + l.AppSecret + "&code=" + code + "&grant_type=" + l.GrantType
	wxData, err := pkg.HttpGet(url)
	if err != nil {
		return nil, err
	}
	//处理返回值
	var res *OauthResponse
	err = pkg.StructTo(&wxData, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
