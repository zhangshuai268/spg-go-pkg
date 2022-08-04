package wechat

import gopay_wechat "github.com/go-pay/gopay/wechat"

func (p *pay) PaySign(nonceStr string, packages string, signType string, timeStamp string) string {
	paySign := gopay_wechat.GetMiniPaySign(p.AppId, nonceStr, packages, signType, timeStamp, p.PaySecret)
	return paySign
}
