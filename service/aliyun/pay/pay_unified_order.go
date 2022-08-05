package aliyun

import (
	"context"
	"errors"
	"github.com/go-pay/gopay"
	gopay_alipay "github.com/go-pay/gopay/alipay"
	spg_go_pkg "github.com/zhangshuai268/spg-go-pkg"
	"github.com/zhangshuai268/spg-go-pkg/pkg/util"
	"github.com/zhangshuai268/spg-go-pkg/service/aliyun"
)

func (p *pay) PayUnifiedOrder(pm spg_go_pkg.ParamMap) (string, error) {
	var bm gopay.BodyMap
	err := util.StructTo(&pm, &bm)
	if err != nil {
		return "", nil
	}
	client, err := gopay_alipay.NewClient(p.AppId, p.PrivateKey, p.IsProd)
	if err != nil {
		return "", nil
	}
	switch p.PayType {
	case aliyun.TradeTypeApp:
		appPay, err := client.TradeAppPay(context.TODO(), bm)
		if err != nil {
			return "", err
		}
		return appPay, nil
	case aliyun.TradeTypeH5:
		pagePay, err := client.TradePagePay(context.TODO(), bm)
		if err != nil {
			return "", err
		}
		return pagePay, nil
	case aliyun.TradeTypeWeb:
		wapPay, err := client.TradeWapPay(context.TODO(), bm)
		if err != nil {
			return "", err
		}
		return wapPay, nil
	default:
		return "", errors.New("暂不支持其他种类的支付")
	}
}
