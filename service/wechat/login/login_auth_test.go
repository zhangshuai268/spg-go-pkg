package wechat

import (
	logger "github.com/zhangshuai268/spg-go-pkg/pkg/log"
	"testing"
)

func TestLogin_Auth(t *testing.T) {
	openId := "oU9FgwpLnEtLduOShKwEZ8Rb02B0"
	accessToken := "58_aef-0wLmtcU5Q6VZRQxoibBFlKUnJl5eg6nuGx7JXf6sxq26IdMjjJbTTkDiMwKiGVYsA75kbbH2zyQIG17Vj-_Ze_GH9uJ8wmAmapaD2jc"
	auth, err := client.Auth(accessToken, openId)
	if err != nil {
		logger.Logger.Error(err)
		return
	}
	logger.Logger.Info(auth.ErrMessage)
	return
}
