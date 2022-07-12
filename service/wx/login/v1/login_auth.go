package v1

import (
	"github.com/zhangshuai268/spg-go-pkg/pkg"
)

type AuthResponse struct {
	ErrCode    int    `json:"errcode"`
	ErrMessage string `json:"errmsg"`
}

func (l *login) Auth(accessToken, openId string) (*AuthResponse, error) {
	url := authUrl + "access_token=" + accessToken + "&openid=" + openId
	wxData, err := pkg.HttpGet(url)
	if err != nil {
		return nil, err
	}
	//处理返回值
	var res *AuthResponse
	err = pkg.StructTo(&wxData, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
