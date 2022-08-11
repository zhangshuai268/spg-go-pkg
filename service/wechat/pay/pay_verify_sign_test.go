package wechat

import (
	spg_go_pkg "github.com/zhangshuai268/spg-go-pkg"
	logger "github.com/zhangshuai268/spg-go-pkg/pkg/logger"
	"github.com/zhangshuai268/spg-go-pkg/service/wechat"
	"testing"
)

func TestPay_PayVerifySign(t *testing.T) {
	var pm spg_go_pkg.ParamMap
	sign, err := client.PayVerifySign(wechat.SignTypeMD5, pm)
	if err != nil {
		return
	}
	logger.Logger.Info(sign)
}
