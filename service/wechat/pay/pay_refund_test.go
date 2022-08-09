package wechat

import (
	spg_go_pkg "github.com/zhangshuai268/spg-go-pkg"
	logger "github.com/zhangshuai268/spg-go-pkg/pkg/log"
	"testing"
)

func TestPay_PayRefund(t *testing.T) {
	err := client.PayAddCertFilePath("", "")
	if err != nil {
		return
	}
	pm := make(spg_go_pkg.ParamMap)
	pm.Set("appid", appId)
	pm.Set("mch_id", mchId)
	pm.Set("out_trade_no", "SASASASASAS")
	pm.Set("out_refund_no", "sdaa")
	pm.Set("total_fee", 1)
	pm.Set("refund_fee", 1)
	refund, resMap, err := client.PayRefund(pm)
	if err != nil {
		return
	}
	logger.Logger.Info(refund)
	logger.Logger.Info(resMap)

}
