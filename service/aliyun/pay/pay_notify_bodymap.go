package aliyun

import (
	gopay_alipay "github.com/go-pay/gopay/alipay"
	spg_go_pkg "github.com/zhangshuai268/spg-go-pkg"
	"github.com/zhangshuai268/spg-go-pkg/pkg/util"
	"net/http"
)

func (p *pay) PayParseNotifyToBodyMap(req *http.Request) (spg_go_pkg.ParamMap, error) {
	notifyReq, err := gopay_alipay.ParseNotifyToBodyMap(req)
	if err != nil {
		return nil, err
	}
	var pm spg_go_pkg.ParamMap
	err = util.StructTo(&notifyReq, &pm)
	if err != nil {
		return nil, err
	}
	return pm, nil
}
