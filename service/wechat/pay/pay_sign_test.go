package wechat

import (
	"github.com/zhangshuai268/spg-go-pkg/service/wechat"
	"testing"
)

func TestPay_PaySign(t *testing.T) {
	nonce := ""
	packages := ""
	timeStamp := ""
	client.PaySign(nonce, packages, wechat.TradeTypeApp, timeStamp)
}
