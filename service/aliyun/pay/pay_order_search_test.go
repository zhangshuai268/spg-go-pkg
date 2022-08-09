package aliyun

import (
	spg_go_pkg "github.com/zhangshuai268/spg-go-pkg"
	logger "github.com/zhangshuai268/spg-go-pkg/pkg/log"
	"testing"
)

func TestPay_PayOrderSearch(t *testing.T) {
	pm := make(spg_go_pkg.ParamMap)
	pm.Set("out_trade_no", "sadasdadasd")
	search, err := client.PayOrderSearch(pm)
	if err != nil {
		return
	}
	logger.Logger.Info(search)
}
