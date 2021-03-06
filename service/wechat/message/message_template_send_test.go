package wechat

import (
	spg_go_pkg "github.com/zhangshuai268/spg-go-pkg"
	logger "github.com/zhangshuai268/spg-go-pkg/pkg/log"
	"testing"
)

func TestMessage_TemplateSend(t *testing.T) {
	accessToken := ""
	pm := make(spg_go_pkg.ParamMap)
	pm.Set("touser", "")
	pm.Set("template_id", "")
	pm.Set("url", "")
	pm.Set("miniprogram", map[string]interface{}{
		"appid":    "xiaochengxuappid12345",
		"pagepath": "index?foo=bar",
	})
	pm.Set("client_msg_id", "MSG_000001")
	pm.Set("data", map[string]interface{}{
		"first": map[string]interface{}{
			"value": "恭喜你购买成功！",
			"color": "#173177",
		},
		"keyword1": map[string]interface{}{
			"value": "巧克力",
			"color": "#173177",
		},
	})
	send, err := client.TemplateSend(pm, accessToken)
	if err != nil {
		logger.Logger.Error(err.Error())
		return
	}
	logger.Logger.Info(*send)
}
