package wechat

import (
	spg_go_pkg "github.com/zhangshuai268/spg-go-pkg"
	logger "github.com/zhangshuai268/spg-go-pkg/pkg/log"
	"testing"
)

func TestMessage_TemplateId(t *testing.T) {
	accessToken := ""
	pm := make(spg_go_pkg.ParamMap)
	pm.Set("template_id_short", "TM00015")
	id, err := client.TemplateId(pm, accessToken)
	if err != nil {
		logger.Logger.Error(err.Error())
		return
	}
	logger.Logger.Info(*id)
}
