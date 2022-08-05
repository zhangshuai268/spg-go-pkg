package aliyun

import gopay_alipay "github.com/go-pay/gopay/alipay"

func (p *pay) PayVerifySign(bean interface{}) (bool, error) {
	sign, err := gopay_alipay.VerifySign(p.PublicKey, bean)
	if err != nil {
		return false, err
	}
	return sign, nil
}
