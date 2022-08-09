package wechat

import (
	"github.com/zhangshuai268/spg-go-pkg/pkg/http"
	"github.com/zhangshuai268/spg-go-pkg/pkg/util"
)

// JsCode2Response 小程序登录Rsp
type JsCode2Response struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMessage string `json:"errmsg"`
}

func (l *login) JsCode2(code string) (*JsCode2Response, error) {
	url := jscode2Url + "appid=" + l.AppId + "&secret=" + l.AppSecret + "&js_code=" + code + "&grant_type=" + l.GrantType
	wxData, err := http.HttpGet(url)
	if err != nil {
		return nil, err
	}
	//处理返回值
	var res JsCode2Response
	err = util.StructTo(&wxData, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
