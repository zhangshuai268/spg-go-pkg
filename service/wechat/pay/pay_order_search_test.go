package wechat

import (
	spg_go_pkg "github.com/zhangshuai268/spg-go-pkg"
	logger "github.com/zhangshuai268/spg-go-pkg/pkg/log"
	"testing"
)

func TestPay_PayOrderSearch(t *testing.T) {
	pm := make(spg_go_pkg.ParamMap)
	pm.Set("appid", appId)
	pm.Set("mch_id", mchId)
	pm.Set("transaction_id", "SASASASASAS")
	pm.Set("nonce_str", "sadafaf")
	search, resMap, err := client.PayOrderSearch(pm)
	if err != nil {
		return
	}
	logger.Logger.Info(search)
	logger.Logger.Info(resMap)

}
