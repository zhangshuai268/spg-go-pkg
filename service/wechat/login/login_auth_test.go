package wechat

import (
	logger "github.com/zhangshuai268/spg-go-pkg/pkg/log"
	"testing"
)

func TestLogin_Auth(t *testing.T) {
	openId := ""
	accessToken := ""
	auth, err := client.Auth(accessToken, openId)
	if err != nil {
		logger.Logger.Error(err)
		return
	}
	logger.Logger.Info(auth.ErrMessage)
	return
}
