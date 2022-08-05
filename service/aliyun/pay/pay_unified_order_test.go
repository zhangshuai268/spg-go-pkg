package aliyun

import (
	spg_go_pkg "github.com/zhangshuai268/spg-go-pkg"
	logger "github.com/zhangshuai268/spg-go-pkg/pkg/log"
	"testing"
)

func TestPay_PayUnifiedOrder(t *testing.T) {
	pm := make(spg_go_pkg.ParamMap)
	pm.Set("out_trade_no", "sadasdadasd")
	pm.Set("total_amount", 0.01)
	pm.Set("subject", "测试")
	pm.Set("product_code", "QUICK_WAP_WAY")
	pm.Set("return_url", "www.baidu.com")
	pm.Set("notify_url", "www.baidu.com")
	order, err := client.PayUnifiedOrder(pm)
	if err != nil {
		logger.Logger.Info(err.Error())
		return
	}
	logger.Logger.Info(order)
}
