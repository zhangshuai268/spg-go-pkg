package wechat

import (
	"github.com/zhangshuai268/spg-go-pkg/pkg/http"
	"github.com/zhangshuai268/spg-go-pkg/pkg/util"
)

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	ErrCode     int    `json:"errcode"`
	ErrMessage  string `json:"errmsg"`
}

func (l *login) GetAccessToken() (*AccessTokenResponse, error) {
	url := tokenUrl + "grant_type=" + grantTypeC + "&appid=" + l.AppId + "&secret=" + l.AppSecret
	wxData, err := http.HttpGet(url)
	if err != nil {
		return nil, err
	}
	//处理返回值
	var res AccessTokenResponse
	err = util.StructTo(&wxData, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
