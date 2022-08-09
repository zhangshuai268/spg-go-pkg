package wechat

import (
	"github.com/zhangshuai268/spg-go-pkg/pkg/http"
	"github.com/zhangshuai268/spg-go-pkg/pkg/util"
)

// UserInfoResponse 获取用户详情Rsp
type UserInfoResponse struct {
	OpenId     string   `json:"openid"`
	NickName   string   `json:"nickname"`
	Sex        int      `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgUrl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
	ErrCode    int      `json:"errcode"`
	ErrMessage string   `json:"errmsg"`
}

func (l *login) UserInfo(accessToken, openId, lang string) (*UserInfoResponse, error) {
	//请求链接
	url := userinfoUrl + "access_token=" + accessToken + "&openid=" + openId + "&lang=" + lang
	wxData, err := http.HttpGet(url)
	if err != nil {
		return nil, err
	}
	//处理返回值
	var res UserInfoResponse
	err = util.StructTo(&wxData, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
