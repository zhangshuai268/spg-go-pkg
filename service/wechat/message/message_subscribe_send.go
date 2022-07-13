package wechat

import spg_go_pkg "github.com/zhangshuai268/spg-go-pkg"

type SubscribeSendResponse struct {
	ErrCode    int    `json:"errcode"`
	ErrMessage string `json:"errmsg"`
}

func (m *message) SubscribeSend(pm spg_go_pkg.ParamMap) (*SubscribeSendResponse, error) {

}

// AppletsSend 小程序发送
func AppletsSend(pm spg_go_pkg.ParamMap)
