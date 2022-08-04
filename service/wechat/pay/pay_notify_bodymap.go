package wechat

import (
	"github.com/go-pay/gopay/wechat"
	spg_go_pkg "github.com/zhangshuai268/spg-go-pkg"
	"github.com/zhangshuai268/spg-go-pkg/pkg/util"
	"net/http"
)

func (p *pay) PayParseNotifyToBodyMap(req *http.Request) (spg_go_pkg.ParamMap, error) {
	notifyReq, err := wechat.ParseNotifyToBodyMap(req)
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
