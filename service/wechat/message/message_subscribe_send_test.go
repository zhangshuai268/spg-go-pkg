package wechat

import (
	spg_go_pkg "github.com/zhangshuai268/spg-go-pkg"
	logger "github.com/zhangshuai268/spg-go-pkg/pkg/logger"
	"testing"
)

func TestMessage_SubscribeSend(t *testing.T) {
	//accessToken := "58_RMEIPfrWtSUvBPcUwSnADrtb6t4u0n8cWEtkObz8fp3Q8_Qe5LOaKS9b4fOgVsZBFFURF4E38oWT9v3I2DrCZdLAcdO5aP4eGwg7v1fPgesenNwxPHWqkXDMSr03NP1PLQO2nGxDqvOSAq4OTMFdAGAEXQ"
	accessTokenJs, _ := client.GetAccessToken()
	openId := "oU9FgwpLnEtLduOShKwEZ8Rb02B0"
	pm := make(spg_go_pkg.ParamMap)
	pm.Set("touser", openId)
	pm.Set("template_id", "Sc9Lf_UApLMf8mwDRVZSvMrWXDKRjV6cIQ6KQnP--Xg")
	pm.Set("page", "")
	pm.Set("data", map[string]interface{}{
		"name1": map[string]string{
			"value": "灰太狼",
		},
		"phrase2": map[string]string{
			"value": "通过",
		},
	})
	send, err := clientJs.SubscribeSend(pm, accessTokenJs.AccessToken)
	if err != nil {
		logger.Logger.Error(err)
		return
	}
	logger.Logger.Info(*send)
	return

}
