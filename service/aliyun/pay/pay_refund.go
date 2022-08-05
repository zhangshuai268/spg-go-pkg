package aliyun

import (
	"context"
	"github.com/go-pay/gopay"
	gopay_alipay "github.com/go-pay/gopay/alipay"
	spg_go_pkg "github.com/zhangshuai268/spg-go-pkg"
	"github.com/zhangshuai268/spg-go-pkg/pkg/util"
)

type PayRefundResponse gopay_alipay.TradeRefundResponse

func (p *pay) PayRefund(pm spg_go_pkg.ParamMap) (*PayRefundResponse, error) {
	var bm gopay.BodyMap
	var res PayRefundResponse
	err := util.StructTo(&pm, &bm)
	if err != nil {
		return nil, nil
	}
	client, err := gopay_alipay.NewClient(p.AppId, p.PrivateKey, p.IsProd)
	refund, err := client.TradeRefund(context.TODO(), bm)
	if err != nil {
		return nil, err
	}
	err = util.StructTo(refund, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
