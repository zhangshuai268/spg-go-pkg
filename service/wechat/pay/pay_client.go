package wechat

import (
	"errors"
	spg_go_pkg "github.com/zhangshuai268/spg-go-pkg"
)

type PayService interface {
	PayUnifiedOrder(pm spg_go_pkg.ParamMap) (*PayUnifiedOrderResponse, error)
}

type pay struct {
	AppId     string `json:"app_id"`
	MchId     string `json:"mch_id"`
	PaySecret string `json:"pay_secret"`
}

func NewPayClient(appId, mchId, paySecret string) (PayService, error) {
	if appId == "" {
		return nil, errors.New("缺少appid")
	}
	if mchId == "" {
		return nil, errors.New("缺少mch_id")
	}

	if paySecret == "" {
		return nil, errors.New("缺少pay_secret")
	}

	return &pay{
		AppId:     appId,
		MchId:     mchId,
		PaySecret: paySecret,
	}, nil

}
