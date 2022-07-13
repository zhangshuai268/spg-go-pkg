package wechat

import (
	logger "github.com/zhangshuai268/spg-go-pkg/pkg/log"
	"github.com/zhangshuai268/spg-go-pkg/service/wechat"
	"testing"
)

func TestLogin_UserInfo(t *testing.T) {
	openId := "oU9FgwpLnEtLduOShKwEZ8Rb02B0"
	accessToken := "58_aef-0wLmtcU5Q6VZRQxoibBFlKUnJl5eg6nuGx7JXf6sxq26IdMjjJbTTkDiMwKiGVYsA75kbbH2zyQIG17Vj-_Ze_GH9uJ8wmAmapaD2jc"
	info, err := client.UserInfo(accessToken, openId, wechat.LandCN)
	if err != nil {
		logger.Logger.Error(err.Error())
		return
	}
	logger.Logger.Info(*info)
}
