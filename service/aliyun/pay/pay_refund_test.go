package aliyun

import (
	spg_go_pkg "github.com/zhangshuai268/spg-go-pkg"
	logger "github.com/zhangshuai268/spg-go-pkg/pkg/log"
	"testing"
)

func TestPay_PayRefund(t *testing.T) {
	pm := make(spg_go_pkg.ParamMap)
	pm.Set("out_trade_no", "sadasdadasd")
	pm.Set("refund_amount", 0.01)
	order, err := client.PayRefund(pm)
	if err != nil {
		logger.Logger.Info(err.Error())
		return
	}
	logger.Logger.Info(order)
}
