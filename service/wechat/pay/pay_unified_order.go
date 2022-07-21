package wechat

import (
	"context"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/qq"
	"github.com/go-pay/gopay/wechat"
	spg_go_pkg "github.com/zhangshuai268/spg-go-pkg"
	"github.com/zhangshuai268/spg-go-pkg/pkg/util"
)

type PayUnifiedOrderResponse qq.UnifiedOrderResponse

func (p *pay) PayUnifiedOrder(pm spg_go_pkg.ParamMap) (*PayUnifiedOrderResponse, error) {
	var res PayUnifiedOrderResponse
	client := wechat.NewClient(p.AppId, p.MchId, p.PaySecret, true)
	//使用gopay的模板参数
	bm := make(gopay.BodyMap)
	err := util.StructTo(&pm, &bm)
	if err != nil {
		return nil, err
	}
	//支付签名
	sign := wechat.GetParamSign(p.AppId, p.MchId, p.PaySecret, bm)
	bm.Set("sign", sign)
	//调用下单
	wxRsp, err := client.UnifiedOrder(context.Background(), bm)
	err = util.StructTo(wxRsp, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
