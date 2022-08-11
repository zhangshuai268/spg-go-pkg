package wechat

import (
	spg_go_pkg "github.com/zhangshuai268/spg-go-pkg"
	logger "github.com/zhangshuai268/spg-go-pkg/pkg/logger"
	"github.com/zhangshuai268/spg-go-pkg/service/wechat"
	"testing"
)

func TestPay_PayUnifiedOrder(t *testing.T) {
	pm := make(spg_go_pkg.ParamMap)
	pm.Set("appid", appId)
	pm.Set("mch_id", mchId)
	pm.Set("body", "body")
	pm.Set("out_trade_no", "SASASASASAS")
	pm.Set("total_fee", 0.01)
	pm.Set("openid", "assas")
	pm.Set("nonce_str", "sadafaf")
	pm.Set("spbill_create_ip", "127.0.0.1")
	pm.Set("sign_type", wechat.SignTypeMD5)
	pm.Set("notify_url", "127.0.0.1/s")
	pm.Set("trade_type", wechat.TradeTypeApp)
	order, err := client.PayUnifiedOrder(pm)
	if err != nil {
		logger.Logger.Error(err.Error())
		return
	}
	logger.Logger.Info(*order)

}
