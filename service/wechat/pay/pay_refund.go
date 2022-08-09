package wechat

import (
	"context"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/qq"
	gopay_wechat "github.com/go-pay/gopay/wechat"
	spg_go_pkg "github.com/zhangshuai268/spg-go-pkg"
	"github.com/zhangshuai268/spg-go-pkg/pkg/util"
	"io/ioutil"
	"os"
)

type PayRefundResponse qq.RefundResponse

func (p *pay) PayRefund(pm spg_go_pkg.ParamMap) (*PayRefundResponse, spg_go_pkg.ParamMap, error) {
	var res PayRefundResponse
	var resPm spg_go_pkg.ParamMap
	//生成gopay客户端
	client := gopay_wechat.NewClient(p.AppId, p.MchId, p.PaySecret, true)
	err := addCertContent(client, p.CertFilePath, p.KeyFilePath)
	if err != nil {
		return nil, nil, err
	}
	//使用gopay的模板参数
	bm := make(gopay.BodyMap)
	err = util.StructTo(&pm, &bm)
	if err != nil {
		return nil, nil, err
	}
	client.SetCountry(gopay_wechat.China)
	//生成签名
	sign := gopay_wechat.GetParamSign(p.AppId, p.MchId, p.PaySecret, bm)
	bm.Set("sign", sign)
	wxRsp, resBm, err := client.Refund(context.Background(), bm)
	if err != nil {
		return nil, nil, err
	}
	//处理参数
	err = util.StructTo(wxRsp, &res)
	if err != nil {
		return nil, nil, err
	}
	err = util.StructTo(&resBm, &resPm)
	if err != nil {
		return nil, nil, err
	}
	return &res, resPm, nil
}

func addCertContent(client *gopay_wechat.Client, certFilePath, keyFilePath string) error {
	certFile, err := os.Open(certFilePath)
	if err != nil {
		return err
	}

	certContent, err := ioutil.ReadAll(certFile)
	if err != nil {
		return err
	}

	keyfile, err := os.Open(keyFilePath)
	if err != nil {
		return err
	}

	keyContent, err := ioutil.ReadAll(keyfile)
	if err != nil {
		return err
	}

	err = client.AddCertPemFileContent(certContent, keyContent)
	if err != nil {
		return err
	}
	return nil
}
