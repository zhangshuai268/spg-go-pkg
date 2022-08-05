package aliyun

import (
	"context"
	"github.com/go-pay/gopay"
	gopay_alipay "github.com/go-pay/gopay/alipay"
	spg_go_pkg "github.com/zhangshuai268/spg-go-pkg"
	"github.com/zhangshuai268/spg-go-pkg/pkg/util"
)

type PayOrderResponse gopay_alipay.TradeQueryResponse

func (p *pay) PayOrderSearch(pm spg_go_pkg.ParamMap) (*PayOrderResponse, error) {
	var bm gopay.BodyMap
	var res PayOrderResponse
	err := util.StructTo(&pm, &bm)
	if err != nil {
		return nil, nil
	}
	client, err := gopay_alipay.NewClient(p.AppId, p.PrivateKey, p.IsProd)
	query, err := client.TradeQuery(context.TODO(), bm)
	if err != nil {
		return nil, err
	}
	err = util.StructTo(query, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
