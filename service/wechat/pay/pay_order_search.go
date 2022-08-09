package wechat

import (
	"context"
	"github.com/go-pay/gopay"
	gopay_wechat "github.com/go-pay/gopay/wechat"
	spg_go_pkg "github.com/zhangshuai268/spg-go-pkg"
	"github.com/zhangshuai268/spg-go-pkg/pkg/util"
)

type PayOrderResponse gopay_wechat.QueryOrderResponse

func (p *pay) PayOrderSearch(pm spg_go_pkg.ParamMap) (*PayOrderResponse, spg_go_pkg.ParamMap, error) {
	var res PayOrderResponse
	var resPm spg_go_pkg.ParamMap

	var bm gopay.BodyMap
	err := util.StructTo(&pm, &bm)
	if err != nil {
		return nil, nil, err
	}
	client := gopay_wechat.NewClient(p.AppId, p.MchId, p.PaySecret, true)
	//支付签名
	sign := gopay_wechat.GetParamSign(p.AppId, p.MchId, p.PaySecret, bm)
	bm.Set("sign", sign)
	wxRsp, resBm, err := client.QueryOrder(context.TODO(), bm)
	if err != nil {
		return nil, nil, nil
	}
	err = util.StructTo(&wxRsp, &res)
	if err != nil {
		return nil, nil, err
	}
	err = util.StructTo(&resBm, &resPm)
	if err != nil {
		return nil, nil, err
	}

	return &res, resPm, nil
}
