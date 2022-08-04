package wechat

import gopay_wechat "github.com/go-pay/gopay/wechat"

func (p *pay) PayVerifySign(signType string, bean interface{}) (bool, error) {
	sign, err := gopay_wechat.VerifySign(p.PaySecret, signType, bean)
	if err != nil {
		return false, err
	}
	return sign, nil
}
