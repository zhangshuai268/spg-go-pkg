package v1

import "github.com/zhangshuai268/spg-go-pkg/pkg"

type JsCode2Response struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    string `json:"errcode"`
	ErrMessage string `json:"errmsg"`
}

func (l *login) JsCode2(code string) (*JsCode2Response, error) {
	url := jscode2Url + "?appid=" + l.Appid + "&secret=" + l.AppSecret + "&js_code=" + code + "&grant_type=" + l.GrantType
	wxData, err := pkg.HttpGet(url)
	if err != nil {
		return nil, err
	}
	//处理返回值
	var res JsCode2Response
	err = pkg.StructTo(&wxData, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
